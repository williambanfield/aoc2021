package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	endOfStream = errors.New("End Of Stream")
)

type intRetriever interface {
	next() (int, error)
}

type scannerIntRetriever struct {
	*bufio.Scanner
}

func (s scannerIntRetriever) next() (int, error) {
	if !s.Scan() {
		return 0, endOfStream
	}
	return strconv.Atoi(s.Text())
}

type windowTrackingIntRetriever struct {
	offset  int
	windows []int
	ir      intRetriever
}

func (wt *windowTrackingIntRetriever) next() (int, error) {
	n, err := wt.ir.next()
	if err != nil {
		return 0, err
	}
	for i := range wt.windows {
		wt.windows[i] += n
	}
	v := wt.windows[wt.offset%len(wt.windows)]
	wt.windows[wt.offset%len(wt.windows)] = 0
	wt.offset++
	return v, nil
}
func (wt *windowTrackingIntRetriever) start() error {
	_, err := wt.next()
	if err != nil {
		return err
	}
	_, err = wt.next()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(fmt.Sprintf("error opening input file %s", err))
	}
	s := bufio.NewScanner(f)
	ir := scannerIntRetriever{s}
	wt := &windowTrackingIntRetriever{
		windows: make([]int, 3),
		ir:      ir,
	}
	wt.start()
	count, err := countDecreases(wt)
	if err != nil {
		panic(fmt.Sprintf("error opening input file %s", err))
	}
	fmt.Println(count)
}

func countDecreases(i intRetriever) (int, error) {
	prev := math.MaxInt
	var err error
	var next int
	var val int
	for next, err = i.next(); err == nil; next, err = i.next() {
		if next > prev {
			val++
		}
		prev = next
	}
	if !errors.Is(err, endOfStream) {
		return 0, err
	}
	return val, nil
}
