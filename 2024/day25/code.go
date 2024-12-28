package main

import (
	"bufio"
	"fmt"
	"os"
)

func p1() int {
	file, _ := os.Open("./inputs/day25.txt")
	scanner := bufio.NewScanner(file)
	locks := [][5]int{}
	keys := [][5]int{}
	new := true
	l := false
	c := make(map[int]int)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		if new && string(line[0]) == "#" {
			l = true
			new = false
		} else if new && string(line[0]) == "." {
			l = false
			new = false
		} else if len(line) == 0 {
			arr := [5]int{}
			for k := range c {
				arr[k] = c[k]
			}
			if l {
				for i := range arr {
					arr[i] -= 1
				}
				locks = append(locks, arr)
			} else {
				for i := range arr {
					arr[i] -= 1
				}
				keys = append(keys, arr)
			}
			new = true
			c = make(map[int]int)
		}
		for i, v := range line {
			if string(v) == "#" {
				c[i] += 1
			}
		}
	}
	for _, key := range keys {
		for _, lock := range locks {

			j := 0
			for i := range key {
				if key[i]+lock[i] <= 5 {
					j += 1
				}
			}
			if j == 5 {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(p1())
}
