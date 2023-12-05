package main

import (
	"bufio"
	"math"
	"os"
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

func findTranslation(i int, m []map[string]int) int {
	if i < m[len(m)-1]["min"] || i > m[len(m)-1]["max"] {
		return i
	}
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
					seedsP2 = append(seedsP2, Tuple{s, e})
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
		m["sStart"] = sourceRangeStart
		m["sEnd"] = sourceRangeStart + rangeLength
		ma[currentMap] = append(ma[currentMap], m)
	}
	for k, v := range ma {
		var mMin, mMax int
		for i, m := range v {
			if i == 0 {
				mMin = m["sStart"]
				mMax = m["sEnd"]
				continue
			}
			mMin = int(math.Min(float64(m["sStart"]), float64(mMin)))
			mMax = int(math.Max(float64(m["sEnd"]), float64(mMax)))
		}
		ma[k] = append(ma[k], map[string]int{
			"min": mMin,
			"max": mMax,
		})
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

	}

	println("Lowest location is ", lowestLocation)
}
