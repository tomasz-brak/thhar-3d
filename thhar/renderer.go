package renderer

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point_3d struct {
	X, Y, Z float64
}

type Point_2d struct {
	X, Y float64
}

type Object_3d struct {
	Origin   Point_3d
	Vertices []Point_3d
	Faces    [][]int
	Rotation Point_3d
}

func screenProjection(screen *ebiten.Image, object Object_3d) {

}

var Window_height = 0
var Window_width = 0

func Set_window_size(screen *ebiten.Image) {
	Window_height = screen.Bounds().Dy()
	Window_width = screen.Bounds().Dx()
}

func Render(screen *ebiten.Image, objects []Object_3d, camera *Camera, colored color.Color) {
	Window_height = screen.Bounds().Dy()
	Window_width = screen.Bounds().Dx()

	for _, obj := range objects {
		render_object(&obj, screen, camera, colored)

	}

}

func GetObjectFromFile(filename string, scale float64) Object_3d {
	file, err := os.Open(filename)
	var obj Object_3d
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.Printf("GEtting an object")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "v ") {
			parts := strings.Split(line, " ")

			var vertex Point_3d
			vertex.X, _ = strconv.ParseFloat(parts[1], 64)
			vertex.Y, _ = strconv.ParseFloat(parts[2], 64)
			vertex.Z, _ = strconv.ParseFloat(parts[3], 64)

			vertex.X *= scale
			vertex.Y *= scale
			vertex.Z *= scale

			obj.Vertices = append(obj.Vertices, vertex)

		}
		if strings.HasPrefix(line, "f") {

			// log.Printf("line: %v", line)
			faces_ := strings.Split(line, " ")[1:]
			var single_face []int
			for _, face_ := range faces_ {
				number_str := strings.Split(face_, "/")[0]
				number, err := strconv.Atoi(number_str)
				if err != nil {
					panic(err)
				}
				single_face = append(single_face, number-1)
			}
			obj.Faces = append(obj.Faces, single_face)
		}
	}
	return obj
}

// angle - rad
func RotatePoint(pointToRotate Point_2d, midPoint Point_2d, angle float64) Point_2d {
	// Get our current angle as a cosine and sine.
	var out Point_2d

	angle = angle - deg2rad(0)

	out.X = math.Cos(angle)*(pointToRotate.X-midPoint.X) - math.Sin(angle)*(pointToRotate.Y-midPoint.Y) + pointToRotate.X
	out.Y = math.Sin(angle)*(pointToRotate.X-midPoint.X) + math.Cos(angle)*(pointToRotate.Y-midPoint.Y) + pointToRotate.Y

	return out
}

var x = 0.0

func is_out_of_circle(p Point_2d, r float64) bool {
	return math.Pow(p.X, 2)+math.Pow(p.Y, 2) > math.Pow(r, 2)
}

