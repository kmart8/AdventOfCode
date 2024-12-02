package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// making new slice every time so as to not modify original
func excludeIndex(slice []string, index int) []string {
	result := make([]string, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func isSafe(report []string) bool {
	dir := 0
	for i := 1; i < len(report); i++ {
		rI, _ := strconv.Atoi(report[i])
		rI1, _ := strconv.Atoi(report[i-1])

		if rI == rI1 {
			break
		} else if absInt(rI-rI1) > 3 {
			break
		} else if rI > rI1 {
			dir++
		} else if rI < rI1 {
			dir--
		}
	}
	if absInt(dir)+1 == len(report) {
		return true
	}
	return false
}

func day2() (int, int) {
	file, _ := os.Open("./inputs/day2.txt")
	scanner := bufio.NewScanner(file)
	safep1 := 0
	safep2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
		if isSafe(report) {
			safep1 += 1
		}
		// try every mutation of report & if any are good, add to safep2 / break
		for i := range len(report) {
			sliceReport := excludeIndex(report, i)
			if isSafe(sliceReport) {
				safep2 += 1
				break
			}
		}
	}
	return safep1, safep2
}

func main() {
	ans1, ans2 := day2()
	fmt.Println(fmt.Sprintf("The answer to day 2 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 2 part 2 is %d", ans2))
}
