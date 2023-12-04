package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isAllSpace(s string) bool {
	for _, v := range s {
		if v != ' ' {
			return false
		}

	}
	return true
}

func extractNums(raw string) (nums []string, winning map[string]interface{}) {
	a := strings.TrimSpace(strings.Split(raw, ":")[1])
	b := strings.Split(a, "|")
	winningNum := strings.TrimSpace(b[0])
	winningNumsList := strings.Split(winningNum, " ")
	winningNumsMap := make(map[string]interface{})
	scratchedNums := strings.TrimSpace(b[1])
	scratchedNumsList := strings.Split(scratchedNums, " ")
	for _, v := range winningNumsList {
		if _, ok := winningNumsMap[v]; !ok && !isAllSpace(v) {
			v = strings.TrimSpace(v)
			winningNumsMap[v] = nil

		}
	}
	return scratchedNumsList, winningNumsMap
}

func dayFourP2Rec(buff []string, occs []int, idx int) []int {
start:
	//offset := 0
	currentCard := buff[idx]
	println(currentCard)
	split1 := strings.Split(currentCard, ":")
	println(strings.Split(split1[0], " ")[1])
	a := strings.Split(split1[0], " ")
	cardNum, _ := strconv.Atoi(strings.TrimSpace(a[len(a)-1]))
	println(cardNum)
	scratchedNumsList, winningNumsMap := extractNums(currentCard)
	matches := 0
	for _, v := range scratchedNumsList {
		if _, ok := winningNumsMap[v]; ok {
			matches++
		}
	}
	println("Card ", cardNum, " won ", matches, " cards")
	occsCardNum := cardNum - 1
	for m := 1; m <= matches; m++ {
		occs[occsCardNum+m] += occs[occsCardNum]
	}
	idx++
	if idx >= len(buff) {
		return occs
	}
	goto start
}

func dayFourP2(file *os.File) {
	var buff []string
	var occs []int
	s := bufio.NewScanner(file)
	for s.Scan() {
		raw := s.Text()
		buff = append(buff, raw)
		occs = append(occs, 1)
	}
	var sum int
	for _, v := range dayFourP2Rec(buff, occs, 0) {
		sum += v
	}
	println(sum)

}

func dayFour(file *os.File) {
	s := bufio.NewScanner(file)
	var totalSum int
	for s.Scan() {
		raw := s.Text()
		scratchedNumsList, winningNumsMap := extractNums(raw)
		var tempSum int
		for _, v := range scratchedNumsList {
			if _, ok := winningNumsMap[v]; ok {
				if tempSum == 0 {
					tempSum = 1
				} else {
					tempSum *= 2
				}
			}
		}
		totalSum += tempSum
	}
	println(totalSum)
}
