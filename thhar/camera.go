package renderer

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"gonum.org/v1/gonum/mat"
)

func Get_Wh_Ww(screen *ebiten.Image) (width, height int) {
	return screen.Bounds().Dx(), screen.Bounds().Dy()
}

var forward = mat.NewDense(1, 4, []float64{0, 0, 1, 1})
var up = mat.NewDense(1, 4, []float64{0, 1, 0, 1})
var right = mat.NewDense(1, 4, []float64{1, 0, 0, 1})
var H_fov float64 = math.Pi / 3
var V_fov float64 = H_fov * float64(900) / float64(1600)
var Near_plane, Far_plane = 0.1, 100
var camera_speed = 0.3
var camera_rotation_speed = 0.003

type Camera struct {
	Position            Point_3d
	Position_last_frame Point_3d
	Angle_pitch         float64
	Angle_yaw           float64
	Angle_roll          float64
	Forward             *mat.Dense
	up                  *mat.Dense
	right               *mat.Dense
}

func New_camera(position Point_3d) *Camera {

	return &Camera{
		Position:    position,
		Angle_pitch: 0,
		Angle_yaw:   0,
		Angle_roll:  0,
		Forward:     forward,
		up:          up,
		right:       right,
	}
}

func camera_yaw(c *Camera, angle float64) {
	c.Angle_yaw += angle
}

func camera_pitch(c *Camera, angle float64) {
	c.Angle_pitch += angle
}

func axii_identity(c *Camera) {
	c.Forward = mat.NewDense(1, 4, []float64{0, 0, 1, 1})
	c.up = mat.NewDense(1, 4, []float64{0, 1, 0, 1})
	c.right = mat.NewDense(1, 4, []float64{1, 0, 0, 1})
}

func camera_update_axii(c *Camera) {
	rotate := mat.NewDense(4, 4, nil)
	rotate.Mul(rotate_x(c.Angle_pitch), rotate_y(c.Angle_yaw))
	axii_identity(c)
	// print_matrix_mat(c.forward)

	c.Forward.Mul(c.Forward, rotate)
	c.right.Mul(c.right, rotate)
	c.up.Mul(c.up, rotate)

}

func camera_matrix(c *Camera) *mat.Dense {
	camera_update_axii(c)
	matrix := mat.NewDense(4, 4, nil)
	tr_matrix := translate_matrix(c)
	rot_matrix := rotate_matrix(c)
	matrix.Mul(tr_matrix, rot_matrix)
	return matrix
}

func translate_matrix(c *Camera) *mat.Dense {
	x, y, z := c.Position.X, c.Position.Y, c.Position.Z

	return mat.NewDense(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1,
	})
}

func rotate_matrix(c *Camera) *mat.Dense {
	rx, ry, rz := c.right.At(0, 0), c.right.At(0, 1), c.right.At(0, 2)
	fx, fy, fz := c.Forward.At(0, 0), c.Forward.At(0, 1), c.Forward.At(0, 2)
	ux, uy, uz := c.up.At(0, 0), c.up.At(0, 1), c.up.At(0, 2)
	return mat.NewDense(4, 4, []float64{
		rx, ux, fx, 0,
		ry, uy, fy, 0,
		rz, uz, fz, 0,
		0, 0, 0, 1,
	})
}

func HandleMovement(c *Camera) {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		scaled := mat.NewDense(1, 4, nil)
		scaled.Scale(camera_speed, c.right)
		c.Position.X += scaled.At(0, 0)
		c.Position.Y += scaled.At(0, 1)
		c.Position.Z += scaled.At(0, 2)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		scaled := mat.NewDense(1, 4, nil)
		scaled.Scale(camera_speed, c.right)
		c.Position.X -= scaled.At(0, 0)
		c.Position.Y -= scaled.At(0, 1)
		c.Position.Z -= scaled.At(0, 2)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		scaled := mat.NewDense(1, 4, nil)
		scaled.Scale(camera_speed, c.Forward)
		c.Position.X -= scaled.At(0, 0)
		c.Position.Y -= scaled.At(0, 1)
		c.Position.Z -= scaled.At(0, 2)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		scaled := mat.NewDense(1, 4, nil)
		scaled.Scale(camera_speed, c.Forward)
		c.Position.X += scaled.At(0, 0)
		c.Position.Y += scaled.At(0, 1)
		c.Position.Z += scaled.At(0, 2)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		scaled := mat.NewDense(1, 4, nil)
		scaled.Scale(camera_speed, c.up)
		c.Position.X -= scaled.At(0, 0)
		c.Position.Y -= scaled.At(0, 1)
		c.Position.Z -= scaled.At(0, 2)
	}

	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		scaled := mat.NewDense(1, 4, nil)
		scaled.Scale(camera_speed, c.up)
		c.Position.X += scaled.At(0, 0)
		c.Position.Y += scaled.At(0, 1)
		c.Position.Z += scaled.At(0, 2)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.Angle_yaw -= camera_rotation_speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.Angle_yaw += camera_rotation_speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		c.Angle_pitch -= camera_rotation_speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		c.Angle_pitch += camera_rotation_speed
	}

}
