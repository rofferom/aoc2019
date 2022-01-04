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
