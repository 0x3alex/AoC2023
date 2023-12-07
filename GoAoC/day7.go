package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	price   int
	m       map[rune]int
	text    string
	rating  int
	oldText string
}

func fixJoker(s string) (string, rune) {
	count := 0
	var maxV rune
	for _, v := range s {
		if c := strings.Count(s, string(v)); c > count && v != 'J' {
			count = c
			maxV = v
		}
	}
	return strings.ReplaceAll(s, "J", string(maxV)), maxV
}

func MapOut(s string) (map[rune]int, string, string) {
	result := make(map[rune]int)
	k := s
	if s == "JJJJJ" {
		return map[rune]int{
			'A': 5,
		}, s, "AAAAA"
	}
	hasJoker := false
	for _, v := range s {
		if v == 'J' {
			hasJoker = true
		}
		if _, ok := result[v]; !ok {
			result[v] = 1
		} else {
			result[v] += 1
		}
	}

	var r rune
	if hasJoker {
		k, r = fixJoker(s)
		if k != s {
			result[r] += result['J']
			delete(result, 'J')
		}

	}
	return result, s, k
}

func determineHandValue(c *card) {
	twoPair := false
	for _, v := range c.m {
		if v == 5 { //five of a kind
			c.rating = 7
		}
		if v == 4 { //four of a kind
			c.rating = 6
		}
		if v == 3 { //three of a kind
			c.rating = 4
		}
		if v == 2 {
			if twoPair {
				c.rating = 3
			} else {
				twoPair = true
			}
		}
	}
	if twoPair && c.rating == 0 { //one pair
		c.rating = 2
	}
	if c.rating == 5 && twoPair {
		c.rating = 5
	}
	if !twoPair && c.rating == 0 {
		c.rating = 1
	}
}
func daySeven(file *os.File) {
	scan := bufio.NewScanner(file)
	var cards []card
	for scan.Scan() {
		raw := scan.Text()
		sp := strings.Split(raw, " ")
		hand := sp[0]
		price := sp[1]
		a, _ := strconv.Atoi(price)
		b, c, d := MapOut(hand)
		cards = append(cards, card{
			price:   a,
			m:       b,
			text:    d,
			oldText: c,
		})
	}
	for i, _ := range cards {
		determineHandValue(&cards[i])
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].rating == cards[j].rating {
			for l := 0; l < len(cards[i].text); l++ {
				if cards[i].text[l] != cards[j].text[l] {
					return cards[i].text[l] > cards[j].text[l]
				}
			}
			if cards[i].text == cards[j].text {
				if strings.Contains(cards[i].oldText, "J") {
					return false
				}
				return true
			}
		}
		return cards[i].rating < cards[j].rating
	})
	sum := 0
	for i, v := range cards {
		println("card ", v.text, " with price ", v.price, " has rating ", i+1)
		sum += v.price * (i + 1)
	}
	println(sum)
}
