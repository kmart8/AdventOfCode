package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Button struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

func solveEqn(button Button, p2 bool) (int, bool) {
	a, b, c, d, px, py := button.ax, button.ay, button.bx, button.by, button.px, button.py
	if p2 {
		px += 10000000000000
		py += 10000000000000
	}
	D := a*d - b*c

	if D == 0 {
		return 0, true
	}

	Dx := px*d - c*py
	Dy := a*py - b*px

	if Dx%D != 0 {
		return 0, false
	}
	if Dy%D != 0 {
		return 0, false
	}

	aC := Dx / D
	bC := Dy / D
	if p2 {
		return (3 * aC) + bC, false
	}
	if aC > 100 || bC > 100 {
		return 0, false
	} else {
		return (3 * aC) + bC, false
	}
}

func part1() (int, int) {
	file, _ := os.Open("./inputs/day13.txt")
	scanner := bufio.NewScanner(file)
	button := Button{}
	var tCost1 int = 0
	var tCost2 int = 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "A") {
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &button.ax, &button.ay)
		} else if strings.Contains(line, "B") {
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &button.bx, &button.by)
		} else if strings.Contains(line, "Prize") {
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &button.px, &button.py)
		} else {
			cost1, err1 := solveEqn(button, false)
			cost2, err2 := solveEqn(button, true)
			button = Button{}
			if !err1 {
				tCost1 += cost1
			}
			if !err2 {
				tCost2 += cost2
			}
		}
	}
	// final iteration
	cost1, err1 := solveEqn(button, false)
	cost2, err2 := solveEqn(button, true)
	if !err1 {
		tCost1 += cost1
	}
	if !err2 {
		tCost2 += cost2
	}
	return tCost1, tCost2
}

func main() {
	fmt.Println(part1())
}
