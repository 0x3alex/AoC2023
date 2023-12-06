package main

import (
	"GoAoC/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func daySixP2(file *os.File) {
	s := bufio.NewScanner(file)
	var d []int
	for s.Scan() {
		raw := s.Text()
		data, _ := strconv.Atoi(strings.ReplaceAll(raw[10:], " ", ""))
		d = append(d, data)

	}
	count := 0
	for i := 0; i <= d[0]; i++ {
		distance := i * (d[0] - i)
		if distance > d[1] {
			count++
		}
	}
	println(count)
}

func daySix(file *os.File) {
	s := bufio.NewScanner(file)
	offset := 0
	var d []int
	m := make(map[int]int)
	for s.Scan() {
		raw := s.Text()
		data := strings.Split(raw, " ")[1:]
		filtered := utils.Filter(data, utils.NotIsAllSpace)
		for _, v := range filtered {
			a, _ := strconv.Atoi(v)
			d = append(d, a)
		}
		offset = len(filtered)
	}
	for i := 0; i < (len(d) / 2); i++ {
		m[d[i]] = d[i+offset]
	}
	var variants []int
	for k, v := range m {
		count := 0
		for i := 0; i <= k; i++ {
			distance := i * (k - i)
			if distance > v {
				count++
			}
		}
		variants = append(variants, count)
	}
	product := 1
	for _, v := range variants {
		product *= v
	}
	println(product)

}
