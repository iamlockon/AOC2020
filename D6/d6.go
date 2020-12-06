package D6


import (
	"fmt"
	"os"
	"path"
	"bufio"
)


func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D6/d6.txt"))
	scanner := bufio.NewScanner(file)
	cnt := 0 // all
	ans := map[rune]bool{} // each
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			cnt += len(ans)
			// fmt.Println(len(ans))
			ans = map[rune]bool{}
		} else {
			for _, c := range val {
				ans[c] = true
			}
		}
	}
	cnt += len(ans)
	fmt.Println(cnt)
}

func Part2() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D6/d6.txt"))
	scanner := bufio.NewScanner(file)
	cnt := 0 // all
	ans := map[rune]int{} // each
	pplcnt := 0 // group people count
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			for _, v := range ans {
				if v == pplcnt {
					cnt++
				}
			}
			// fmt.Println(len(ans))
			ans = map[rune]int{}
			pplcnt = 0
		} else {
			for _, c := range val {
				ans[c]++
			}
			pplcnt++
		}
	}
	for _, v := range ans {
		if v == pplcnt {
			cnt++
		}
	}
	fmt.Println(cnt)
}