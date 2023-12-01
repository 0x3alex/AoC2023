package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	numsMap = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func replaceToNum(s string) string {
	var acc, res string
	for _, v := range s {
		acc += string(v)
		for k, v1 := range numsMap {
			if strings.Contains(acc, k) {
				res += v1
				acc = string(v)
				break
			}
		}
		if unicode.IsDigit(v) {
			acc = ""
			res += string(v)
		}
	}
	return res
}

func dayOne(file *os.File) {
	s := bufio.NewScanner(file)
	sum := 0
	for s.Scan() {
		raw := s.Text()
		digits := replaceToNum(raw)
		if len(digits) > 2 {
			digits = digits[:1] + digits[len(digits)-1:]
		} else if len(digits) == 1 {
			digits += digits
		}
		i, err := strconv.Atoi(digits)
		if err != nil {
			panic(err.Error())
		}

		sum += i
	}
	println(sum)
}
