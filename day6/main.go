package main

import (
	"aoc2019/utils"
	"fmt"
	"math"
	"strings"
)

const INPUT_PATH = "day6/input.txt"

type node struct {
	parent string
	childs []string
}

type graph map[string]*node

func parseInput(input string) graph {
	nodes := make(graph)

	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, ")")
		parentName, childName := items[0], items[1]

		// Add/Update parent
		if parent, ok := nodes[parentName]; ok {
			parent.childs = append(parent.childs, childName)
		} else {
			nodes[parentName] = &node{
				parent: "",
				childs: []string{childName},
			}
		}

		if _, ok := nodes[childName]; ok {
			// Link child to parent if it was unknown
			child := nodes[childName]
			if child.parent == "" {
				child.parent = parentName
			}
		} else {
			// Create child if required
			nodes[childName] = &node{
				parent: parentName,
				childs: []string{},
			}
		}
	}

	return nodes
}

func orbitsCount(nodes graph) int {
	type nodeVisitor struct {
		name  string
		depth int
	}

	count := 0
	visitList := []nodeVisitor{{name: "COM", depth: 0}}

	for len(visitList) > 0 {
		currentNode := visitList[0]
		visitList = visitList[1:]

		for _, childName := range nodes[currentNode.name].childs {
			visitList = append(visitList, nodeVisitor{name: childName, depth: currentNode.depth + 1})
		}

		count += currentNode.depth
	}

	return count
}

func distanceToSanta(nodes graph) int {
	// Init distances
	dist := make(map[string]int)

	for name := range nodes {
		if name == "YOU" {
			dist[name] = 0
		} else {
			dist[name] = math.MaxInt
		}
	}

	// List of nodes to visit
	toVisit := []string{"YOU"}

	for {
		// Pop the node to visit: the one with the lowest cost
		// AKA: stupid implementation of a priority queue
		var position string
		visitIdx := -1
		cost := math.MaxInt

		for i, name := range toVisit {
			distance := dist[name]

			if distance < cost {
				cost = distance
				position = name
				visitIdx = i
			}
		}

		utils.AssertNeq(visitIdx, -1)

		toVisit[visitIdx] = toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		// Target reached. Substract 2 because we don't need distances between
		// YOU and SAN and their object.
		if position == "SAN" {
			return cost - 2
		}

		if cost > dist[position] {
			continue
		}

		// Try to find better paths from this node
		edges := append(nodes[position].childs, nodes[position].parent)
		for _, childName := range edges {
			if cost+1 < dist[childName] {
				dist[childName] = cost + 1
				toVisit = append(toVisit, childName)
			}
		}
	}
}

func part1() int {
	example := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	nodes := parseInput(example)
	utils.AssertEq(orbitsCount(nodes), 42)

	nodes = parseInput(utils.LoadFile(INPUT_PATH))
	count := orbitsCount(nodes)
	utils.AssertEq(count, 119831)

	return orbitsCount(nodes)
}

func part2() int {
	example := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	nodes := parseInput(example)
	utils.AssertEq(distanceToSanta(nodes), 4)

	nodes = parseInput(utils.LoadFile(INPUT_PATH))
	count := distanceToSanta(nodes)
	utils.AssertEq(count, 322)

	return count
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
