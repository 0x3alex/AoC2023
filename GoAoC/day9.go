package main

import (
	"GoAoC/utils"
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

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
		sp := strings.Fields(raw)
		for _, v := range sp {
			a, _ := strconv.Atoi(v)
			i = append(i, a)
		}
		if !p2 {
			sum += i[len(i)-1] + solveDayNine1(i)
		} else {
			slices.Reverse(i)
			sum += i[len(i)-1] + solveDayNine1(i)
		}
	}
	println(sum)
}
