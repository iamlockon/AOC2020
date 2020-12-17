package D14

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

// mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
// mem[8] = 11

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D14/d14.txt"))
	scanner := bufio.NewScanner(file)
	mask := ""
	
	mem := map[int64]int64{}
	for scanner.Scan() {
		val := scanner.Text()
		if val[:3] == "mem" {
			idx, _ :=  strconv.Atoi(val[strings.Index(val, "[")+1:strings.Index(val, "]")])
			v, _ := strconv.Atoi(strings.Split(val, " ")[len(strings.Split(val, " "))-1])
			for i,p := len(mask)-1, 1 ; i > -1; i, p = i-1, p*2 {
				if mask[i] == 'X' {
					continue
				}
				if mask[i] == '0' {
					v &= ^p
				} else {
					v |= p
				}
			}
			mem[int64(idx)] = int64(v)
		} else {
			mask = strings.Split(val, " ")[len(strings.Split(val, " "))-1]
		}
	}
	res := int64(0)
	for _, v := range mem {
		if v != 0 {
			res += int64(v)
		}
	}
	fmt.Println(res)
}

func Part2() {
	w := make(chan int)
	go func (){
		cwd, _ := os.Getwd()
		file, _ := os.Open(path.Join(cwd, "D14/d14.txt"))
		scanner := bufio.NewScanner(file)
		mask := ""
		mem := map[int64]int64{}
		for scanner.Scan() {
			val := scanner.Text()
			fmt.Println(val)
			if val[:3] == "mem" {
				idx, _ :=  strconv.Atoi(val[strings.Index(val, "[")+1:strings.Index(val, "]")])
				v, _ := strconv.Atoi(strings.Split(val, " ")[len(strings.Split(val, " "))-1])
				dest := []int64{int64(idx)}
				for i,p := len(mask)-1, int64(1) ; i > -1; i, p = i-1, p*2 {
					if mask[i] == '0' {
						continue
					}
					if mask[i] == '1' {
						for j := 0; j < len(dest); j++ {
							dest[j] |= p
						}
					} else {
						ndest := []int64{}
						for x := 0; x < len(dest); x++ {
							A, B := dest[x] | p, dest[x] & (^p)
							ndest = append(ndest, A, B)
						}
						dest = ndest
						
					}
				}
				for _, i := range dest {
					mem[i] = int64(v)
				}
			} else {
				mask = strings.Split(val, " ")[len(strings.Split(val, " "))-1]
			}
		}
		res := int64(0)
		for _, v := range mem {
			if v != 0 {
				res += v
			}
		}	
		fmt.Println(res)
		w<-1
	}()
	<-w
}