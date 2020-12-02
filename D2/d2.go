package D2

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

type Password struct {
	MinL int;
	MaxL int;
	Char rune;
	Password string;
}

var list []Password

func Part1() {
	cwd, _ := os.Getwd()
	file, err := os.Open(path.Join(cwd,"/D2/d2.txt"))
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	scanner := bufio.NewScanner(file)
	var res = 0
	for scanner.Scan() {
		val := scanner.Text()
		var mn, mx int
		var ch rune
		var passwd string
		fmt.Sscanf(val, "%d-%d %c: %s", &mn, &mx, &ch, &passwd)
		fmt.Println(mn, mx, ch, passwd)
		
		var cnt = 0
		for _, c := range passwd {
			if c == ch {
				cnt++
			}
		}
		if mn <= cnt && cnt <= mx {
			res++
		} 

		list = append(list, Password{
			MinL: mn,
			MaxL: mx,
			Char: ch,
			Password: passwd,
		})
	}
	fmt.Println(res) 
}

func Part2() {
	var res = 0
	for _, password := range list {
		first, second, ch, pass := password.MinL, password.MaxL, password.Char, password.Password
		A, B := rune(pass[first-1]) == ch, rune(pass[second-1]) == ch
		if A != B {
			res++
		}
	}
	fmt.Println(res)
}