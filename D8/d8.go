package D8

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	NOP = "nop"
	JMP = "jmp"
	ACC = "acc"
)

type Op struct {
	Type string
	Arg  int
}

type GameConsole struct {
	StackPtr    int
	Accumulator int
}

var ops = []Op{}

func Part1() {
	fmt.Println("===DAY8===")
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D8/d8.txt"))
	scanner := bufio.NewScanner(file)
	vis := map[int]bool{}
	for scanner.Scan() {
		val := scanner.Text()
		v := strings.Split(val, " ")
		op, arg := v[0], v[1]
		iarg, _ := strconv.Atoi(arg)
		ops = append(ops, Op{Type: op, Arg: iarg})
	}
	console := &GameConsole{}
	for vis[console.StackPtr] != true {
		vis[console.StackPtr] = true
		switch t, a := ops[console.StackPtr].Type, ops[console.StackPtr].Arg; t {
		case NOP:
			console.StackPtr++
		case JMP:
			console.StackPtr += a
		case ACC:
			console.Accumulator += a
			console.StackPtr++
		}
	}
	fmt.Println(console.Accumulator)
}

func backtrack(cur int, N int, console *GameConsole, vis map[int]bool, res *int, found *bool, changed bool) {
	// fmt.Println(cur, console)
	if cur > N || vis[cur] || *found {
		return
	}
	if cur == N {
		*res = console.Accumulator
		*found = true
		// fmt.Println(*res)
		return
	}
	vis[console.StackPtr] = true
	switch t, a := ops[console.StackPtr].Type, ops[console.StackPtr].Arg; t {
	case NOP:
		console.StackPtr++
		backtrack(console.StackPtr, N, console, vis, res, found, changed)
		console.StackPtr--
		if changed {
			break
		}
		changed = true
		console.StackPtr += a
		backtrack(console.StackPtr, N, console, vis, res, found, changed)
		console.StackPtr -= a
		changed = false
	case JMP:
		console.StackPtr += a
		backtrack(console.StackPtr, N, console, vis, res, found, changed)
		console.StackPtr -= a
		if changed {
			break
		}
		changed = true
		console.StackPtr++
		backtrack(console.StackPtr, N, console, vis, res, found, changed)
		console.StackPtr--
		changed = false
	case ACC:
		console.Accumulator += a
		console.StackPtr++
		backtrack(console.StackPtr, N, console, vis, res, found, changed)
		console.Accumulator -= a
		console.StackPtr--
	}
	vis[console.StackPtr] = false
}

func Part2() {
	res := 0
	console := &GameConsole{}
	vis := map[int]bool{}
	found := false

	backtrack(0, len(ops), console, vis, &res, &found, false)
	if ops[0].Type == NOP || ops[0].Type == JMP {
		backtrack(0, len(ops), console, vis, &res, &found, true)
	}
	fmt.Println(res)
}
