package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func intersection(arr1 []int, arr2 []int) []int {
	result := []int{}
	set := make(map[int]bool)

	// Add elements of arr1 to set
	for _, num := range arr1 {
		set[num] = true
	}

	// Check if elements of arr2 exist in set
	for _, num := range arr2 {
		if set[num] {
			result = append(result, num)
			set[num] = false // To avoid duplicates in the result
		}
	}

	return result
}

func day5() (int, int) {
	file, _ := os.Open("./inputs/day5.txt")
	scanner := bufio.NewScanner(file)
	// map of key:value where value is all the numbers that can be after key
	// each iteration, see if value is in set. then take intersection of currSet & new value array, make new currSet
	order := make(map[int][]int)
	matrix := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			fields := strings.Split(line, "|")
			num1, _ := strconv.Atoi(fields[0])
			num2, _ := strconv.Atoi(fields[1])
			order[num1] = append(order[num1], num2)
		} else if strings.Contains(line, ",") {
			fields := strings.Split(line, ",")
			arr := []int{}
			for _, j := range fields {
				j, _ := strconv.Atoi(j)
				arr = append(arr, j)
			}
			matrix = append(matrix, arr)
		}
	}
	ansp1 := 0
	ansp2 := 0

	for _, row := range matrix {
		keys := make([]int, 0, len(order))
		for k := range order {
			keys = append(keys, k)
		}
		for j, value := range row {
			if !slices.Contains(keys, value) {
				break
			} else if j == len(row)-1 {
				ansp1 += row[(len(row)-1)/2]
			}
			keys = intersection(keys, order[value])
		}
	}

	for _, row := range matrix {
		keys := make([]int, 0, len(order))
		for k := range order {
			keys = append(keys, k)
		}
		incorrect := false
		for j, value := range row {
			if !slices.Contains(keys, value) {
				incorrect = true
			}
			if j == len(row)-1 && incorrect {
				row = s(row, order)
				ansp2 += row[(len(row)-1)/2]
			}
			keys = intersection(keys, order[value])
		}
	}
	return ansp1, ansp2
}

func s(row []int, order map[int][]int) []int {
	sorted := false
	for sorted == false {
		swaps := 0
		for i, v := range row {
			if i == 0 {
				continue
			}
			if !slices.Contains(order[row[i-1]], v) {
				row[i-1], row[i] = row[i], row[i-1]
				swaps++
			}
		}
		if swaps == 0 {
			sorted = true
		}
	}
	return row
}

func main() {
	ans1, ans2 := day5()
	fmt.Println(fmt.Sprintf("The answer to day 4 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 4 part 2 is %d", ans2))
}
