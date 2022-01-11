package main

import (
	"aoc2019/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

const INPUT_PATH = "day10/input.txt"

type vector struct {
	x int
	y int
}

type grid struct {
	points [][]bool
	width  int
	height int
}

func (g *grid) positionValid(v *vector) bool {
	return 0 <= v.x && v.x < g.width && 0 <= v.y && v.y < g.height
}

func parseInput(input string) grid {
	points := [][]bool{}

	for _, strLine := range strings.Split(input, "\n") {
		line := make([]bool, len(strLine))

		for i := 0; i < len(strLine); i++ {
			line[i] = strLine[i] == '#'
		}

		points = append(points, line)
	}

	return grid{
		points: points,
		width:  len(points[0]),
		height: len(points),
	}
}

func getDirections(width, height int) []vector {
	// Generate unordered directions
	directionMap := make(map[vector]bool)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x == 0 && y == 0 {
				continue
			}

			d := utils.Gcd(x, y)

			deltaX := x / d
			deltaY := y / d

			directionMap[vector{x: deltaX, y: deltaY}] = true
			directionMap[vector{x: -deltaX, y: deltaY}] = true
			directionMap[vector{x: deltaX, y: -deltaY}] = true
			directionMap[vector{x: -deltaX, y: -deltaY}] = true
		}
	}

	// Sort them by angle. Oriented from (0,1), clockwise
	getAngle := func(v vector) float64 {
		norm := math.Sqrt(float64(v.x*v.x + v.y*v.y))
		angle := math.Acos((float64(v.x) / norm))
		if v.y < 0 {
			angle = -3. * math.Pi / 2
		} else {
			angle = math.Pi/2 - angle
		}

		if angle < 0 {
			return angle + 2*math.Pi
		} else {
			return angle
		}
	}

	directions := []vector{}
	for dir := range directionMap {
		directions = append(directions, dir)
	}

	sort.Slice(directions, func(i, j int) bool {
		return getAngle(directions[i]) < getAngle(directions[j])
	})

	return directions
}

func getReachablePoint(g *grid, dir *vector, startPos *vector) (bool, vector) {
	for i := 1; ; i++ {
		visitPos := vector{
			x: startPos.x + dir.x*i,
			y: startPos.y - dir.y*i,
		}

		if !g.positionValid(&visitPos) {
			break
		}

		if g.points[visitPos.y][visitPos.x] {
			return true, visitPos
		}
	}

	return false, vector{x: -1, y: -1}
}

func getReachableCount(g *grid, directions []vector, startPos *vector) int {
	count := 0

	for i := 0; i < len(directions); i++ {
		found, _ := getReachablePoint(g, &directions[i], startPos)
		if found {
			count++
		}
	}

	return count
}

func getMonitoringStationPos(g *grid) (vector, int) {
	directions := getDirections(g.width, g.height)

	maxCount := math.MinInt
	maxPosition := vector{x: -1, y: -1}

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			if !g.points[y][x] {
				continue
			}

			count := getReachableCount(g, directions, &vector{x: x, y: y})
			if count > maxCount {
				maxCount = count
				maxPosition = vector{x: x, y: y}
			}
		}
	}

	return maxPosition, maxCount
}

func getNthHit(g *grid, wantedHitId int) int {
	monitorStatPos, _ := getMonitoringStationPos(g)
	directions := getDirections(g.width, g.height)

	hitIdx := 0

	for dirIdx := 0; ; dirIdx = (dirIdx + 1) % len(directions) {
		found, coords := getReachablePoint(g, &directions[dirIdx], &monitorStatPos)
		if found {
			hitIdx++
			if hitIdx == wantedHitId {
				return coords.x*100 + coords.y
			} else {
				g.points[coords.y][coords.x] = false
			}
		}
	}
}

func testInput(input string, x, y, count int) {
	g := parseInput(input)
	pos, c := getMonitoringStationPos(&g)

	utils.AssertEq(pos.x, x)
	utils.AssertEq(pos.y, y)
	utils.AssertEq(c, count)
}

func part1() int {
	testInput(".#..#\n.....\n#####\n....#\n...##",
		3, 4, 8)

	testInput("......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####",
		5, 8, 33)

	testInput("#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###.",
		1, 2, 35)

	testInput(".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##",
		11, 13, 210)

	g := parseInput(utils.LoadFile(INPUT_PATH))
	pos, count := getMonitoringStationPos(&g)
	utils.AssertEq(pos.x, 26)
	utils.AssertEq(pos.y, 29)
	utils.AssertEq(count, 299)

	return count
}

func part2() int {
	{
		g := parseInput(".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##")
		utils.AssertEq(getNthHit(&g, 200), 802)
	}

	g := parseInput(utils.LoadFile(INPUT_PATH))
	result := getNthHit(&g, 200)
	utils.AssertEq(result, 1419)
	return result
}

func main() {
	fmt.Printf("Part1: %d\n", part1())
	fmt.Printf("Part2: %d\n", part2())
}
