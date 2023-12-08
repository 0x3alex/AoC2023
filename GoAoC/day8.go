package main

import (
	"GoAoC/utils"
	"bufio"
	"os"
	"strings"
)

type rl struct {
	name    string
	left    string
	right   string
	visited int
}

func gdc(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gdc(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func solve2(pos []rl, path string) int {
	m := make(map[string]*rl)
	var currentPoses []*rl
	for i, v := range pos {
		if strings.HasSuffix(v.name, "A") {
			currentPoses = append(currentPoses, &pos[i])
		}
		m[v.name] = &pos[i]
	}
	c := make(chan int)
	count := 0
	for _, v := range currentPoses {
		go func(name string) {
			c <- solve1(pos, path, name, false)
		}(v.name)
	}
	var res []int
	for count != len(currentPoses) {
		res = append(res, <-c)
		count++
		println(count, "/", len(currentPoses))
	}
	close(c)
	//,
	return lcm(res[0], res[1], res[2], res[3], res[4], res[5])
}

func solve1(pos []rl, path string, starting string, allZ bool) int {
	m := make(map[string]*rl)
	for i, v := range pos {
		m[v.name] = &pos[i]
	}
	count := 0
	offset := 0
	current, ok := m[starting]
	if !ok {
		println(starting)
		panic("not possible")
	}
	for (allZ && current.name != "ZZZ") || (!allZ && !strings.HasSuffix(current.name, "Z")) {
		op := path[offset]
		if rune(op) == 'R' {
			current = m[current.right]
		} else {
			current = m[current.left]
		}
		count++
		offset++
		if offset >= len(path) {
			offset = 0
		}
	}
	return count

}

func dayEight(file *os.File) {
	scanner := bufio.NewScanner(file)
	var pos []rl
	i := 0
	path := ""
	for scanner.Scan() {
		raw := scanner.Text()
		if i == 0 {
			path = raw
			i++
			continue
		}
		if i == 1 {
			i++
			continue
		}
		sp := strings.Split(raw, "=")
		loc := strings.TrimSpace(sp[0])
		k := strings.ReplaceAll(sp[1], "(", "")
		k = strings.ReplaceAll(k, ")", "")
		ksp := utils.Map(strings.Split(k, ","), strings.TrimSpace)
		pos = append(pos, rl{
			name:  loc,
			left:  ksp[0],
			right: ksp[1],
		})

	}
	/*for _, v := range pos {
		println(v.name, " has left ", v.left, " and right ", v.right)
	}*/
	println("Solve1 ", solve1(pos, path, "AAA", true))
	println("Solve2 ", solve2(pos, path))

}
