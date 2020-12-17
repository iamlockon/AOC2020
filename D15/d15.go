package D15

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"bufio"
	"strings"
)

const (
	WANTED = 30000000
)

var nums []int

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D15/d15.txt"))
	scanner := bufio.NewScanner(file)
	nums = []int{-1} // padding number
	prev_mp := map[int]int{}
	for scanner.Scan() {
		for i, s := range strings.Split(scanner.Text(), ",") {
			v, _ := strconv.Atoi(s)
			prev_mp[v] = i+1
			nums = append(nums, v)
		}
	}

	nth := []int{0}
	var last int
	cur := len(nums)+1
	for cur <= WANTED {
		last = nth[len(nth)-1]
		
		if _, ok := prev_mp[last]; ok {
			// diff
			diff := cur - 1 - prev_mp[last]
			nth = append(nth, diff)
		} else {
			// never spoken
			nth = append(nth, 0)
		}
		prev_mp[last] = cur-1
		// fmt.Println(cur, nth)
		cur++
	}
	// for i := len(nth); i <= 10; i++ { // i denotes turn
	// 	prev := nth[i-1] // previous spoken
	// 	if v, ok := prev_mp[prev]; ok {
	// 		diff := i-1 - v
	// 		nth = append(nth, diff)
	// 		if _, ok := prev_mp[diff]; ok {
	// 			prev_mp[diff] = i
	// 		}
	// 	} else {
	// 		nth = append(nth, 0)
	// 		prev_mp[0] = i
	// 		prev_mp[prev] = i-1
	// 	}

	// 	fmt.Println(nth, prev_mp)
	// } 
	fmt.Println(nth[WANTED-len(nums)])
}


func Part2() {

}