func radar(s *ebiten.Image, objects []Object_3d, camera *Camera) {
	Window_height = s.Bounds().Dy()
	Window_width = s.Bounds().Dx()

	pos_x := (Window_width / 2) + 50
	const pos_y = 100
	r := float64(80)

	color_red := color.RGBA{255, 0, 0, 255}

	// radar shape
	vector.StrokeCircle(s, float32(pos_x), float32(pos_y), float32(r), 1, color_red, true)

	// fov indicator
	var x1, y1 = math.Floor(r * math.Sin(deg2rad(330))), math.Floor(r * math.Cos(deg2rad(330)))
	var x2, y2 = math.Floor(r * math.Sin(deg2rad(30))), math.Floor(r * math.Cos(deg2rad(30)))

	vector.StrokeLine(s, float32(pos_x), float32(pos_y), float32(pos_x)-float32(x1), float32(pos_y)-float32(y1), 1, color_red, true)
	vector.StrokeLine(s, float32(pos_x), float32(pos_y), float32(pos_x)-float32(x2), float32(pos_y)-float32(y2), 1, color_red, true)
	if ebiten.IsKeyPressed(ebiten.KeyF2) {

		ebitenutil.DebugPrintAt(s, fmt.Sprintf("1: (%v, %v)\n2: (%v, %v)", x1, y1, x2, y2), 0, 550)
		ebitenutil.DebugPrintAt(s, fmt.Sprintf("cam: (%v)", camera.Angle_yaw), 0, 600)
	}
	// display objects
	camera_x, camera_z := camera.Position.X, camera.Position.Z

	for i, obj := range objects {
		obj_x, obj_z := obj.Origin.X, obj.Origin.Z

		scaledX, scaledY := (float64(pos_x) - (obj_x - camera_x)), (float64(pos_y) + (obj_z - camera_z))

		//	TODO
		// filter obj out of bounds

		scaled := Point_2d{scaledX, scaledY}
		origin := Point_2d{float64(pos_x), float64(pos_y)}

		rotated := RotatePoint(scaled, origin, camera.Angle_yaw-deg2rad(90))

		// flip the detected enemy to match rotation direction
		offset_x := rotated.X - float64(pos_x)

		distance_x := (float64(pos_x) - origin.X)
		// distance_y := (float64(pos_y) - origin.Y)

		rotated.X -= distance_x * 2
		// rotated.Y -= distance_y * 2
		rotated.Y += 20

		check_out_point := rotated
		check_out_point.X = rotated.X - origin.X
		check_out_point.Y = rotated.Y - origin.Y

		if is_out_of_circle(check_out_point, r) {
			continue
		}

		vector.DrawFilledCircle(s, float32(origin.X+offset_x), float32(rotated.Y), 2, color.RGBA{180, 0, 255, 255}, true)
		if ebiten.IsKeyPressed(ebiten.KeyF2) {
			ebitenutil.DebugPrintAt(s, fmt.Sprintf("Scaled object: X: %v, Z: %v, id: %v", float32(origin.X+offset_x), float32(rotated.Y), i), 0, 230+i*50)
		}
	}

}

func Debug_map(s *ebiten.Image, objects []Object_3d, camera *Camera) {
	//new radar
	radar(s, objects, camera)

	if !ebiten.IsKeyPressed(ebiten.KeyF2) {
		return
	}

	//old radar
	box_x, box_y := s.Bounds().Dx()/5, s.Bounds().Dy()/5
	centerX, centerY := box_x/2, box_y/2

	vector.StrokeRect(s, 0, 0, float32(box_x), float32(box_y), 2, color.RGBA{0, 255, 0, 255}, true)

	camera_x, camera_y, camera_z := camera.Position.X, camera.Position.Y, camera.Position.Z

	// vector.DrawFilledCircle(s, float32(camera_x), float32(camera_z), 2, color.RGBA{255, 0, 0, 255}, true)
	vector.DrawFilledCircle(s, float32(centerX), float32(centerY), 2, color.RGBA{0, 255, 0, 255}, true)

	if ebiten.IsKeyPressed(ebiten.KeyF2) {
		ebitenutil.DebugPrintAt(s, fmt.Sprintf("X: %v, Y: %v, Z: %v", camera_x, camera_y, camera_z), 0, 200)
	}
	//Camera centered mapping

	for _, obj := range objects {
		obj_x, obj_z := obj.Origin.X, obj.Origin.Z

		scaledX, scaledY := (float64(centerX) - (obj_x - camera_x)), (float64(centerY) + (obj_z - camera_z))

		if scaledX > float64(box_x) {
			continue
		}
		if scaledY > float64(box_y) {
			continue
		}

		vector.DrawFilledCircle(s, float32(scaledX), float32(scaledY), 2, color.RGBA{0, 0, 255, 255}, true)
		if ebiten.IsKeyPressed(ebiten.KeyF2) {
			// ebitenutil.DebugPrintAt(s, fmt.Sprintf("Scaled object: X: %v, Z: %v, id: %v", scaledX, scaledY, i), 0, 230+i*50)
		}
	}

}
