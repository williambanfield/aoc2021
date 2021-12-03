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

type directionInput struct {
	direction, val int
}

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

func parseInput(fname string) ([]directionInput, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	d := []directionInput{}
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
			d = append(d, directionInput{direction: up, val: val})
		case "down":
			d = append(d, directionInput{direction: down, val: val})
		case "forward":
			d = append(d, directionInput{direction: forward, val: val})
		}
	}
	return d, nil
}

func part1(in []directionInput) int {
	var vertical, horizontal int
	for _, d := range in {
		switch d.direction {
		case up:
			vertical -= d.val
		case down:
			vertical += d.val
		case forward:
			horizontal += d.val
		}
	}
	return horizontal * vertical
}

func part2(in []directionInput) int {
	var depth, horizontal, aim int
	for _, d := range in {
		switch d.direction {
		case up:
			aim -= d.val
		case down:
			aim += d.val
		case forward:
			depth += d.val * aim
			horizontal += d.val
		}
	}
	return horizontal * depth
}
