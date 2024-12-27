package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func valid(x int, y int, sizex int, sizey int) bool {
	return (x >= 0 && x < sizex) && (y >= 0 && y < sizey)
}

type Position struct {
	x int
	y int
}

func prettyPrint(m [][]string) {
	for _, row := range m {
		for _, v := range row {
			fmt.Printf("%s", v)
		}
		fmt.Print("\n")
	}
}

// returns matrix & starting position
func buildMatrix(p2 bool) ([][]string, int, int, []string) {
	var file *os.File
	if p2 {
		file, _ = os.Open("./inputs/day152.txt")
	} else {
		file, _ = os.Open("./inputs/day15.txt")
	}
	scanner := bufio.NewScanner(file)
	matrix := [][]string{}
	instructions := []string{}
	i, x, y := 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 105 {
			newRow := []string{}
			matrix = append(matrix, newRow)
			for j, v := range line {
				if v == '@' {
					x, y = i, j
				}
				matrix[i] = append(matrix[i], string(v))
			}
			i++
		} else {
			for _, v := range line {
				instructions = append(instructions, string(v))
			}
		}
	}
	return matrix, x, y, instructions
}

func p1() int {
	dMap := make(map[string]Position)
	dMap["^"] = Position{-1, 0}
	dMap["v"] = Position{1, 0}
	dMap[">"] = Position{0, 1}
	dMap["<"] = Position{0, -1}

	matrix, x, y, instructions := buildMatrix(false)
	p := Position{x, y}
	for _, in := range instructions {
		x, y := push(p, dMap[in], matrix)
		p = Position{x, y}
	}
	return calculateGPS(matrix)
}

func push(p Position, d Position, m [][]string) (int, int) {
	nx, ny := p.x, p.y
	toMove := []Position{p}
	nx, ny = nx+d.x, ny+d.y
	for m[nx][ny] != "#" {
		if valid(nx, ny, len(m), len(m[0])) {
			if m[nx][ny] == "O" {
				toMove = append(toMove, Position{nx, ny})
				nx, ny = nx+d.x, ny+d.y
			} else if m[nx][ny] == "." {
				for i := len(toMove) - 1; i >= 0; i-- {
					n := toMove[i]
					m[n.x][n.y], m[n.x+d.x][n.y+d.y] = m[n.x+d.x][n.y+d.y], m[n.x][n.y]
				}
				return p.x + d.x, p.y + d.y
			}
		}
	}
	return p.x, p.y
}

func calculateGPS(m [][]string) int {
	ans := 0
	for i, row := range m {
		for j, v := range row {
			if v == "O" {
				ans += (100 * i) + j
			}
		}
	}
	return ans
}

func p2() int {
	dMap := make(map[string]Position)
	dMap["^"] = Position{-1, 0}
	dMap["v"] = Position{1, 0}
	dMap[">"] = Position{0, 1}
	dMap["<"] = Position{0, -1}

	matrix, x, y, instructions := buildMatrix(true)
	p := Position{x, y}
	for _, in := range instructions {
		x, y := push2(p, dMap[in], matrix)
		p = Position{x, y}
	}
	return calculateGPS2(matrix)
}

func push2(p Position, d Position, m [][]string) (int, int) {
	q := []Position{}
	seen := []Position{}
	q = append(q, p)
	seen = append(seen, p)
	dMap := make(map[string]string)
	dMap["["] = "]"
	dMap["]"] = "["
	dMap2 := make(map[string]int)
	dMap2["["] = 1
	dMap2["]"] = -1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		nx, ny := node.x+d.x, node.y+d.y
		if valid(nx, ny, len(m), len(m[0])) {
			if m[nx][ny] == "[" || m[nx][ny] == "]" && !slices.Contains(seen, Position{nx, ny}) {
				if d.x == 1 || d.x == -1 {
					v := dMap2[m[nx][ny]]
					if valid(nx, ny+v, len(m), len(m[0])) &&
						(m[nx][ny+v] == dMap[m[nx][ny]]) &&
						!slices.Contains(seen, Position{nx, ny + v}) {
						seen = append(seen, Position{nx, ny + v})
						q = append(q, Position{nx, ny + v})
					}
				}
				if !slices.Contains(seen, Position{nx, ny}) {
					seen = append(seen, Position{nx, ny})
					q = append(q, Position{nx, ny})
				}
			} else if m[nx][ny] == "#" {
				return p.x, p.y
			}
		}
	}
	for i := len(seen) - 1; i >= 0; i-- {
		n := seen[i]
		m[n.x][n.y], m[n.x+d.x][n.y+d.y] = m[n.x+d.x][n.y+d.y], m[n.x][n.y]
	}
	return p.x + d.x, p.y + d.y
}

func calculateGPS2(m [][]string) int {
	ans := 0
	for i, row := range m {
		for j, v := range row {
			if v == "[" {
				ans += (100 * i) + j
			}
		}
	}
	return ans
}

func main() {
	ans1 := p1()
	ans2 := p2()
	fmt.Println(fmt.Sprintf("The answer to day 15 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 15 part 2 is %d", ans2))
}
