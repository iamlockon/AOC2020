package D11

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

const FILE_NAME = "D11/d11.txt"

const (
	FLOOR    = 2
	EMP_SEAT = 0
	OCC_SEAT = 1
)

const (
	X = iota
	XP
	Y
	YP
	FIR
	SEC
	THI
	FOU
)

var revmp = map[int]rune{
	2: '.',
	0: 'L',
	1: '#',
}

var mp = map[rune]int{
	'.': 2,
	'L': 0,
	'#': 1,
}

type Directions struct {
	X   int
	XP  int
	Y   int
	YP  int
	FIR int
	SEC int
	THI int
	FOU int
}

type Cell struct {
	State int
	Next  int
	Dir   Directions
}

var graph [][]*Cell
var R, C int

func Part1() {
	fmt.Println("===DAY11===")
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, FILE_NAME))
	scanner := bufio.NewScanner(file)
	graph = [][]*Cell{}
	for scanner.Scan() {
		val := scanner.Text()
		graph = append(graph, make([]*Cell, len(val)))
		for i, r := range val {
			graph[len(graph)-1][i] = &Cell{State: mp[r], Next: FLOOR}
		}
	}
	R = len(graph)
	C = len(graph[0])

	// iterate
	cc := 0
	for {
		var changed = false
		for i, row := range graph {
			for j, cell := range row {
				if cell.State == FLOOR {
					// fmt.Print(" ")
					continue
				}
				nei := [][2]int{
					{i, j - 1},
					{i, j + 1},
					{i + 1, j - 1},
					{i + 1, j},
					{i + 1, j + 1},
					{i - 1, j - 1},
					{i - 1, j},
					{i - 1, j + 1},
				}
				occ := 0
				for k := 0; k < len(nei); k++ {
					if nei[k][0] < 0 || nei[k][0] >= R || nei[k][1] < 0 || nei[k][1] >= C {
						continue
					}
					if graph[nei[k][0]][nei[k][1]].State == OCC_SEAT {
						occ++
					}
				}
				if cell.State == EMP_SEAT && occ == 0 {
					cell.Next = OCC_SEAT
					changed = true
				}
				if cell.State == OCC_SEAT && occ >= 4 {
					cell.Next = EMP_SEAT
					changed = true
				}
			}
			// fmt.Println()
		}
		if !changed {
			break
		}

		for _, row := range graph {
			for _, cell := range row {
				if cell.Next != FLOOR {
					cell.State = cell.Next
					cell.Next = FLOOR // reset
				}
			}
		}
		cc++
		// if cc > 1 {
		// 	break
		// }
	}

	// fmt.Println()
	// fmt.Println(cc)
	cnt := 0
	for _, row := range graph {
		for _, cell := range row {
			// fmt.Print(cell.State)
			if cell.State == OCC_SEAT {
				cnt++
			}
		}
		// fmt.Println()
	}
	fmt.Println(cnt)
}

// func Part2() {
// 	cwd, _ := os.Getwd()
// 	file, _ := os.Open(path.Join(cwd, "D11/d11_.txt"))
// 	scanner := bufio.NewScanner(file)
// 	graph = [][]*Cell{}
// 	for scanner.Scan() {
// 		val := scanner.Text()
// 		graph = append(graph, make([]*Cell, len(val)))
// 		for i,r := range val {
// 			graph[len(graph)-1][i] = &Cell{State: mp[r], Next: FLOOR}
// 		}
// 	}
// }

