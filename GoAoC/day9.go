package main

import (
	"GoAoC/utils"
	"bufio"
	"fmt"
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
		n := i[k+1] - i[k]
		tmp = append(tmp, n)
	}

	if utils.All(i, func(i int) bool { return i == 0 }) {
		return 0
	}
	if len(tmp) == 1 {
		return tmp[0]
	}
	n := solveDayNine2(tmp)
	return tmp[0] - n
}

func solveDayNine1(i []int) int {
	var tmp []int
	for k := 0; k < len(i); k++ {
		if k+1 >= len(i) {
			continue
		}
		n := i[k+1] - i[k]
		tmp = append(tmp, n)
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
			if v == " " {
				continue
			}
			a, _ := strconv.Atoi(v)
			i = append(i, a)
		}
		for _, v := range i {
			fmt.Printf("%d ", v)
		}
		println()
		if !p2 {
			n := solveDayNine1(i)
			println("adding ", n)
			sum += i[len(i)-1] + n
		} else {
			n := solveDayNine2(i)
			println("adding ", n)
			sum += i[0] - n
		}

		//2107948156
		//2098490942
	}
	println(sum)
}
