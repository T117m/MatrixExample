package main

import (
	"testing"
)

func BenchmarkInverseProduct(b *testing.B) {
	var (
		m = NewMatrix(5, 5)
		n = NewMatrix(5, 5)
	)

	m.Fill(
		[]float64{0, 1, 2, 3, 4},
		[]float64{1, 2, 3, 4, 5},
		[]float64{2, 3, 4, 5, 6},
		[]float64{3, 4, 5, 6, 7},
		[]float64{4, 5, 6, 7, 8},
	)
	n.Fill(
		[]float64{0, 3, 4, 5, 6},
		[]float64{4, 5, 6, 7, 8},
		[]float64{9, 0, 2, 4, 5},
		[]float64{0, 9, 0, 2, 4},
		[]float64{3, 4, 5, 6, 7},
	)

	for b.Loop() {
		InverseProduct(*m, *n)
	}
}

func BenchmarkInverseProductConsequent(b *testing.B) {
	var (
		m = NewMatrix(5, 5)
		n = NewMatrix(5, 5)
	)

	m.Fill(
		[]float64{0, 1, 2, 3, 4},
		[]float64{1, 2, 3, 4, 5},
		[]float64{2, 3, 4, 5, 6},
		[]float64{3, 4, 5, 6, 7},
		[]float64{4, 5, 6, 7, 8},
	)
	n.Fill(
		[]float64{0, 3, 4, 5, 6},
		[]float64{4, 5, 6, 7, 8},
		[]float64{9, 0, 2, 4, 5},
		[]float64{0, 9, 0, 2, 4},
		[]float64{3, 4, 5, 6, 7},
	)

	for b.Loop() {
		InverseProductConsequent(*m, *n)
	}
}

func InverseProductConsequent(a, b Matrix) (Matrix, bool) {
	return Product(Inverse(a), Inverse(b))
}
