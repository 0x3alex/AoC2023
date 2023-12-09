package main

import (
	"GoAoC/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solveDayNine2(i []int) int {
	var tmp []int
	for k := 0; k < len(i); k++ {
		if k+1 >= len(i) {
			continue
		}
		tmp = append(tmp, i[k+1]-i[k])
	}

	if utils.All(i, func(i int) bool { return i == 0 }) {
		return 0
	}
	if len(tmp) == 1 {
		return tmp[0]
	}
	return tmp[0] - solveDayNine2(tmp)
}

func solveDayNine1(i []int) int {
	var tmp []int
	for k := 0; k < len(i); k++ {
		if k+1 >= len(i) {
			continue
		}
		tmp = append(tmp, i[k+1]-i[k])
	}

	if utils.All(i, func(i int) bool { return i == 0 }) {
		return 0
	}
	if len(tmp) == 1 {
		return tmp[0]
	}
	return tmp[len(tmp)-1] + solveDayNine1(tmp)
}

func dayNine(file *os.File, p2 bool) {
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		var i []int
		raw := scanner.Text()
		sp := strings.Split(raw, " ")
		for _, v := range sp {
			if v != " " {
				a, _ := strconv.Atoi(v)
				i = append(i, a)
			}
		}
		if !p2 {
			n := solveDayNine1(i)
			sum += i[len(i)-1] + n
		} else {
			n := solveDayNine2(i)
			sum += i[0] - n
		}
	}
	println(sum)
}
