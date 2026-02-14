package main

import "math"

type Matrix [][]int

func NewMatrix(n int) Matrix {
	a := make(Matrix, n)

	for i := range a {
		a[i] = make([]int, n)
	}

	return a
}

func CompareSize(a, b Matrix) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
	}

	return true
}

func Product(a, b Matrix) Matrix {
	if !CompareSize(a, b) {
		return NewMatrix(0)
	}

	c := NewMatrix(len(a))

	for i := range a {
		for j := range b {
			for k := range c {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}

func simpleDet(a Matrix) int {
	if len(a) == 1 {
		return a[0][0]
	}

	return a[0][0]*a[1][1] - a[0][1]*a[1][0]
}

func minor(a Matrix, x, y int) int {
	n := len(a)

	if n < 3 {
		return simpleDet(a)
	}

	m := a[:x]
	m = append(m, a[x+1:]...)

	for i := range n {
		tmp := m[i]
		m[i] = m[i][:y]
		m[i] = append(m[i], tmp[x+1:]...)
	}

	return determinant(m)
}

func determinant(a Matrix) int {
	n := len(a)

	if n < 3 {
		return simpleDet(a)
	}

	var det int

	for j := range n {
		m := a[1:]

		switch j {
		case 0:
			for i := range n - 1 {
				m[i] = m[i][1:]
			}
		case n - 1:
			for i := range n - 1 {
				m[i] = m[i][:len(m)]
			}
		default:
			for i := range n - 1 {
				tmp := m[i]
				m[i] = m[i][:j]
				m[i] = append(m[i], tmp[j+1:]...)
			}
		}

		det += int(math.Pow(float64(-1), float64(1+j))) * a[0][j] * minor(m, 0, j)
	}

	return det
}

func transpose(a Matrix) Matrix {
	var (
		n = len(a)
		t = NewMatrix(n)
	)

	for i := range n {
		for j := range n {
			t[i][j] = a[j][i]
		}
	}

	return t
}

func adj(a Matrix) Matrix {
	var (
		n = len(a)
		b = NewMatrix(n)
	)

	return b
}

func Inverse(a Matrix) Matrix {
}

func InverseProduct(a, b Matrix) Matrix {
}

func InverseFuture(a Matrix) <-chan Matrix {
}
