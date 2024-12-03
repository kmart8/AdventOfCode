package main

import (
	"fmt"
	"os"
	"regexp"
)

func p1() int {
	file, _ := os.ReadFile("./inputs/day3.txt")
	strs := regexp.MustCompile(`mul\(\d+,\d+\)`).FindAllString(string(file), -1)
	ans := 0
	for _, v := range strs {
		var a, b int
		fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
		ans += a * b
	}
	return ans
}

func p2() int {
	file, _ := os.ReadFile("./inputs/day3.txt")
	strs := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(string(file), -1)
	multiply := true
	ans := 0
	for _, v := range strs {
		if v == "do()" {
			multiply = true
		} else if v == "don't()" {
			multiply = false
		} else {
			if multiply {
				var a, b int
				fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
				ans += a * b
			}
		}
	}
	return ans
}

func main() {
	ans1 := p1()
	ans2 := p2()
	fmt.Println(fmt.Sprintf("The answer to day 3 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 3 part 2 is %d", ans2))
}
