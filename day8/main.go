package main

import (
	"aoc2019/utils"
	"fmt"
	"math"
)

const INPUT_PATH = "day8/input.txt"
const IMG_WIDTH = 25
const IMG_HEIGHT = 6

func main() {
	input := utils.LoadFile(INPUT_PATH)

	layerLen := IMG_WIDTH * IMG_HEIGHT
	layerCount := len(input) / layerLen

	if layerCount*layerLen != len(input) {
		panic("Invalid size")
	}

	// Extract layers
	layers := []string{}

	for i := 0; i < layerCount; i++ {
		layer := input[i*layerLen : (i+1)*layerLen]
		layers = append(layers, layer)
	}

	// Find layer with the lower zero count
	foundLayerIdx := 0
	foundLayerZeroCount := math.MaxInt

	for i := 0; i < layerCount; i++ {
		zeroCount := 0

		for j := 0; j < layerLen; j++ {
			if layers[i][j] == '0' {
				zeroCount++
			}
		}

		if zeroCount < foundLayerZeroCount {
			foundLayerZeroCount = zeroCount
			foundLayerIdx = i
		}
	}

	// Compute part1 result
	oneCount := 0
	twoCount := 0

	for i := 0; i < layerLen; i++ {
		switch layers[foundLayerIdx][i] {
		case '1':
			oneCount++
		case '2':
			twoCount++
		}
	}

	result := oneCount * twoCount
	utils.AssertEq(result, 2440)

	// Part2
	// Prepare output image
	output := [][]byte{}

	for y := 0; y < IMG_HEIGHT; y++ {
		line := []byte{}

		for x := 0; x < IMG_WIDTH; x++ {
			line = append(line, ' ')
		}

		output = append(output, line)
	}

	// Draw image
	for i := 0; i < layerCount; i++ {
		for y := 0; y < IMG_HEIGHT; y++ {
			for x := 0; x < IMG_WIDTH; x++ {
				if output[y][x] != ' ' {
					continue
				}

				idx := y*IMG_WIDTH + x
				switch layers[i][idx] {
				case '0':
					output[y][x] = '.'
				case '1':
					output[y][x] = '#'
				}
			}
		}
	}

	// Should display "AZCJC"
	for y := 0; y < IMG_HEIGHT; y++ {
		for x := 0; x < IMG_WIDTH; x++ {
			fmt.Printf("%c", output[y][x])
		}

		fmt.Println()
	}
}
