package main

type Matrix [][]int

func NewMatrix(n int) Matrix {
	a := make(Matrix, n)

	for i := range a {
		a[i] = make([]int, n)
	}

	return a
}

func Product(a, b Matrix) Matrix {
}

func Inverse(a Matrix) Matrix {
}

func InverseProduct(a, b Matrix) Matrix {
}

func InverseFuture(a Matrix) <-chan Matrix {
}
