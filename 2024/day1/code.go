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

func parseInput() ([1000]int, [1000]int) {
	file, _ := os.Open("./inputs/day1.txt")

	// initialize 2 arrays
	col1 := [1000]int{}
	col2 := [1000]int{}
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		col1[i] = num1
		col2[i] = num2
		i += 1
	}
	return col1, col2
}

func part1() int {
	col1, col2 := parseInput()

	slices.Sort(col1[:])
	slices.Sort(col2[:])

	sum := 0

	for i := 0; i < len(col1); i++ {
		sum += absInt(col2[i] - col1[i])
	}

	return sum
}

func part2() int {
	col1, col2 := parseInput()
	frequency := make(map[int]int)
	simScore := 0

	for _, value := range col2 {
		frequency[value]++
	}

	for _, value := range col1 {
		simScore += value * frequency[value]
	}
	return simScore

}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
