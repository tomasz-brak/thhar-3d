package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	renderer "github.com/tomasz-brak/thhar-3d/thhar"
)

type Game struct{}

// Objects is a slice of 3d objects that will be rendered.
var Objects []renderer.Object_3d

// Camera is the camera that will be used to render the 3d objects.
var Camera renderer.Camera

// Ebiten Boilerplate
func (g *Game) Update() error {
	// Animation is as simple as changing the variables here.
	Objects[0].Rotation.X += 0.01
	Objects[0].Rotation.Y += 0.01
	Objects[0].Rotation.Z += 0.01

	return nil
}

// Game loop
func (g *Game) Draw(screen *ebiten.Image) {
	// Camera movement code is included in the library. If you want to use your own code for that you can by modifying the "origin" of the camera.
	renderer.HandleMovement(&Camera)
	// Rendering 3d objects
	renderer.Render(screen, Objects, &Camera, color.White)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %v, use arrows and WASD to move around", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	// Setting up a new camera in the 3d space, you can setup multiple with different angles and positions. The camera passed to the Render function will be used.
	Camera = *renderer.New_camera(renderer.Point_3d{X: 0, Y: 0, Z: 10})

	// Importing a model from a file, appending it to global objects. || File from cwd.
	Objects = append(Objects, renderer.GetObjectFromFile("/3dModels/basic.obj", 1))

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
