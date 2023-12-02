package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func dayTwoP2(file *os.File) {
	s := bufio.NewScanner(file)
	var sum int
	for s.Scan() {
		raw := s.Text()
		split1 := strings.Split(raw, ":")
		data := strings.TrimSpace(split1[1])
		eachDraw := strings.Split(data, ";")
		var sBlue, sRed, sGreen int
		for _, v := range eachDraw {
			v = strings.TrimSpace(v)
			eachQube := strings.Split(v, ",")
			for _, k := range eachQube {
				k = strings.TrimSpace(k)
				args := strings.Split(k, " ")
				num, _ := strconv.Atoi(args[0])
				switch args[1] {
				case "blue":
					if num > sBlue {
						sBlue = num
					}
					break
				case "red":
					if num > sRed {
						sRed = num
					}
					break
				case "green":
					if num > sGreen {
						sGreen = num
					}
					break
				}
			}

		}
		sum += sGreen * sRed * sBlue

	}
	println(sum)
}

func dayTwo(file *os.File) {
	s := bufio.NewScanner(file)
	var sum, blue, red, green int
	for s.Scan() {
		raw := s.Text()
		split1 := strings.Split(raw, ":")
		gameSection := split1[0]
		gameId := strings.Split(gameSection, " ")[1]
		data := strings.TrimSpace(split1[1])
		eachDraw := strings.Split(data, ";")
		valid := true
		for _, v := range eachDraw {
			red = 0
			green = 0
			blue = 0
			v = strings.TrimSpace(v)
			eachQube := strings.Split(v, ",")
			for _, k := range eachQube {
				k = strings.TrimSpace(k)
				args := strings.Split(k, " ")
				num, _ := strconv.Atoi(args[0])
				switch args[1] {
				case "blue":
					blue += num
					break
				case "red":
					red += num
					break
				case "green":
					green += num
					break
				}
			}
			if red > 12 || green > 13 || blue > 14 {
				valid = false
				break
			}

		}

		if valid {
			a, _ := strconv.Atoi(gameId)
			sum += a
		}

	}
	println(sum)
}
