package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concatNums(a, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	combined, _ := strconv.Atoi(aStr + bStr)
	return combined
}

func valid(arr []int, target int, p2 bool) int {
	operatorCombinations := [][]rune{{}}
	for i := 0; i < (len(arr) - 1); i++ {
		var newResults [][]rune
		for _, combination := range operatorCombinations {
			add := append([]rune(nil), combination...)
			add = append(add, '+')
			newResults = append(newResults, add)

			multi := append([]rune(nil), combination...)
			multi = append(multi, '*')
			newResults = append(newResults, multi)
			if p2 {
				comb := append([]rune(nil), combination...)
				comb = append(comb, '|')
				newResults = append(newResults, comb)
			}
		}
		operatorCombinations = newResults
	}

	for _, ops := range operatorCombinations {
		result := arr[0]
		for i, op := range ops {
			switch string(op) {
			case "+":
				result += arr[i+1]
			case "*":
				result *= arr[i+1]
			case "|":
				result = concatNums(result, arr[i+1])
			}
		}
		if result == target {
			return result
		}
	}
	return 0
}

func day7() (int, int) {
	file, _ := os.Open("./inputs/day7.txt")
	scanner := bufio.NewScanner(file)
	ans1 := 0
	ans2 := 0
	for scanner.Scan() {
		arr := []int{}
		line := scanner.Text()
		fields := strings.Split(line, ":")
		target, _ := strconv.Atoi(string(fields[0]))
		values := strings.Split(fields[1][1:], " ")
		for _, v := range values {
			v, _ := strconv.Atoi(string(v))
			arr = append(arr, v)
		}
		ans1 += valid(arr, target, false)
		ans2 += valid(arr, target, true)
	}
	return ans1, ans2
}

func main() {
	p1, p2 := day7()
	fmt.Println(fmt.Sprintf("The answer to day 7 part 1 is %d", p1))
	fmt.Println(fmt.Sprintf("The answer to day 7 part 2 is %d", p2))
}
