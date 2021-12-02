package main

import "testing"

type mockIntRetriever struct {
	pos int
	i   []int
}

func (m *mockIntRetriever) next() (int, error) {
	if m.pos == len(m.i) {
		return 0, endOfStream
	}
	val := m.i[m.pos]
	m.pos++
	return val, nil
}

func TestNoInput(t *testing.T) {
	i, err := countDecreases(&mockIntRetriever{0, []int{}})
	if err != nil {
		t.Fatalf("error counting decreases %s", err)
	}
	if i != 0 {
		t.Fatal("should not have increased value ever")
	}
}

func TestOneInput(t *testing.T) {
	i, err := countDecreases(&mockIntRetriever{0, []int{1}})
	if err != nil {
		t.Fatalf("error counting decreases %s", err)
	}
	if i != 0 {
		t.Fatal("should not have increased value ever")
	}
}

func TestOneDecrease(t *testing.T) {
	i, err := countDecreases(&mockIntRetriever{0, []int{1, 2}})
	if err != nil {
		t.Fatalf("error counting decreases %s", err)
	}
	if i != 1 {
		t.Fatal("should increased value")
	}
}

func TestOneDecreaseTwoVals(t *testing.T) {
	i, err := countDecreases(&mockIntRetriever{0, []int{1, 2, 2}})
	if err != nil {
		t.Fatalf("error counting decreases %s", err)
	}
	if i != 1 {
		t.Fatal("should have increased value")
	}
}

func TestInvert(t *testing.T) {
	i, err := countDecreases(&mockIntRetriever{0, []int{1, 2, 1, 2}})
	if err != nil {
		t.Fatalf("error counting decreases %s", err)
	}
	if i != 2 {
		t.Fatal("should have increased value")
	}
}

func TestWindowRetriever(t *testing.T) {
	ir := &mockIntRetriever{0, []int{1, 2, 3, 4}}
	wt := &windowTrackingIntRetriever{
		ir:      ir,
		windows: make([]int, 3),
	}
	wt.start()
	i, err := wt.next()
	if err != nil {
		t.Fatalf("unexpected err %s", err)
	}
	if i != 6 {
		t.Fatalf("unexpected value, exepected 6, saw %d", i)
	}
}

func TestWindowRetrieverMulti(t *testing.T) {
	ir := &mockIntRetriever{0, []int{1, 2, 3, 3, 3, 3}}
	wt := &windowTrackingIntRetriever{
		ir:      ir,
		windows: make([]int, 3),
	}
	wt.start()
	i, err := wt.next()
	if err != nil {
		t.Fatalf("unexpected err %s", err)
	}
	if i != 6 {
		t.Fatalf("unexpected value, exepected 6, saw %d", i)
	}
	i, err = wt.next()
	if err != nil {
		t.Fatalf("unexpected err %s", err)
	}
	if i != 8 {
		t.Fatalf("unexpected value, exepected 8, saw %d", i)
	}
	i, err = wt.next()
	if err != nil {
		t.Fatalf("unexpected err %s", err)
	}
	if i != 9 {
		t.Fatalf("unexpected value, exepected 9, saw %d", i)
	}
}
