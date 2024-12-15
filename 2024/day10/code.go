package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func valid(x int, y int, sizex int, sizey int) bool {
	return (x >= 0 && x < sizex) && (y >= 0 && y < sizey)
}

func buildMatrix() [][]int {
	file, _ := os.Open("./inputs/day10.txt")
	scanner := bufio.NewScanner(file)
	matrix := [][]int{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		newRow := []int{}
		matrix = append(matrix, newRow)
		for _, v := range line {
			intV, _ := strconv.Atoi(string(v))
			matrix[i] = append(matrix[i], intV)
		}
		i++
	}
	return matrix
}

type Position struct {
	x int
	y int
}

func bfs(i int, j int, matrix [][]int) int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	q := []Position{}
	seen := make(map[Position]bool)
	q = append(q, Position{i, j})
	seen[Position{i, j}] = true
	ans := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, d := range directions {
			dx, dy := d[0], d[1]
			nx, ny := node.x+dx, node.y+dy
			v := matrix[node.x][node.y]
			if valid(nx, ny, len(matrix), len(matrix[0])) &&
				!seen[Position{nx, ny}] &&
				matrix[nx][ny] == v+1 {
				q = append(q, Position{nx, ny})
				seen[Position{nx, ny}] = true
				if matrix[nx][ny] == 9 {
					ans += 1
				}
			}
		}
	}
	return ans
}

func p1() int {
	matrix := buildMatrix()
	ans := 0
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				ans += bfs(i, j, matrix)
			}
		}
	}
	return ans
}

func dfs(x int, y int, matrix [][]int) int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	v := matrix[x][y]
	if v == 9 {
		return 1
	}
	n := len(matrix)
	ans := 0
	for _, d := range directions {
		dx, dy := d[0], d[1]
		nx, ny := x+dx, y+dy
		if valid(nx, ny, n, n) && matrix[nx][ny] == v+1 {
			ans += dfs(nx, ny, matrix)
		}
	}
	return ans
}

func p2() int {
	matrix := buildMatrix()
	ans := 0
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				ans += dfs(i, j, matrix)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(fmt.Sprintf("The answer to day 10 part 1 is %d", p1()))
	fmt.Println(fmt.Sprintf("The answer to day 10 part 2 is %d", p2()))
}
