package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	up int = iota
	down
	forward
)

func main() {
	in, err := parseInput("input")
	if err != nil {
		panic(err)
	}
	p1 := part1(in)
	fmt.Println(p1)
	p2 := part2(in)
	fmt.Println(p2)
}

func parseInput(fname string) ([][2]int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	vals := [][2]int{}
	for s.Scan() {
		strs := strings.Split(s.Text(), " ")
		if len(strs) != 2 {
			return nil, errors.New("malformed input")
		}
		val, err := strconv.Atoi(strs[1])
		if err != nil {
			return nil, errors.New("malformed input")
		}
		switch strs[0] {
		case "up":
			vals = append(vals, [2]int{val, up})
		case "down":
			vals = append(vals, [2]int{val, down})
		case "forward":
			vals = append(vals, [2]int{val, forward})
		}
	}
	return vals, nil
}

func part1(vals [][2]int) int {
	var vertical, horizontal int
	for i := range vals {
		switch vals[i][1] {
		case up:
			vertical -= vals[i][0]
		case down:
			vertical += vals[i][0]
		case forward:
			horizontal += vals[i][0]
		}
	}
	return horizontal * vertical
}

func part2(vals [][2]int) int {
	var depth, horizontal, aim int
	for i := range vals {
		switch vals[i][1] {
		case up:
			aim -= vals[i][0]
		case down:
			aim += vals[i][0]
		case forward:
			depth += vals[i][0] * aim
			horizontal += vals[i][0]
		}
	}
	return horizontal * depth
}
