package D4

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Iyr int;
	Cid int;
	Pid string;
	Eyr int;
	Hcl string;
	Ecl string;
	Byr int;
	Hgt string;
}

var passports []Passport

func Part1() {
	cwd, _ := os.Getwd()
	file, _ := os.Open(path.Join(cwd, "D4/d4.txt"))
	scanner := bufio.NewScanner(file)
	pass := Passport{Cid: -1}
	s := map[string]bool{}
	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(s) == 8 || (len(s) == 7 && pass.Cid == -1) {
				cnt++
			}
			passports = append(passports, pass)
			pass = Passport{Cid: -1}
			s = map[string]bool{}
		} else {
			for _, pair := range strings.Split(line, " ") {
				keyval := strings.Split(pair, ":")
				key, val := keyval[0], keyval[1]
				switch key {
					case "iyr":
						ival, _ := strconv.Atoi(val)
						s["iyr"] = true
						pass.Iyr = ival
					case "cid":
						ival, _ := strconv.Atoi(val)
						s["cid"] = true
						pass.Cid = ival
					case "pid":
						s["pid"] = true
						pass.Pid = val
					case "eyr":
						ival, _ := strconv.Atoi(val)
						s["eyr"] = true
						pass.Eyr = ival
					case "hcl":
						s["hcl"] = true
						pass.Hcl = val
					case "ecl":
						s["ecl"] = true
						pass.Ecl = val
					case "byr":
						s["byr"] = true
						ival, _ := strconv.Atoi(val)
						pass.Byr = ival
					case "hgt":
						s["hgt"] = true
						pass.Hgt = val
					default:
						// do nothing
				}
			}
		}
	}
	fmt.Println(cnt)
}

/*
    byr (Birth Year) - four digits; at least 1920 and at most 2002.
    iyr (Issue Year) - four digits; at least 2010 and at most 2020.
    eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
    hgt (Height) - a number followed by either cm or in:
        If cm, the number must be at least 150 and at most 193.
        If in, the number must be at least 59 and at most 76.
    hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
    ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    pid (Passport ID) - a nine-digit number, including leading zeroes.
    cid (Country ID) - ignored, missing or not.
**/

func Part2() {
	cnt := 0
	hclReg := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eclMap := map[string]bool {
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	pidReg := regexp.MustCompile(`^[0-9]{9}$`)
	for i := 0; i < len(passports); i++ {
		p := passports[i]
		fmt.Println(p)
		if p.Byr < 1920 || p.Byr > 2002 {
			fmt.Println("byr: ", p.Byr)
			continue
		}
		if p.Iyr < 2010 || p.Iyr > 2020 {
			fmt.Println("Iyr: ", p.Iyr)
			continue
		}
		if p.Eyr < 2020 || p.Eyr > 2030 {
			fmt.Println("Eyr: ", p.Eyr)
			continue
		}

		heiOK := false
		if !strings.ContainsAny(p.Hgt, "cm") {
			if !strings.ContainsAny(p.Hgt, "in") {
				fmt.Println("Hgt1: ", p.Hgt)
				continue
			}
		} else {
			h, _ := strconv.Atoi(strings.TrimSuffix(p.Hgt, "cm"))
			if h < 150 || h > 193 {
				fmt.Println("Hgt2: ", h)
				continue
			}
			heiOK = true
		}
		if !heiOK {
			if !strings.ContainsAny(p.Hgt, "in") {
				fmt.Println("Hgt3: ", p.Hgt)
				continue
			} else {
				h, _ := strconv.Atoi(strings.TrimSuffix(p.Hgt, "in"))
				if h < 59 || h > 76 {
					fmt.Println("Hgt4: ", h)
					continue
				}
			}
		}

		if !hclReg.MatchString(p.Hcl) {
			fmt.Println("Hcl: ", p.Hcl)
			continue
		}
		 
		if _, ok := eclMap[p.Ecl]; !ok {
			fmt.Println("Ecl: ", p.Ecl)
			continue
		}
		if !pidReg.MatchString(string(p.Pid)) {
			fmt.Println("Pid: ", p.Pid)
			continue
		}
		cnt++
	}
	fmt.Println(cnt)
}