func dfs(i, j int, dir int, vis map[int]map[int]map[int]int) {
	if i < 0 || i >= R || j < 0 || j >= C {
		return
	}

	if _, ok := vis[i][j][dir]; ok {
		return
	}

	if vis[i] == nil {
		vis[i] = map[int]map[int]int{}
	}
	if vis[i][j] == nil {
		vis[i][j] = map[int]int{}
	}

	curOcc := 0
	if graph[i][j].State == OCC_SEAT {
		curOcc++
	}
	if graph[i][j].State == EMP_SEAT {
		vis[i][j][dir] = 0
		return
	}

	ret := curOcc
	switch dir {
	case Y:
		dfs(i+1, j, Y, vis)
		vis[i][j][dir] = ret + vis[i+1][j][dir]
		// graph[i][j].Dir.Y = ret
	case YP:
		dfs(i-1, j, YP, vis)
		vis[i][j][dir] = ret + vis[i-1][j][dir]
		// graph[i][j].Dir.YP = ret
	case X:
		dfs(i, j-1, X, vis)
		vis[i][j][dir] = ret + vis[i][j-1][dir]
		// graph[i][j].Dir.YP = ret
	case XP:
		dfs(i, j+1, XP, vis)
		vis[i][j][dir] = ret + vis[i][j+1][dir]
		// graph[i][j].Dir.XP = ret
	case FIR:
		dfs(i-1, j+1, FIR, vis)
		vis[i][j][dir] = ret + vis[i-1][j+1][dir]
		// graph[i][j].Dir.FIR = ret
	case SEC:
		dfs(i-1, j-1, SEC, vis)
		vis[i][j][dir] = ret + vis[i-1][j-1][dir]
		// graph[i][j].Dir.SEC = ret
	case THI:
		dfs(i+1, j-1, THI, vis)
		vis[i][j][dir] = ret + vis[i+1][j-1][dir]
		// graph[i][j].Dir.THI = ret
	case FOU:
		dfs(i+1, j+1, FOU, vis)
		vis[i][j][dir] = ret + vis[i+1][j+1][dir]
		// graph[i][j].Dir.FOU = ret
	}

}

func Part2() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, FILE_NAME))
	scanner := bufio.NewScanner(file)
	graph = [][]*Cell{}
	for scanner.Scan() {
		val := scanner.Text()
		graph = append(graph, make([]*Cell, len(val)))
		for i, r := range val {
			graph[len(graph)-1][i] = &Cell{State: mp[r], Next: FLOOR}
		}
	}

	cc := 0
	for {
		changed := false
		vis := map[int]map[int]map[int]int{}
		for i := 0; i < R; i++ {
			if vis[i] == nil {
				vis[i] = map[int]map[int]int{}
			}
			for j := 0; j < C; j++ {
				// for each direction, get seen occupied

				if vis[i][j] == nil {
					vis[i][j] = map[int]int{}
				}
				cell := graph[i][j]
				dfs(i-1, j, YP, vis)
				dfs(i+1, j, Y, vis)
				dfs(i, j+1, XP, vis)
				dfs(i, j-1, X, vis)
				dfs(i-1, j+1, FIR, vis)
				dfs(i-1, j-1, SEC, vis)
				dfs(i+1, j-1, THI, vis)
				dfs(i+1, j+1, FOU, vis)

				occ := 0
				if vis[i-1][j][YP] > 0 {
					occ++
				}
				if vis[i+1][j][Y] > 0 {
					occ++
				}
				if vis[i][j+1][XP] > 0 {
					occ++
				}
				if vis[i][j-1][X] > 0 {
					occ++
				}
				if vis[i-1][j+1][FIR] > 0 {
					occ++
				}
				if vis[i-1][j-1][SEC] > 0 {
					occ++
				}
				if vis[i+1][j-1][THI] > 0 {
					occ++
				}
				if vis[i+1][j+1][FOU] > 0 {
					occ++
				}
				if cell.State == EMP_SEAT && occ == 0 {
					cell.Next = OCC_SEAT
					changed = true
				}
				if cell.State == OCC_SEAT && occ >= 5 {
					cell.Next = EMP_SEAT
					changed = true
				}

			}
		}
		if !changed {
			break
		}
		cc++
		for _, row := range graph {
			for _, cell := range row {
				if cell.Next != FLOOR {
					cell.State = cell.Next
					cell.Next = FLOOR // reset
				}
			}
		}

		// for _, row := range graph {
		// 	for _, cell := range row {
		// 		fmt.Print(string(revmp[cell.State]))
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
	}

	// fmt.Println()
	// fmt.Println(cc)
	cnt := 0
	for _, row := range graph {
		for _, cell := range row {
			// fmt.Print(cell.State)
			if cell.State == OCC_SEAT {
				cnt++
			}
		}
		// fmt.Println()
	}
	fmt.Println(cnt)
}
