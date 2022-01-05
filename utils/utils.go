package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func AssertEq(a, b int) {
	if a != b {
		panic(fmt.Sprintf("%d != %d", a, b))
	}
}

func AssertIntArrayEq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func CloneIntArray(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)

	return b
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func LoadIntFile(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		out = append(out, v)
	}

	return out
}

func LoadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}
