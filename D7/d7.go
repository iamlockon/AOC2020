package D7

import (
	"fmt"
	"os"
	"path"
	"bufio"
	"strings"
	"unicode"
	"strconv"
)

var graph map[string][]string

func dfs(cur string, res *int, vis map[string]bool) {
	vis[cur] = true
	for _, adj := range graph[cur] {
		if !vis[adj] {
			*res = *res + 1
			dfs(adj, res, vis)
		}
	}
}


func dfs2(cur string, mp map[string]map[string]int) int {
	// fmt.Println(cur, mp[cur])
	ret := 0
	if mp[cur] == nil {
		// fmt.Println(cur, " contains no bags ")
	}

	for k, v := range mp[cur] {
		ret += v * dfs2(k, mp)
	}
	// fmt.Println(cur, ret)
	return ret + 1
}



// shiny chartreuse bags contain 5 mirrored cyan bags, 3 posh chartreuse bags, 4 dotted aqua bags.

// A <- {B, C, D}
func Part1() {
	fmt.Println("===DAY 7===")
	graph = map[string][]string{}
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D7/d7.txt"))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if strings.Contains(val, "no other") {
			continue
		} else {
			index := strings.Index(val, " bags")
			a := val[:index]
			val = val[index+6:]
			for len(val) > 0 {
				if unicode.IsDigit(rune(val[0])) {
					nidx := strings.Index(val, " bag")
					if graph[val[2:nidx]] == nil {
						graph[val[2:nidx]] = []string{}
					}
					graph[val[2:nidx]] = append(graph[val[2:nidx]], a)
					val = val[nidx+4:]
					continue
				}
				val = val[1:]
			}
		}
	}
	// dfs
	cnt := 0
	vis := map[string]bool{}
	dfs("shiny gold", &cnt, vis)
	fmt.Println(cnt)
}

func Part2() {
	graph := map[string]map[string]int{}
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D7/d7.txt"))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if strings.Contains(val, "no other") {
			continue
		} else {
			index := strings.Index(val, " bags")
			a := val[:index]
			val = val[index+6:]
			for len(val) > 0 {
				if unicode.IsDigit(rune(val[0])) {
					nidx := strings.Index(val, " bag")
					if graph[a] == nil {
						graph[a] = map[string]int{}
					}
					v, _ := strconv.Atoi(string(val[0]));
					graph[a][val[2:nidx]] = v;
					val = val[nidx+4:]
					continue
				}
				val = val[1:]
			}
		}
	}

	fmt.Println(dfs2("shiny gold", graph)-1)
	
}