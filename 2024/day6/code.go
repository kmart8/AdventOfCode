package main

import (
	"bufio"
	"fmt"
	"os"
)

func valid(x int, y int) bool {
	return (x >= 0 && x < 130) && (y >= 0 && y < 130)
}

// returns matrix & starting coordinate
func buildMatrix() ([130][]string, int, int) {
	file, _ := os.Open("./inputs/day6.txt")
	scanner := bufio.NewScanner(file)
	matrix := [130][]string{}
	x := 0
	y := 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, v := range line {
			if v == '^' {
				y, x = i, j
			}
			matrix[i] = append(matrix[i], string(v))
		}
		i++
	}
	return matrix, x, y
}

func p1() int {
	matrix, x, y := buildMatrix()
	d := []int{0, -1}
	seen := make(map[string]bool)
	newX, newY := x, y

	for valid(newX, newY) {
		if matrix[newY][newX] == "." || matrix[newY][newX] == "^" {
			seen[fmt.Sprintf("(%d, %d)", newX, newY)] = true
		} else if matrix[newY][newX] == "#" {
			newX, newY = newX-d[0], newY-d[1]
			d[0], d[1] = -d[1], d[0]
		}
		newX, newY = newX+d[0], newY+d[1]
	}

	keys := make([]string, 0, len(seen))
	for k := range seen {
		keys = append(keys, k)
	}

	return len(keys)
}

func p2() int {
	matrix, x, y := buildMatrix()
	ans := 0

	for j, row := range matrix {
		for i := range row {
			d := []int{0, -1}
			seen := make(map[string]bool)
			newX, newY := x, y
			for valid(newX, newY) {
				if seen[fmt.Sprintf("(%d, %d, %d, %d)", newX, newY, d[0], d[1])] {
					ans += 1
					break
				}
				if (matrix[newY][newX] == "." || matrix[newY][newX] == "^") && (newX != i && newY != j) {
					seen[fmt.Sprintf("(%d, %d, %d, %d)", newX, newY, d[0], d[1])] = true
				} else if (matrix[newY][newX] == "#") || (newX == i && newY == j) {
					seen[fmt.Sprintf("(%d, %d, %d, %d)", newX, newY, d[0], d[1])] = true
					newX, newY = newX-d[0], newY-d[1]
					d[0], d[1] = -d[1], d[0]
				}
				newX, newY = newX+d[0], newY+d[1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(fmt.Sprintf("The answer to day 6 part 1 is %d", p1()))
	fmt.Println(fmt.Sprintf("The answer to day 6 part 2 is %d", p2()))
}
