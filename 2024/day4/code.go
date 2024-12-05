package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func valid(x int, y int) bool {
	return (x >= 0 && x < 140) && (y >= 0 && y < 140)
}

func buildMatrix() [140][]string {
	file, _ := os.Open("./inputs/day4.txt")
	scanner := bufio.NewScanner(file)
	// build the matrix
	matrix := [140][]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, v := range line {
			matrix[i] = append(matrix[i], string(v))
		}
		i++
	}
	return matrix
}

func hasXmas(x int, y int, dx int, dy int, matrix [140][]string) bool {
	for k, v := range "XMAS" {
		newX, newY := x+(dx*k), y+(dy*k)
		if !valid(newX, newY) || matrix[newY][newX] != string(v) {
			return false
		}
	}
	return true
}

func findXmas(direction [2][]int, x int, y int, matrix [140][]string) []string {
	curr := []string{}
	for _, coord := range direction {
		dx := coord[0]
		dy := coord[1]
		newX, newY := x+dx, y+dy
		if valid(newX, newY) && (matrix[newY][newX] == "M" || matrix[newY][newX] == "S") {
			curr = append(curr, matrix[newY][newX])
		}
	}
	return curr
}

func p1() int {
	matrix := buildMatrix()

	directions := [8][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	ans := 0
	for y, row := range matrix {
		for x := range row {
			for _, coord := range directions {
				dx := coord[0]
				dy := coord[1]
				if hasXmas(x, y, dx, dy, matrix) {
					ans += 1
				}
			}
		}
	}
	return ans
}

func p2() int {
	matrix := buildMatrix()

	// each direction must contain an 'M' & 'S'
	direction1 := [2][]int{
		{-1, -1},
		{1, 1},
	}
	direction2 := [2][]int{
		{-1, 1},
		{1, -1},
	}

	ans := 0
	for y, row := range matrix {
		for x, v := range row {
			if v == "A" {
				curr1 := findXmas(direction1, x, y, matrix)
				curr2 := findXmas(direction2, x, y, matrix)
				if (slices.Contains(curr1, "M") && slices.Contains(curr1, "S")) && (slices.Contains(curr2, "M") && slices.Contains(curr2, "S")) {
					ans += 1
				}
			}
		}
	}
	return ans
}

func main() {
	ans1 := p1()
	ans2 := p2()
	fmt.Println(fmt.Sprintf("The answer to day 4 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 4 part 2 is %d", ans2))
}
