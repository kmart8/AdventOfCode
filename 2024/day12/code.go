package main

import (
	"bufio"
	"fmt"
	"os"
)

func valid(x int, y int, sizex int, sizey int) bool {
	return (x >= 0 && x < sizex) && (y >= 0 && y < sizey)
}

type Position struct {
	x int
	y int
}

func validEqual(p1 Position, p2 Position, matrix [][]string) bool {
	return (valid(p2.x, p2.y, len(matrix), len(matrix[0])) && matrix[p1.x][p1.y] == matrix[p2.x][p2.y])
}

func buildMatrix() [][]string {
	file, _ := os.Open("./inputs/day12.txt")
	scanner := bufio.NewScanner(file)
	matrix := [][]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		newRow := []string{}
		matrix = append(matrix, newRow)
		for _, v := range line {
			matrix[i] = append(matrix[i], string(v))
		}
		i++
	}
	return matrix
}

func bfs(i int, j int, matrix [][]string, seen map[Position]bool) (int, int, int) {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	q := []Position{}
	q = append(q, Position{i, j})
	seen[Position{i, j}] = true
	p := 0
	a := 1
	c := 0
	sizex := len(matrix)
	sizey := len(matrix[0])
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for i := range directions {
			s1 := directions[i%4]
			s2 := directions[(i+1)%4]
			s3 := []int{s1[0] + s2[0], s1[1] + s2[1]}
			nx, ny := node.x+s1[0], node.y+s1[1]
			p2 := Position{node.x + s2[0], node.y + s2[1]}
			p3 := Position{node.x + s3[0], node.y + s3[1]}
			p1 := Position{nx, ny}
			v := matrix[node.x][node.y]
			if (validEqual(node, p1, matrix) &&
				validEqual(node, p2, matrix) &&
				!validEqual(node, p3, matrix)) ||
				(!validEqual(node, p1, matrix) &&
					!validEqual(node, p2, matrix)) {
				c++
			}
			if !valid(nx, ny, sizex, sizey) {
				p++
			}
			if valid(nx, ny, sizex, sizey) &&
				matrix[nx][ny] != v {
				p++
			}
			if valid(nx, ny, sizex, sizey) &&
				matrix[nx][ny] == v &&
				!seen[Position{nx, ny}] {
				q = append(q, Position{nx, ny})
				seen[Position{nx, ny}] = true
				a++
			}
		}
	}
	return a, p, c
}

func day12() (int, int) {
	ans1 := 0
	ans2 := 0
	matrix := buildMatrix()
	a := 0
	p := 0
	c := 0
	seen := make(map[Position]bool)
	for i, row := range matrix {
		for j := range row {
			if !seen[Position{i, j}] {
				a, p, c = bfs(i, j, matrix, seen)
				// part 1, area * perimeter
				ans1 += a * p
				// part 2, area * corners - since # of corners = # of sides
				ans2 += a * c
			}
		}
	}

	return ans1, ans2
}

func main() {
	ans1, ans2 := day12()
	fmt.Println(fmt.Sprintf("The answer to day 12 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 12 part 2 is %d", ans2))
}
