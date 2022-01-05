package main

import (
	"aoc2019/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const INPUT_PATH = "day3/input.txt"

type instruction struct {
	direction byte
	count     int
}

type position struct {
	x int
	y int
}

func (p position) distance() int {
	return utils.AbsInt(p.x) + utils.AbsInt(p.y)
}

func parseInput(input string) [][]instruction {
	var instr_list [][]instruction

	for _, l := range strings.Split(input, "\n") {
		var instr []instruction

		for _, str_inst := range strings.Split(l, ",") {
			direction := str_inst[0]
			count, err := strconv.Atoi(str_inst[1:])

			if err != nil {
				panic(err)
			}

			instr = append(instr, instruction{
				direction: direction,
				count:     count,
			})
		}

		instr_list = append(instr_list, instr)
	}

	return instr_list
}

func gen_hitmap(instructions []instruction) map[position]int {
	hitmap := make(map[position]int)

	current_pos := position{
		x: 0,
		y: 0,
	}

	steps := 0

	for _, instr := range instructions {
		x_delta := 0
		y_delta := 0

		switch instr.direction {
		case 'R':
			x_delta = 1
		case 'L':
			x_delta = -1
		case 'U':
			y_delta = 1
		case 'D':
			y_delta = -1
		default:
			panic(fmt.Sprintf("Unexpected direction %c", instr.direction))
		}

		for i := 0; i < instr.count; i++ {
			steps++

			current_pos.x += x_delta
			current_pos.y += y_delta

			if _, ok := hitmap[current_pos]; !ok {
				hitmap[current_pos] = steps
			}

		}
	}

	return hitmap
}

func part1_solve(str_input string) int {
	input := parseInput(str_input)

	wire0_hitmap := gen_hitmap(input[0])
	wire1_hitmap := gen_hitmap(input[1])

	result := math.MaxInt

	for pos := range wire0_hitmap {
		if _, ok := wire1_hitmap[pos]; ok {
			result = utils.MinInt(result, pos.distance())
		}
	}

	return result
}

func part2_solve(str_input string) int {
	input := parseInput(str_input)

	wire0_hitmap := gen_hitmap(input[0])
	wire1_hitmap := gen_hitmap(input[1])

	result := math.MaxInt

	for pos, wire0_steps := range wire0_hitmap {
		if wire1_steps, ok := wire1_hitmap[pos]; ok {
			result = utils.MinInt(result, wire0_steps+wire1_steps)
		}
	}

	return result
}

func part1() int {
	utils.AssertEq(part1_solve("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"), 159)
	utils.AssertEq(part1_solve("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"), 135)

	result := part1_solve(utils.LoadFile(INPUT_PATH))
	utils.AssertEq(result, 217)

	return result
}

func part2() int {
	utils.AssertEq(part2_solve("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"), 610)
	utils.AssertEq(part2_solve("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"), 410)

	result := part2_solve(utils.LoadFile(INPUT_PATH))
	utils.AssertEq(result, 3454)

	return result
}

func main() {
	fmt.Printf("Part1: %d\n", part1())
	fmt.Printf("Part2: %d\n", part2())
}
