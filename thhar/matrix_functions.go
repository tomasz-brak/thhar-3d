package renderer

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func translate(pos Point_3d) *mat.Dense {
	tx, ty, tz := pos.X, pos.Y, pos.Z
	return mat.NewDense(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		tx, ty, tz, 1,
	})
}

func rotate_x(a float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		1, 0, 0, 0,
		0, math.Cos(a), math.Sin(a), 0,
		0, -math.Sin(a), math.Cos(a), 0,
		0, 0, 0, 1,
	})
}

func rotate_y(a float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		math.Cos(a), 0, -math.Sin(a), 0,
		0, 1, 0, 0,
		math.Sin(a), 0, math.Cos(a), 0,
		0, 0, 0, 1,
	})
}

func rotate_z(a float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		math.Cos(a), math.Sin(a), 0, 0,
		-math.Sin(a), math.Cos(a), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

func scale(s float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		s, 0, 0, 0,
		0, s, 0, 0,
		0, 0, s, 0,
		0, 0, 0, 1,
	})
}
