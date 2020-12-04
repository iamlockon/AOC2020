package D3

import (
	"fmt"
	"os"
	"bufio"
	"path"
)

var graph []string

const (
	TREE = "#"
	SQ = "."
)

type Slope struct {
	R int;
	D int;
}

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "/D3/d3.txt"))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	row, col, cnt := 0, 0, 0
	for scanner.Scan() {
		val := scanner.Text()
		if row > 0 {
			if string(val[col]) == TREE {
				cnt++
			}
		}
		row, col = row+1, (col+3)%len(val)
		graph = append(graph, val)
	}
	fmt.Println(cnt)
}


func Part2() {
	var res = 1
	slopes := []Slope {
		{R: 1, D: 1},
		{R: 3, D: 1},
		{R: 5, D: 1},
		{R: 7, D: 1},
		{R: 1, D: 2},
	}
	for _, slope := range slopes {
		c, r, cnt := slope.R, slope.D, 0
		for ; r < len(graph); r = r+slope.D {
			if string(graph[r][c]) == TREE {
				cnt++
			}
			c = (c+slope.R)%len(graph[0])
		}
		res *= cnt
	}
	fmt.Println(res)
}