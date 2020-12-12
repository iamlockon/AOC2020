package D10

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l List) Less(i, j int) bool { return l[i] < l[j] }

var list List
var mx int

func Part1() {
	fmt.Println("===DAY10===")
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D10/d10.txt"))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		list = append(list, val)
	}
	sort.Sort(list)
	mx = list[len(list)-1]
	mp := map[int]int{}
	cur := 0
	for i := 0; i < len(list); i++ {
		mp[list[i]-cur]++
		cur = list[i]
	}
	mp[3]++
	fmt.Println(mp[1] * mp[3])
}   

func Part2() {
	// from 0 to mx, how many sequences exist?
	list = append(list, mx, 0)
	copy(list[1:], list)
	list[0] = 0
	exist := map[int]bool{}
	for _,num := range list {
		exist[num] = true
	}
	dp := make([]int64, mx+1)
	dp[0] = 1
	for i := 1; i <= mx; i++ {
		if !exist[i] {
			continue
		}
		if i-3 > -1 {
			dp[i] += dp[i-3]
		}
		if i-2 > -1 {
			dp[i] += dp[i-2]
		}
		dp[i] += dp[i-1]
	} 
	fmt.Println(dp[mx])
}