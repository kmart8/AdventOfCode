package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func parseInput() ([1000]int, [1000]int, map[int]int) {
	file, _ := os.Open("./inputs/day1.txt")

	col1 := [1000]int{}
	col2 := [1000]int{}
	frequency := make(map[int]int)
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		col1[i] = num1
		col2[i] = num2
		frequency[num2]++
		i += 1
	}
	return col1, col2, frequency
}

func day1() (int, int) {
	col1, col2, frequency := parseInput()

	slices.Sort(col1[:])
	slices.Sort(col2[:])

	sum := 0
	simScore := 0

	for i := 0; i < len(col1); i++ {
		num1, num2 := col1[i], col2[i]
		sum += absInt(num2 - num1)
		simScore += num1 * frequency[num1]
	}

	return sum, simScore
}

func main() {
	ans1, ans2 := day1()
	fmt.Println(fmt.Sprintf("The answer to day 1 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 1 part 2 is %d", ans2))
}
