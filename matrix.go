package main

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

func Inverse(a Matrix) Matrix {
}

func InverseProduct(a, b Matrix) Matrix {
}

func InverseFuture(a Matrix) <-chan Matrix {
}
