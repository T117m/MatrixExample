package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestProduct(t *testing.T) {
	var (
		a = NewMatrix(4, 4)
		b = NewMatrix(4, 4)
	)

	a.Fill(
		[]float64{2, 0, 1, 3},
		[]float64{4, 0, 0, 2},
		[]float64{0, 1, 2, 1},
		[]float64{5, 0, 3, 0},
	)

	b.Fill(
		[]float64{0, 0, 0, 0},
		[]float64{0, 0, 0, 0},
		[]float64{0, 0, 0, 0},
		[]float64{0, 0, 0, 0},
	)

	c, ok := Product(*a, *b)
	if !ok {
		t.Error("Multiplication failed\n", c)
	}

	t.Log("c:\n", c)

	assertEqual(t, c, b)
}

func TestInverse(t *testing.T) {
	var (
		a    = NewMatrix(3, 3)
		want = NewMatrix(3, 3)
	)

	a.Fill(
		[]float64{1, 5, 6},
		[]float64{3, 0, 2},
		[]float64{1, 1, 4},
	)
	want.Fill(
		[]float64{1, 0, 0},
		[]float64{0, 1, 0},
		[]float64{0, 0, 1},
	)

	aInv := Inverse(*a)

	t.Log("aInv:\n", &aInv)

	c, ok := Product(*a, aInv)
	if !ok {
		t.Error("Multiplication failed\n", c)
	}

	t.Log("c:\n", c)

	assertEqual(t, want, c)
}

func (m *Matrix) String() string {
	var s strings.Builder
	for i := range m.Rows {
		for j := range m.Cols {
			s.WriteString(fmt.Sprintf("%.2f ", m.at(i, j)))
		}
		s.WriteString("\n")
	}
	return s.String()
}

func assertEqual(t *testing.T, want, got *Matrix) {
	t.Helper()

	if want.Rows != got.Rows || want.Cols != got.Cols || len(want.Data) != len(got.Data) {
		t.Errorf("wanted:\n%s\ngot:\n%s", want, got)
	}

	for i := range want.Data {
		if math.Ceil(want.Data[i]*100) != math.Ceil(got.Data[i]*100) {
			t.Errorf("wanted:\n%s\ngot:\n%s", want, got)
		}
	}
}
