package main

import (
	"GoAoC/utils"
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var values = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
	"0": 0,
}

type card struct {
	fields map[int]int
	intTxt []int

	filed2   map[int]int
	intText2 []int
	price    int
	rating   int
	text     string
	text2    string
}

func getRating(m map[int]int) int {
	var tmpRaitings []int
	for k, v := range m {
		println(k, ":", v)
		if v == 5 {
			return 7
		}
		if v == 4 {
			return 6
		}
		if v == 3 {
			tmpRaitings = append(tmpRaitings, 5)
		}
		if v == 2 {
			if len(tmpRaitings) == 1 && tmpRaitings[0] == 2 {
				return 3
			} else {
				tmpRaitings = append(tmpRaitings, 2)

			}
		}
	}
	for _, v := range tmpRaitings {
		println(v)
	}
	if utils.Any(tmpRaitings, func(i int) bool { return i == 5 }) && utils.Any(tmpRaitings, func(i int) bool { return i == 2 }) {
		return 5
	} else if len(tmpRaitings) == 1 && tmpRaitings[0] == 2 {
		return 2
	} else if utils.Any(tmpRaitings, func(i int) bool { return i == 5 }) && !utils.Any(tmpRaitings, func(i int) bool { return i == 2 }) {
		return 4
	} else {
		return 1
	}
}

func fixJokers(s string) string {
	if s == "JJJJJ" {
		return "AAAAA"
	}
	count := 0
	var ru rune
	for _, v := range s {
		if c := strings.Count(s, string(v)); c > count && v != 'J' {
			ru = v
			count = c
		}
	}
	return strings.ReplaceAll(s, "J", string(ru))
}

func buildCard(sp []string) card {
	hand := sp[0]
	rating, _ := strconv.Atoi(sp[1])
	r := make(map[int]int)
	var intTxt []int
	for _, p := range strings.Split(hand, "") {

		v, _ := values[p]
		if _, ok := r[v]; ok {
			r[v] += 1
		} else {
			r[v] = 1
		}
		intTxt = append(intTxt, v)
	}
	return card{
		text:   hand,
		fields: r,
		intTxt: intTxt,
		price:  rating,
		rating: getRating(r),
	}
}

func daySeven(file *os.File, p2 bool) {
	scanner := bufio.NewScanner(file)
	var cards []card
	for scanner.Scan() {
		raw := scanner.Text()
		sp := strings.Split(raw, " ")
		if p2 {
			//println(sp[0])
			//println(fixJokers(sp[0]))
			//sp[0] = fixJokers(sp[0])
			var tmp []string
			tmp = append(tmp, strings.ReplaceAll(sp[0], "J", "1"))
			tmp = append(tmp, sp[1])
			c := buildCard(tmp)
			println(sp[0])
			sp[0] = fixJokers(sp[0])
			println(sp[0])
			c1 := buildCard(sp)
			cr := card{
				fields:   c1.fields,
				intTxt:   c1.intTxt,
				filed2:   c.fields,
				intText2: c.intTxt,
				text:     c1.text,
				text2:    c.text,
				rating:   c1.rating,
				price:    c.price,
			}
			cards = append(cards, cr)
		} else {
			cards = append(cards, buildCard(sp))
		}

	}
	sort.SliceStable(cards, func(i, j int) bool {
		if cards[i].rating == cards[j].rating {
			if p2 {
				for k, v := range cards[i].intText2 {
					if v != cards[j].intText2[k] {
						return v < cards[j].intText2[k]
					}
				}
			} else {
				for k, v := range cards[i].intTxt {
					if v != cards[j].intTxt[k] {
						return v < cards[j].intTxt[k]
					}
				}
			}
		}
		return cards[i].rating < cards[j].rating
	})
	sum := 0
	for i, v := range cards {
		println(i, " Card ", v.text, "("+v.text2+")", " has rating ", v.rating, " and price ", v.price)
		sum += v.price * (i + 1)
	}
	println(sum)
}
