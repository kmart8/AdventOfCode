package main

import (
	"bufio"
	"fmt"
	"os"
)

type Robot struct {
	px int
	py int
	vx int
	vy int
}

func mod(n int, m int) int {
	return (((n % m) + m) % m)
}

func (r *Robot) step() {
	r.px = mod((r.px + r.vx), 103)
	r.py = mod((r.py + r.vy), 101)
}

func quadSum(matrix [103][101]int) int {
	s1, s2, s3, s4 := 0, 0, 0, 0
	ax1 := (len(matrix)+1)/2 - 1
	ax2 := (len(matrix[0])+1)/2 - 1
	for i, row := range matrix {
		for j, v := range row {
			// q1
			if i < ax1 && j < ax2 {
				s1 += v
				// q2
			} else if i < ax1 && j > ax2 {
				s2 += v
				// q3
			} else if i > ax1 && j < ax2 {
				s3 += v
				// q4
			} else if i > ax1 && j > ax2 {
				s4 += v
			}
		}
	}
	return s1 * s2 * s3 * s4
}

// problem defines x as right & y as down, but I'm using x as down & y as right
func p1() int {
	file, _ := os.Open("./inputs/day14.txt")
	scanner := bufio.NewScanner(file)
	matrix := [103][101]int{}
	for scanner.Scan() {
		robot := Robot{}
		line := scanner.Text()
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.py, &robot.px, &robot.vy, &robot.vx)
		for range 100 {
			robot.step()
		}
		matrix[robot.px][robot.py] += 1
	}
	return quadSum(matrix)
}

func p2() int {
	file, _ := os.Open("./inputs/day14.txt")
	scanner := bufio.NewScanner(file)
	matrix := [103][101]string{}
	robots := []Robot{}
	for scanner.Scan() {
		robot := Robot{}
		line := scanner.Text()
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.py, &robot.px, &robot.vy, &robot.vx)
		robots = append(robots, robot)
	}
	for seconds := range 100000 {
		// simulate
		for i, r := range robots {
			matrix[r.px][r.py] = ""
			r.step()
			robots[i] = r
			matrix[r.px][r.py] = "X"
		}
		for _, row := range matrix {
			for j := range len(row) - 15 {
				if tree(row[j : j+15]) {
					prettyPrint(matrix)
					return seconds + 1
				}
			}
		}
	}
	return -1
}

func prettyPrint(m [103][101]string) {
	for _, row := range m {
		for _, v := range row {
			if v == "X" {
				fmt.Printf("%s", v)
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Print("\n")
	}
}
func tree(s1 []string) bool {
	for _, char := range s1 {
		if char != "X" {
			return false
		}
	}
	return true
}

func main() {
	ans1 := p1()
	ans2 := p2()
	fmt.Println(fmt.Sprintf("The answer to day 14 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 14 part 2 is %d", ans2))
}
