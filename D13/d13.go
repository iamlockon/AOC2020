package D13

import (
	"os"
	"path"
	"bufio"
	"strconv"
	"fmt"
	"strings"
)

var bus []string

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D13/d13.txt"))
	scanner := bufio.NewScanner(file)
	var ts int 
	if scanner.Scan() {
		ts, _ = strconv.Atoi(scanner.Text())
	}
	if scanner.Scan() {
		avail := scanner.Text()
		bus = strings.Split(avail, ",")
	}
	mn := 999999999999
	nbus := 0
	for _, b := range bus {
		if b == "x"{
			continue
		}
		nb, _ := strconv.Atoi(b)
		earliest := (ts/nb+1)*nb
		// fmt.Println(earliest)
		if earliest < mn {
			mn = earliest
			nbus = nb
		}
	}
	fmt.Println(nbus * (mn-ts))
}

func Part2() {
	// b1 : b1, 2b1, ...
	// b2 : b2, 2b2, ...
	// b1 | n*b1 && b2 | n*b1+idx(b2) && b3 | n*b1+idx(b3) ...

	bs := map[int]int{}
	for idx, b := range bus {
		if b != "x"{
			nb, _ := strconv.Atoi(b)
			bs[nb] = idx
		}
	}
	ans := 0
	product := 1 // running product
	for b, offset := range bs {
		for (ans+offset) % b != 0 { // while not divisible
			ans += product // add one product to current ts
		}
		product *= b // update running product
	}
	fmt.Println(ans)
}