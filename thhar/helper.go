package renderer

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func print_matrix_mat(m *mat.Dense) {
	x, y := m.Dims()
	fmt.Print("Matrix: ")
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%v, ", m.At(i, j))
		}
		fmt.Print("\n")
	}
}

func deg2rad(x float64) float64 {
	return (x * math.Pi / 180)
}
