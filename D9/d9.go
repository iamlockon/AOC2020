package D9

import (
	"fmt"
	"os"
	"path"
	"bufio"
	"strconv"
)

var nums []int64
const (
	PREAMBLE_SZ = 25
)

var res int64
func Part1() {
	fmt.Println("===DAY9===")
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D9/d9.txt"))
	scanner := bufio.NewScanner(file)
	mp := map[int64]int{}
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if len(nums) >= PREAMBLE_SZ {
			ok := false
			for _, num := range nums[len(nums)-PREAMBLE_SZ:] {
				if mp[i-num] > 0 && i-num != num{
					ok = true
					break
				}
			}
			if !ok {
				fmt.Println(i)
				res = i
				return
			}
			mp[nums[len(nums)-PREAMBLE_SZ]]--
			if mp[nums[len(nums)-PREAMBLE_SZ]] <= 0 {
				delete(mp, nums[len(nums)-PREAMBLE_SZ])
			}
		}
		mp[i]++
		nums = append(nums, i)
		// fmt.Println(mp)
	}
}

func Part2() {
	for i := 0; i < len(nums)-1; i++ {
		sm := nums[i]
		for j := i+1; j < len(nums); j++ {
			if nums[j] >= res {
				break
			}
			sm += nums[j]
			if sm == res {
				mn,mx := res, int64(0)
				// fmt.Println(i, j)
				for _, k := range nums[i:j+1] {
					if mx < k {
						mx = k
					}
					if mn > k {
						mn = k
					}
				}
				fmt.Println(mn+mx)
				return
			}
		}
	}
}