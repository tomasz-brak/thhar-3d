package renderer

import (
	"image/color"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gonum.org/v1/gonum/mat"
)

var rotation = true

func cutof2(matrix *mat.Dense) {
	// var values = []float64{
	// 	matrix.At(0, 0), matrix.At(0, 1), matrix.At(0, 2), matrix.At(0, 3),
	// }

	// for idx, value := range values {
	// 	if value > 2 {
	// 		values[idx] = 0
	// 	}
	// 	if value < -2 {
	// 		values[idx] = 0
	// 	}
	// }

	// for i := 0; i < 4; i++ {
	// 	matrix.Set(0, i, values[i])
	// }
}

func drawLine(screen *ebiten.Image, x1, y1, x2, y2 float32, color color.Color) {
	vector.StrokeLine(screen, x1, y1, x2, y2, 2, color, true)
}

func scale_vertex_by_origin(v, origin *Point_3d) {
	v.X -= origin.X
	v.Y -= origin.Y
	v.Z -= origin.Z
}

func render_object(o *Object_3d, screen *ebiten.Image, cam *Camera, colored color.Color) {
	log.Printf("New Frame!")
	vertices := o.Vertices
	var matrixified []mat.Dense

	rotation := o.Rotation
	var to_cut_out []int

	for _, vertex := range vertices {
		// scale_vertex_by_origin(&vertex, &o.Origin)
		matrix := mat.NewDense(1, 4, []float64{vertex.X, vertex.Y, vertex.Z, 1})
		matrixified = append(matrixified, *matrix)
		// print_matrix_mat(matrix)
	}

	// Test()

	// for idx := range matrixified {
	// 	print_matrix_mat(&matrixified[idx])
	// }

	for idx := range matrixified {
		matrixified[idx].Mul(&matrixified[idx], rotate_x(rotation.X))
		matrixified[idx].Mul(&matrixified[idx], rotate_y(rotation.Y))
		matrixified[idx].Mul(&matrixified[idx], rotate_z(rotation.Z))
	}

	for idx := range matrixified {
		matrixified[idx] = *mat.NewDense(1, 4, []float64{matrixified[idx].At(0, 0) - o.Origin.X, matrixified[idx].At(0, 1) - o.Origin.Y, matrixified[idx].At(0, 2) - o.Origin.Z, matrixified[idx].At(0, 3)})
	}

	var zs []float64    // get the d coordinates
	var cordz []float64 // get the z coordinates

	for idx := range matrixified {
		matrixified[idx].Mul(&matrixified[idx], camera_matrix(cam))
		matrixified[idx].Mul(&matrixified[idx], Projection_matrix)
		// print_matrix_mat(&matrixified[idx])

		zs = append(zs, matrixified[idx].At(0, 3))
		cordz = append(cordz, matrixified[idx].At(0, 2))

	}

	for idx := range cordz {
		if cordz[idx] <= 0 {
			// TODO Remove the point compleatly!
			to_cut_out = append(to_cut_out, idx)
		}
		// log.Printf("Not Cutting: %v", matrixified[idx].At(0, 2))
	}

	for idx := range matrixified {
		matrixified[idx] = *mat.NewDense(1, 4, []float64{matrixified[idx].At(0, 0) / zs[idx], matrixified[idx].At(0, 1) / zs[idx], matrixified[idx].At(0, 2) / zs[idx], 1})

		cutof2(&matrixified[idx])

		matrixified[idx].Mul(&matrixified[idx], To_screen_matrix(screen.Bounds().Dx(), screen.Bounds().Dy()))

	}

	var cords_2d [][]float64

	for _, matrix := range matrixified {
		cords_2d = append(cords_2d, []float64{matrix.At(0, 0), matrix.At(0, 1)})
	}

	for idx, _ := range cords_2d {
		// log.Printf("Trying to draw circle! with cords: %v, %v", cord[0], cord[1])
		if slices.Contains(to_cut_out, idx) {
			// log.Printf("Skipping circle!")
			continue
		}
		// vector.DrawFilledCircle(screen, float32(cord[0]), float32(cord[1]), 5, color.White, true)
		// log.Printf("Drawing a point on 3d cords: %v; %v; %v [x: %v, y: %v] (translated: %v)", vertices[idx].X, vertices[idx].Y, vertices[idx].Z, cord[0], cord[1], matrixified[idx])
	}

	// draw lines with faces

	for _, face := range o.Faces {
		var xs []float64
		var ys []float64
		var is_displayed []bool
		for _, index := range face {
			xs, ys = append(xs, cords_2d[index][0]), append(ys, cords_2d[index][1])
			is_displayed = append(is_displayed, !slices.Contains(to_cut_out, index))
		}

		for idx := range xs {
			plus_one := idx + 1
			if plus_one >= len(xs) {
				plus_one = 0
			}
			if !is_displayed[idx] || !is_displayed[plus_one] {
				continue
			}
			drawLine(screen, float32(xs[idx]), float32(ys[idx]), float32(xs[plus_one]), float32(ys[plus_one]), colored)
		}

	}

}
