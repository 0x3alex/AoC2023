package main

import "os"

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err.Error())
	}
	dayEight(file)
	file.Close()
}
