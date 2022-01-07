package main

import (
	"aoc2019/utils"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const INPUT_PATH = "day5/input.txt"

func parseInput(s string) []int {
	str_values := strings.Split(s, ",")
	int_values := make([]int, len(str_values))

	for i, s := range str_values {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		int_values[i] = v
	}

	return int_values
}

func loadInputFile() []int {
	input, err := ioutil.ReadFile(INPUT_PATH)
	if err != nil {
		panic(err)
	}

	return parseInput(string(input))
}

func getParam(opcode, paramIdx int, values []int, pc int) int {
	addressingMode := (opcode / int(math.Pow10(paramIdx+1))) % 10
	param := values[pc+paramIdx]

	switch addressingMode {
	case 0:
		return values[param]

	case 1:
		return param
	}

	return 0
}

func solve(values []int, inputValue int) int {
	lastOutput := 0

	for pc := 0; values[pc] != 99; {
		opcode := values[pc]

		switch opcode % 100 {
		case 1:
			a := getParam(opcode, 1, values, pc)
			b := getParam(opcode, 2, values, pc)
			r := values[pc+3]
			pc += 4

			values[r] = a + b

		case 2:
			a := getParam(opcode, 1, values, pc)
			b := getParam(opcode, 2, values, pc)
			r_idx := values[pc+3]
			pc += 4

			values[r_idx] = a * b

		case 3:
			a_idx := values[pc+1]
			pc += 2

			values[a_idx] = inputValue

		case 4:
			a_idx := values[pc+1]
			pc += 2

			lastOutput = values[a_idx]

		case 5:
			a := getParam(opcode, 1, values, pc)
			if a != 0 {
				pc = getParam(opcode, 2, values, pc)
			} else {
				pc += 3
			}

		case 6:
			a := getParam(opcode, 1, values, pc)
			if a == 0 {
				pc = getParam(opcode, 2, values, pc)
			} else {
				pc += 3
			}

		case 7:
			a := getParam(opcode, 1, values, pc)
			b := getParam(opcode, 2, values, pc)
			r_idx := values[pc+3]
			pc += 4

			if a < b {
				values[r_idx] = 1
			} else {
				values[r_idx] = 0
			}

		case 8:
			a := getParam(opcode, 1, values, pc)
			b := getParam(opcode, 2, values, pc)
			r_idx := values[pc+3]
			pc += 4

			if a == b {
				values[r_idx] = 1
			} else {
				values[r_idx] = 0
			}

		default:
			panic(fmt.Sprintf("Unexpected opcode %d", opcode))
		}
	}

	return lastOutput
}

func part1() int {
	input := loadInputFile()

	result := solve(input, 1)
	utils.AssertEq(result, 9431221)

	return result
}

func part2() int {
	input := loadInputFile()

	result := solve(input, 5)
	utils.AssertEq(result, 1409363)

	return result
}

func main() {
	fmt.Printf("Part1: %d\n", part1())
	fmt.Printf("Part2: %d\n", part2())
}
