package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildArr() ([]int, []string) {
	file, _ := os.ReadFile("./inputs/day9.txt")
	str := string(file)
	stra := strings.Split(str, "")
	arr := []int{}
	k := 0
	for i, v := range stra {
		n, _ := strconv.Atoi(v)
		if i%2 == 0 {
			for range n {
				arr = append(arr, k)
			}
			k++
		} else {
			for range n {
				arr = append(arr, -1)
			}
		}
	}
	return arr, stra
}

func fragment() []int {
	arr, _ := buildArr()
	l := 0
	r := len(arr) - 1

	for l != r {
		if arr[l] == -1 && arr[r] != -1 {
			arr[r], arr[l] = arr[l], arr[r]
		} else if arr[l] != -1 {
			l++
		} else if arr[r] == -1 {
			r--
		}
	}
	return arr
}

func p1() int {
	arr := fragment()
	return checksum(arr)
}

func p2() int {
	arr, str := buildArr()
	r := len(str) - 1
	for r > 0 {
		size, _ := strconv.Atoi(str[r])
		s2 := []int{}
		for range size {
			s2 = append(s2, -1)
		}
		for i, v := range arr {
			if v == r/2 {
				break
			}
			if areSlicesEqual(arr[i:i+size], s2) {
				for i, v := range arr {
					if v == r/2 {
						arr[i] = -1
					}
				}
				for j := range size {
					arr[i+j] = r / 2
				}
				break
			}
		}
		r -= 2
	}
	return checksum(arr)
}

func areSlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func checksum(arr []int) int {
	s := 0
	for i, v := range arr {
		if v == -1 {
			continue
		}
		s += i * v
	}
	return s
}

func main() {
	fmt.Println(fmt.Sprintf("The answer to day 9 part 1 is %d", p1()))
	fmt.Println(fmt.Sprintf("The answer to day 9 part 2 is %d", p2()))
}
