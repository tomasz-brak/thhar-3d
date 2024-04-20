package renderer

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

var p_right = math.Tan(H_fov / 2)
var left = -p_right
var top = math.Tan(V_fov / 2)
var bottom = -top

var m00 = 2 / (p_right - left)
var m11 = 2 / (top - bottom)
var m22 = (float64(Far_plane) + Near_plane) / (float64(Far_plane) - float64(Near_plane))
var m32 = -2 * float64(Far_plane) * float64(Near_plane) / (float64(Far_plane) - float64(Near_plane))

var Projection_matrix = mat.NewDense(4, 4, []float64{
	m00, 0, 0, 0,
	0, m11, 0, 0,
	0, 0, m22, 1,
	0, 0, m32, 0,
})

func To_screen_matrix(Window_width, Window_height int) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		float64(Window_width) / 2, 0, 0, 0,
		0, float64(-Window_height) / 2, 0, 0,
		0, 0, 1, 0,
		float64(Window_width) / 2, float64(Window_height) / 2, 0, 1,
	})
}

func Test() {
	fmt.Printf("top - bottom: %v, top is: %v, bottom is: %v\n", top-bottom, top, bottom)
	fmt.Printf("%v; %v; %v; %v", m00, m11, m22, m32)
}
