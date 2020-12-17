package D12

import (
	"os"
	"path"
	"bufio"
	"strconv"
	"fmt"
	"math"
)

const (
	EAST = iota
	NORTH
	WEST
	SOUTH
)

type Ship struct {
	X int64;
	Y int64;
	Facing int64;
}
type Instr struct {
	Op rune
	Arg int64
}


var instr []Instr

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D12/d12.txt"))
	scanner := bufio.NewScanner(file)
	ship := Ship{}
	for scanner.Scan() {
		val := scanner.Text()
		op := rune(val[0])
		_arg, _ := strconv.Atoi(val[1:])
		arg := int64(_arg)
		switch op {
		case 'N':
			ship.Y += arg
		case 'W':
			ship.X -= arg
		case 'S':
			ship.Y -= arg
		case 'E':
			ship.X += arg
		case 'F':
			switch ship.Facing {
			case EAST:
				ship.X += arg
			case WEST:
				ship.X -= arg
			case SOUTH:
				ship.Y -= arg
			case NORTH:
				ship.Y += arg
			}
		case 'L':
			ship.Facing = (ship.Facing + arg / 90) % 4
		case 'R':
			ship.Facing = (ship.Facing + (360-arg)/ 90) % 4
		}
		instr = append(instr, Instr{Op: op, Arg: arg})
	}
	if ship.X < 0 {
		ship.X = -ship.X
	}
	if ship.Y < 0 {
		ship.Y = -ship.Y
	}
	fmt.Println(ship.X + ship.Y)
}

func Part2() {
	wp := Ship{X: 10, Y: 1}
	ship := Ship{}
	for _, inst := range instr {
		op, arg := inst.Op, inst.Arg
		rad := math.Pi * (float64(arg) / 180)
		switch op {
		case 'N':
			wp.Y += arg
		case 'W':
			wp.X -= arg
		case 'S':
			wp.Y -= arg
		case 'E':
			wp.X += arg
		case 'F': // move the ship toward the waypoint
			ship.X = ship.X + wp.X * arg
			ship.Y = ship.Y + wp.Y * arg
		case 'L':
			leng := math.Sqrt(float64(wp.X * wp.X + wp.Y * wp.Y))
			wp.X = int64(float64(wp.X) * math.Cos(rad) - float64(wp.Y) * math.Sin(rad))
			wp.Y = int64(float64(wp.X) * math.Sin(rad) + float64(wp.Y) * math.Cos(rad))
			ratio := leng / math.Sqrt(float64(wp.X*wp.X + wp.Y *wp.Y))
			wp.X *= int64(ratio)
			wp.Y *= int64(ratio)
		case 'R':
			leng := math.Sqrt(float64(wp.X * wp.X + wp.Y * wp.Y))
			wp.X = int64(float64(wp.X) * math.Cos(-rad) - float64(wp.Y) * math.Sin(-rad))
			wp.Y = int64(float64(wp.X) * math.Sin(-rad) + float64(wp.Y) * math.Cos(-rad))
			ratio := leng / math.Sqrt(float64(wp.X*wp.X + wp.Y *wp.Y))
			wp.X *= int64(ratio)
			wp.Y *= int64(ratio)
		}
		fmt.Println(op, arg, ship)
		fmt.Println(wp)
		fmt.Println()
	}
	if ship.X < 0 {
		ship.X = -ship.X
	}
	if ship.Y < 0 {
		ship.Y = -ship.Y
	}
	// fmt.Println(ship.X, ship.Y)
	fmt.Println(ship.X + ship.Y)
}