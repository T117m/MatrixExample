package main

import "math"

type Matrix [][]float64

func NewMatrix(n int) Matrix {
	a := make(Matrix, n)

	for i := range a {
		a[i] = make([]float64, n)
	}

	return a
}

func compareSize(a, b Matrix) bool {
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
	if !compareSize(a, b) {
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

func simpleDet(a Matrix) float64 {
	if len(a) == 1 {
		return a[0][0]
	}

	return a[0][0]*a[1][1] - a[0][1]*a[1][0]
}

func minor(a Matrix, x, y int) float64 {
	n := len(a)

	if n == 1 {
		return simpleDet(a)
	}

	m := NewMatrix(n)

	copy(m, a[:x])
	if x != n {
		m = append(m, a[x+1:]...)
	}

	for i := range n - 1 {
		tmp := m[i]
		m[i] = m[i][:y]
		if y != n {
			m[i] = append(m[i], tmp[y+1:]...)
		}
	}

	return determinant(m)
}

func cofactor(a Matrix, x, y int) float64 {
	return math.Pow(float64(-1), float64(x+y)) * minor(a, x, y)
}

func determinant(a Matrix) float64 {
	n := len(a)

	if n < 3 {
		return simpleDet(a)
	}

	var det float64

	for j := range n {
		det += cofactor(a, 0, j) * a[0][j] 
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

	for i := range n {
		for j := range n {
			b[i][j] = cofactor(a, i, j)
		}
	}

	return transpose(b)
}

func div(a Matrix, x float64) Matrix {
	var (
		n = len(a)
		b = NewMatrix(n)
	)

	for i := range n {
		for j := range n {
			b[i][j] = a[i][j] / x
		}
	}

	return b
}

func Inverse(a Matrix) Matrix {
	return div(adj(a), determinant(a))
}

func InverseProduct(a, b Matrix) Matrix {
	var (
		aInvF = InverseFuture(a)
		bInvF = InverseFuture(b)
		aInv  = <-aInvF
		bInv  = <-bInvF
	)

	return Product(aInv, bInv)
}

func InverseFuture(a Matrix) <-chan Matrix {
	future := make(chan Matrix)
	go func() { future <- Inverse(a) }()
	return future
}
