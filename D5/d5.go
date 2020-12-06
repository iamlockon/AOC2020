package D5

import (
	"fmt"
	"os"
	"path"
	"bufio"
)

// BFBFBBBRLL

type Seat struct {
	Row int
	Col int
}

const MAX_ID_PLUS_ONE = 128*8

var seats []Seat
var exist [MAX_ID_PLUS_ONE]bool

func GetRow(r string) int {
	lo, hi := 0, 127
	for cur := 0; cur < 7; cur++ {
		mid := (lo + hi) / 2
		if string(r[cur]) == "B" {
			lo = mid+1
		} else {
			hi = mid
		}
	}
	return lo
}

func GetCol(c string) int {
	lo, hi := 0, 7
	for cur := 0; cur < 3; cur++ {
		mid := (lo + hi) / 2
		if string(c[cur]) == "R" {
			lo = mid+1
		} else {
			hi = mid
		}
	}
	return lo
}

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "/D5/d5.txt"))
	scanner := bufio.NewScanner(file)
	mx := 0
	mn := 9999
	for scanner.Scan() {
		val := scanner.Text()
		// val := "BFFFBBFRRR"
		sr, sc := val[:7], val[7:]
		r, c := GetRow(sr), GetCol(sc)
		fmt.Println(r, c)
		if exist[r*8+c] == true {
			fmt.Println("repeat:", r*8+c)
		}
		exist[r*8+c] = true
		if r*8+c > mx {
			mx = r*8+c
		}
		if r*8+c < mn {
			mn = r*8+c
		}
		seats = append(seats, Seat{Row: r, Col: c})
	}
	fmt.Println(mx, mn)
}


func Part2() {
	fmt.Println("===========")
	for id := 6; id < 822; id++ {
		if exist[id] == false && exist[id-1] == true && exist[id+1] == true {
			fmt.Println(id)
		}
	}
}