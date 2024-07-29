package main

import "testing"

func TestSum(t *testing.T) {
	rs := []struct {
		a, b, r int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 3, 5},
	}

	for _, r := range rs {
		want := r.r
		got := Sum(r.a, r.b)
		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	}
}

func TestSt_Sum(t *testing.T) {
	rs := []struct {
		a, b, r int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 3, 5},
	}
	st := St{}
	for _, r := range rs {
		want := r.r
		got := st.Sum(r.a, r.b)
		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	}
}
