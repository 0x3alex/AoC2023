package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var field [][]rune

func printField() {
	for _, v := range field {
		for _, k := range v {
			fmt.Printf("%c", k)
		}
		println()
	}
}
func takeNum(col, row int) int {
	tC := col
	var r string
	for tC >= 0 && unicode.IsDigit(field[row][tC]) {
		tC--
	}
	if tC != col {
		tC++
	}
	for tC < len(field[col]) && unicode.IsDigit(field[row][tC]) {
		r += fmt.Sprintf("%c", field[row][tC])
		tC++
	}
	a, _ := strconv.Atoi(r)
	return a

}
func checkForNums(row, col int, p2 bool) int {
	if row < 0 || col < 0 || col >= len(field) || row >= len(field[0]) {
		return 0
	}
	if p2 && field[row][col] != '*' {
		return 0
	}
	results := make(map[int]string)
	oR := row - 1
	oC := col - 1
	for i := 0; i < 3; i++ {
		toR := oR + i
		if toR < 0 || toR >= len(field[0]) {
			continue
		}
		for j := 0; j < 3; j++ {
			toC := oC + j
			if toC < 0 || toC >= len(field) {
				continue
			}
			if toC == col && toR == row {
				continue
			}
			if unicode.IsDigit(field[toR][toC]) {
				println("found number adjecent to symbol")
				n := takeNum(toC, toR)
				if _, ok := results[n]; !ok {
					results[n] = ""
					println(n)
				}
			}
		}
	}
	if len(results) != 2 {
		return 0
	}
	s := 1
	for k, _ := range results {
		s *= k
	}
	return s

}

func goTroughField() (sum int) {
	for idx, i := range field {
		for jdx, j := range i {
			if j != '.' && !unicode.IsDigit(j) {
				println("found symbol")
				fmt.Printf("%c\n", j)
				//sum += checkForNums(idx, jdx, false) // p1
				sum += checkForNums(idx, jdx, true)
			}
		}
	}
	return sum
}

func dayThree(file *os.File) {
	s := bufio.NewScanner(file)
	for s.Scan() {
		var tmp []rune
		t := strings.Split(s.Text(), "")
		for _, v := range t {
			var tmp1 []rune
			for _, k := range v {
				tmp1 = append(tmp1, k)
			}
			//println(tmp1[0])
			tmp = append(tmp, tmp1[0])
		}
		field = append(field, tmp)
	}
	println(goTroughField())
}
