package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tuple struct {
	a int
	b int
}

func isEmptyLine(s string) bool {
	for _, v := range s {
		if v != ' ' {
			return false
		}
	}
	return true
}

func findMin(m []map[string]int, ig []int) (int, int) {
	min := 0
	max := 0
outerLoop:
	for i, v := range m {
		for _, k := range ig {
			if v["sStart"] == k {
				continue outerLoop
			}
		}
		if i == 0 {
			min = v["sStart"]
		} else {
			min = int(math.Min(float64(min), float64(v["sStart"])))
		}
		max = v["sEnd"]
	}
	return min, max
}

func findMinSrcFromDest(i int, m []map[string]int) int {
	for _, a := range m {
		dStart := a["dStart"]
		dEnd := a["dEnd"]
		if dStart <= i && i <= dEnd {
			o := i - dStart
			return a["sStart"] + o
		}
	}
	return i
}

func findTranslation(i int, m []map[string]int) int {
	for _, v := range m {
		sStart := v["sStart"]
		sEnd := v["sEnd"]
		if sStart <= i && i <= sEnd {
			o := i - sStart
			return v["dStart"] + o
		}
	}
	return i
}

func dayFive(file *os.File, p2 bool) {
	s := bufio.NewScanner(file)
	var seeds []int
	var seedsP2 []Tuple
	var currentMap string
	ma := map[string][]map[string]int{
		"seed-to-soil":            {},
		"soil-to-fertilizer":      {},
		"fertilizer-to-water":     {},
		"water-to-light":          {},
		"light-to-temperature":    {},
		"temperature-to-humidity": {},
		"humidity-to-location":    {},
	}
outerloop:
	// build map
	for s.Scan() {
		raw := s.Text()
		if isEmptyLine(raw) {
			continue
		}
		if strings.Contains(raw, "seeds") {
			sp := strings.Split(raw, " ")[1:]
			if !p2 {
				for _, v := range sp {
					k, _ := strconv.Atoi(v)
					seeds = append(seeds, k)
				}
				continue
			} else {
				for i := 0; i < len(sp); i += 2 {
					s, _ := strconv.Atoi(sp[i])
					e, _ := strconv.Atoi(sp[i+1])
					seedsP2 = append(seedsP2, Tuple{s, s + e})
				}
				continue
			}

		}
		for k, _ := range ma {
			if strings.Contains(raw, k) {
				currentMap = k
				continue outerloop
			}
		}
		data := strings.Split(raw, " ")
		destinationRangeStart, _ := strconv.Atoi(data[0])
		sourceRangeStart, _ := strconv.Atoi(data[1])
		rangeLength, _ := strconv.Atoi(data[2])
		m := make(map[string]int)
		m["dStart"] = destinationRangeStart
		m["dEnd"] = destinationRangeStart + rangeLength
		m["sStart"] = sourceRangeStart
		m["sEnd"] = sourceRangeStart + rangeLength
		ma[currentMap] = append(ma[currentMap], m)
	}
	lowestLocation := 0
	if !p2 {
		for i, v := range seeds {
			a := findTranslation(v, ma["seed-to-soil"])
			b := findTranslation(a, ma["soil-to-fertilizer"])
			c := findTranslation(b, ma["fertilizer-to-water"])
			d := findTranslation(c, ma["water-to-light"])
			e := findTranslation(d, ma["light-to-temperature"])
			f := findTranslation(e, ma["temperature-to-humidity"])
			g := findTranslation(f, ma["humidity-to-location"])
			if i == 0 {
				lowestLocation = g
				continue
			}
			lowestLocation = int(math.Min(float64(lowestLocation), float64(g)))
		}
	} else {
		var res []int
		c := make(chan int)
		for _, v := range seedsP2 {
			a := v.a
			b := v.b
			go func(s, e int, k chan int) {
				lowest := 0
				first := true
				for i := s; i <= e; i++ {
					a := findTranslation(i, ma["seed-to-soil"])
					b := findTranslation(a, ma["soil-to-fertilizer"])
					c := findTranslation(b, ma["fertilizer-to-water"])
					d := findTranslation(c, ma["water-to-light"])
					e := findTranslation(d, ma["light-to-temperature"])
					f := findTranslation(e, ma["temperature-to-humidity"])
					g := findTranslation(f, ma["humidity-to-location"])
					if first {
						lowest = g
						first = false
						continue
					}
					lowest = int(math.Min(float64(lowest), float64(g)))

				}
				res = append(res, lowest)
				k <- 1
			}(a, b, c)
		}
		sum := 0
		for {
			_ = <-c
			sum++
			println("routine is done, ", sum, "/", len(seedsP2))
			if sum == len(seedsP2) {
				break
			}
		}
		sort.Ints(res)
		lowestLocation = res[0]
	}

	println("Lowest location is ", lowestLocation)
}
