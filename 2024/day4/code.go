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

func p1() int {
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
		for x, v := range row {
			if v == "X" {
				for _, coord := range directions {
					dx := coord[0]
					dy := coord[1]
					newX, newY := x+dx, y+dy
					if valid(newX, newY) && matrix[newY][newX] == "M" {
						newX, newY := newX+dx, newY+dy
						if valid(newX, newY) && matrix[newY][newX] == "A" {
							newX, newY := newX+dx, newY+dy
							if valid(newX, newY) && matrix[newY][newX] == "S" {
								ans += 1
							}
						}
					}
				}
			}
		}
	}

	return ans

}

func p2() int {
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
				curr1 := []string{}
				curr2 := []string{}
				for _, coord := range direction1 {
					dx := coord[0]
					dy := coord[1]
					newX, newY := x+dx, y+dy
					if valid(newX, newY) && (matrix[newY][newX] == "M" || matrix[newY][newX] == "S") {
						curr1 = append(curr1, matrix[newY][newX])
					}
				}
				for _, coord := range direction2 {
					dx := coord[0]
					dy := coord[1]
					newX, newY := x+dx, y+dy
					if valid(newX, newY) && (matrix[newY][newX] == "M" || matrix[newY][newX] == "S") {
						curr2 = append(curr2, matrix[newY][newX])
					}
				}
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
