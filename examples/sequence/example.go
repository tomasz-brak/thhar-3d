package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	renderer "github.com/tomasz-brak/thhar-3d/thhar"
)

type Game struct{}

// Objects is a slice of 3d objects that will be rendered.
var Objects []renderer.Object_3d

// Camera is the camera that will be used to render the 3d objects.
var Camera1 renderer.Camera
var Camera2 renderer.Camera

var current_camera = &Camera1
var timer = 200

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		current_camera = &Camera1
	}
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		current_camera = &Camera2
	}
	timer--
	if timer <= 0 {
		timer = 200
		if current_camera == &Camera1 {
			current_camera = &Camera2
			return nil
		}
		current_camera = &Camera1
	}
	return nil
}

// Game loop
func (g *Game) Draw(screen *ebiten.Image) {
	// Rendering 3d objects
	renderer.Render(screen, Objects, current_camera, color.White)

	// Changing the camera used for rendering based on a key press.
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Press R to use Camera1, T to use Camera2, automatic switch every 200 Ticks, next switch in: %v", timer))

	// Provides debug view while pressing F2, can be omitted
	renderer.Debug_map(screen, Objects, current_camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func main() {
	fmt.Printf("Current os is: %v", runtime.GOOS)

	ebiten.SetWindowSize(1600, 900)
	// Setting up 2 new cameras in the 3d space. Setting their positions to different values.
	Camera1 = *renderer.New_camera(renderer.Point_3d{X: 0, Y: -1, Z: 10})
	Camera2 = *renderer.New_camera(renderer.Point_3d{X: 10, Y: 5, Z: 0})

	// Changing the angles of the second camera
	Camera2.Angle_yaw = math.Pi * 0.4
	Camera2.Angle_pitch = math.Pi * (-0.2)

	current_camera = &Camera1

	// Importing a model from a file, appending it to global objects. || File from cwd.
	Objects = append(Objects, renderer.GetObjectFromFile("/3dModels/sequence.obj", 1))

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
