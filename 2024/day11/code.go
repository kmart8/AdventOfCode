package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildArr() []string {
	file, _ := os.ReadFile("./inputs/day11.txt")
	arr := strings.Split(string(file), " ")
	return arr
}

func blink(arr []string) []string {
	ans := []string{}
	for _, v := range arr {
		// if stone is '0', change it to '1'
		if v == "0" {
			ans = append(ans, "1")
			// if stone has even # of digits, split
		} else if len(v)%2 == 0 {
			ans = append(ans, v[0:len(v)/2])
			test0, _ := strconv.Atoi(v[len(v)/2:])
			ans = append(ans, fmt.Sprint(test0))
			// if nothing else, m
		} else {
			in, _ := strconv.Atoi(v)
			in = in * 2024
			ans = append(ans, fmt.Sprint(in))
		}
	}
	return ans
}

func blinkBetter(d map[int]int) map[int]int {
	newD := make(map[int]int)
	for key := range d {
		l := len(fmt.Sprint(key))
		if key == 0 {
			newD[1] += d[key]
		} else if l%2 == 0 {
			nL, _ := strconv.Atoi(fmt.Sprint(key)[0 : l/2])
			nR, _ := strconv.Atoi(fmt.Sprint(key)[l/2:])
			newD[nL] += d[key]
			newD[nR] += d[key]
		} else {
			newD[2024*key] += d[key]
		}
	}
	return newD
}

func day11() (int, int) {
	arr := buildArr()
	d := make(map[int]int)
	for _, v := range arr {
		in, _ := strconv.Atoi(v)
		d[in] += 1
	}
	for range 25 {
		arr = blink(arr)
	}
	ans1 := len(arr)
	for range 75 {
		d = blinkBetter(d)
	}
	s := 0
	for key := range d {
		s += d[key]
	}
	ans2 := s

	return ans1, ans2
}

func main() {
	ans1, ans2 := day11()
	fmt.Println(fmt.Sprintf("The answer to day 11 part 1 is %d", ans1))
	fmt.Println(fmt.Sprintf("The answer to day 11 part 2 is %d", ans2))
}
