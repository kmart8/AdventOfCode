package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func valid(position Position) bool {
	return 0 <= position.x && position.x < 50 && 0 <= position.y && position.y < 50
}

func antinodes(p1 Position, p2 Position, anti map[Position]bool) int {
	dy := absInt(p2.y - p1.y)
	dx := absInt(p2.x - p1.x)
	a2x, a2y, a1x, a1y := 0, 0, 0, 0
	if p2.x > p1.x {
		a2x = p2.x + dx
		a1x = p1.x - dx
	} else {
		a2x = p2.x - dx
		a1x = p1.x + dx
	}
	if p2.y > p1.y {
		a2y = p2.y + dy
		a1y = p1.y - dy
	} else {
		a2y = p2.y - dy
		a1y = p1.y + dy
	}

	a1 := Position{a1x, a1y}
	a2 := Position{a2x, a2y}

	ans := 0
	if valid(a1) && !anti[a1] {
		anti[a1] = true
		ans += 1
	}
	if valid(a2) && !anti[a2] {
		anti[a2] = true
		ans += 1
	}
	return ans
}

func antiLine(p1 Position, p2 Position, anti map[Position]bool) int {
	dy := absInt(p2.y - p1.y)
	dx := absInt(p2.x - p1.x)
	a := p1
	ans := 0
	d1 := 0
	d2 := 0

	if p2.x > p1.x {
		d1 = -1
	} else {
		d1 = 1
	}
	if p2.y > p1.y {
		d2 = -1
	} else {
		d2 = 1
	}

	for _, v := range []int{-1, 1} {
		a = p1
		for valid(a) {
			if valid(a) && !anti[a] {
				anti[a] = true
				ans += 1
			}
			a = Position{a.x + (v * dx * d1), a.y + (v * dy * d2)}
		}
	}
	return ans
}

func day8() (int, int) {
	file, _ := os.Open("./inputs/day8.txt")
	scanner := bufio.NewScanner(file)
	stations := make(map[string][]Position)
	anti1 := make(map[Position]bool)
	anti2 := make(map[Position]bool)
	i := 0
	ans1 := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "")
		for j, v := range points {
			if v != "." {
				currPos := Position{i, j}
				for _, pos := range stations[v] {
					ans1 += antinodes(currPos, pos, anti1)
					ans2 += antiLine(currPos, pos, anti2)
				}
				stations[v] = append(stations[v], currPos)
			}
		}
		i++
	}
	return ans1, ans2
}

func main() {
	ans1, ans2 := day8()
	fmt.Println(fmt.Sprintf("The answer to day 8 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 8 part 2 is %d", ans2))
}
