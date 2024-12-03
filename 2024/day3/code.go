package main

import (
	"fmt"
	"os"
	"regexp"
)

func day3() (int, int) {
	file, _ := os.ReadFile("./inputs/day3.txt")
	strs := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(string(file), -1)
	multiply := true
	ans1 := 0
	ans2 := 0
	for _, v := range strs {
		if v == "do()" {
			multiply = true
		} else if v == "don't()" {
			multiply = false
		} else {
			var a, b int
			fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
			ans1 += a * b
			if multiply {
				ans2 += a * b
			}
		}
	}
	return ans1, ans2
}

func main() {
	ans1, ans2 := day3()
	fmt.Println(fmt.Sprintf("The answer to day 3 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 3 part 2 is %d", ans2))
}
