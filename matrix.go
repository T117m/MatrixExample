package main

import "math"

type Matrix struct {
	Rows, Cols int
	Data       []float64
}

func NewMatrix(r, c int) *Matrix {
	return &Matrix{
		Rows: r,
		Cols: c,
		Data: make([]float64, r*c),
	}
}

func (m *Matrix) at(i, j int) float64 {
	return m.Data[i*m.Cols+j]
}

func (m *Matrix) isSquare() bool {
	return m.Rows == m.Cols
}

func Product(a, b Matrix) (*Matrix, bool) {
	if a.Rows != b.Cols {
		return &Matrix{}, false
	}

	var (
		m = NewMatrix(a.Rows, b.Cols)
		n = a.Rows
	)

	for i := range n {
		for j := range n {
			var mul float64
			for k := range n {
				mul += a.at(i, k) * b.at(j, k)
			}
			m.Data[i*m.Cols+j] = mul
		}
	}

	return m, true
}

func (m *Matrix) simpleDet() float64 {
	if m.isSquare() && m.Rows == 2 {
		return m.at(0, 0)*m.at(1, 1) - m.at(0, 1)*m.at(1, 0)
	}

	return 0
}

func (m *Matrix) minorMatrix(y, x int) *Matrix {
	var (
		res              = NewMatrix(m.Rows-1, m.Cols-1)
		vOffset, hOffset = 0, 0
	)

	for i := range m.Rows {
		if i == y {
			hOffset = 1
			continue
		}

		for j := range m.Cols {
			if j == x {
				vOffset = 1
				continue
			}

			res.Data[(i-hOffset)*res.Cols+j-vOffset] = m.at(i, j)
		}

		vOffset = 0
	}

	return res
}

//									   bool: true - row, false - column
func (m *Matrix) findMostZeros() (int, bool) {
	var (
		rZeros = make(map[int]int, m.Rows)
		cZeros = make(map[int]int, m.Cols)

		maxR, maxC = 0, 0
	)

	for i := range m.Rows {
		for j := range m.Cols {
			if m.at(i, j) == 0 {
				rZeros[i]++
				if rZeros[i] > rZeros[maxR] {
					maxR = i
				}

				cZeros[i]++
				if cZeros[i] > cZeros[maxC] {
					maxC = j
				}
			}
		}
	}

	if cZeros[maxC] > rZeros[maxR] {
		return maxC, false
	}

	return maxR, true
}

func (m *Matrix) laplasDet() float64 {
	if !m.isSquare() {
		return 0
	}

	if m.Rows == 2 {
		return m.simpleDet()
	}

	var (
		det          float64
		index, isRow = m.findMostZeros()
	)

	if isRow {
		for j := range m.Cols {
			det += m.at(index, j) * math.Pow((-1), float64(index+j)) * m.minorMatrix(index, j).laplasDet()
		}
	} else {
		for i := range m.Rows {
			det += m.at(i, index) * math.Pow((-1), float64(i+index)) * m.minorMatrix(i, index).laplasDet()
		}
	}

	return det
}

func (m *Matrix) adjMatrix() *Matrix {
	var (
		adj = NewMatrix(m.Rows, m.Cols)
	)

	for i := range adj.Rows {
		for j := range adj.Cols {
			adj.Data[i*adj.Cols+j] = math.Pow((-1), float64(i+j)) * m.minorMatrix(i, j).laplasDet()
		}
	}

	return adj
}

func (m *Matrix) divide(n float64) *Matrix {
	var (
		res = NewMatrix(m.Rows, m.Cols)
	)

	for i := range res.Rows {
		for j := range res.Cols {
			res.Data[i*res.Cols+j] = m.at(i, j) / n
		}
	}

	return res
}

func Inverse(a Matrix) Matrix {
	return *a.adjMatrix().divide(a.laplasDet())
}

func InverseProduct(a, b Matrix) (*Matrix, bool) {
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
