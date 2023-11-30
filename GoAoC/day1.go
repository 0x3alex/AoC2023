package main

import (
	"bufio"
	"os"
)

func dayOne(file *os.File) {
	s := bufio.NewScanner(file)
	for s.Scan() {
		println(s.Text())
	}
}
