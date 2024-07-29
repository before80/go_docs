package main

func Sum(a, b int) int {
	return a + b
}

type St struct {
	a, b, r int
}

func (st *St) Sum(a, b int) int {
	st.r = a + b
	return st.r
}
