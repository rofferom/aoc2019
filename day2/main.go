package main

import (
	"aoc2019/utils"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPUT_PATH = "day2/input.txt"

func parse_input(s string) []int {
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

func load_inputfile() []int {
	input, err := ioutil.ReadFile(INPUT_PATH)
	if err != nil {
		panic(err)
	}

	return parse_input(string(input))
}

func solve(input []int) []int {
	for i := 0; input[i] != 99; i += 4 {
		opcode := input[i]
		a_idx := input[i+1]
		b_idx := input[i+2]
		r_idx := input[i+3]

		switch opcode {
		case 1:
			input[r_idx] = input[a_idx] + input[b_idx]
		case 2:
			input[r_idx] = input[a_idx] * input[b_idx]
		default:
			panic(fmt.Sprintf("Unexpected opcode %d", opcode))
		}
	}

	return input
}

func part1() int {
	utils.AssertIntArrayEq(solve(parse_input("1,0,0,0,99")), []int{2, 0, 0, 0, 99})
	utils.AssertIntArrayEq(solve(parse_input("2,3,0,3,99")), []int{2, 3, 0, 6, 99})
	utils.AssertIntArrayEq(solve(parse_input("2,4,4,5,99,0")), []int{2, 4, 4, 5, 99, 9801})
	utils.AssertIntArrayEq(solve(parse_input("1,1,1,4,99,5,6,0,99")), []int{30, 1, 1, 4, 2, 5, 6, 0, 99})

	input := load_inputfile()
	input[1] = 12
	input[2] = 2
	solve(input)

	utils.AssertEq(input[0], 3706713)

	return input[0]
}

func part2() int {
	input := load_inputfile()

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			values := utils.CloneIntArray(input)
			values[1] = noun
			values[2] = verb

			solve(values)

			if values[0] == 19690720 {
				result := noun*100 + verb
				utils.AssertEq(result, 8609)

				return result
			}
		}
	}

	return 0
}

func main() {
	fmt.Printf("Part1: %d\n", part1())
	fmt.Printf("Part2: %d\n", part2())
}
