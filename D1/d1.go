package D1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"path"
)

var MAGIC_NUMBER = 2020

var nums []int
var mp map[int]bool

func init() {
	nums = make([]int, 0)
	mp = make(map[int]bool)
}

// Part1
func Part1() {
	s, _ :=  os.Getwd()
	file, err := os.Open(path.Join(s,"D1/d1.txt"))
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "read file: ", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		if _, ok := mp[MAGIC_NUMBER-val]; ok {
			fmt.Println((MAGIC_NUMBER - val) * val)
		}
		mp[val] = true
		nums = append(nums, val)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "stdin:", err)
	}
	return
}

// Part2
func Part2() {
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if _, ok := mp[MAGIC_NUMBER-nums[i]-nums[j]]; ok {
				fmt.Println(nums[i]*nums[j]*(MAGIC_NUMBER-nums[i]-nums[j]))
				return
			}
		}
	}
}
