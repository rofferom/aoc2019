package main

import (
	"aoc2019/utils"
	"fmt"
)

const INPUT_PATH = "day1/input.txt"

func getRequiredFuel(mass int) int {
	return (mass / 3) - 2
}

func part1() int {
	utils.AssertEq(getRequiredFuel(12), 2)
	utils.AssertEq(getRequiredFuel(14), 2)
	utils.AssertEq(getRequiredFuel(1969), 654)
	utils.AssertEq(getRequiredFuel(100756), 33583)

	sum := 0
	for _, fuel := range utils.LoadIntFile(INPUT_PATH) {
		sum += getRequiredFuel(fuel)
	}

	utils.AssertEq(sum, 3233481)

	return sum
}

func getRequiredFuelPart2(mass int) int {
	sum := 0

	for {
		fuel := getRequiredFuel(mass)
		if fuel <= 0 {
			break
		}

		sum += fuel
		mass = fuel
	}

	return sum
}

func part2() int {
	utils.AssertEq(getRequiredFuelPart2(14), 2)
	utils.AssertEq(getRequiredFuelPart2(1969), 966)
	utils.AssertEq(getRequiredFuelPart2(100756), 50346)

	sum := 0
	for _, fuel := range utils.LoadIntFile(INPUT_PATH) {
		sum += getRequiredFuelPart2(fuel)
	}

	utils.AssertEq(sum, 4847351)

	return sum
}

func main() {
	fmt.Printf("Part1: %d\n", part1())
	fmt.Printf("Part2: %d\n", part2())
}
