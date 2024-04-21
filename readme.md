# Thhar

Thhar is a 3d rendering engine written in go. Based on [ebiten](https://ebitengine.org/), it allows for rendering complex shapes. It supports [Wavefront's `.obj` format](https://en.wikipedia.org/wiki/Wavefront_.obj_file). ![cube](https://i.imgur.com/nNJsauP.png)

## Live demos

Demos can be found on my blog [https://blog-per.vercel.app/demos](https://blog-per.vercel.app/demos), check the rendering with thhar section

## Usage

Only requirement is to put a `.Render` function in any [ebiten](https://ebitengine.org/) application's main loop. Calling it with our camera instance and passing in details about the color of rendered lines.
```go
func (g *Game) Draw(screen *ebiten.Image) {
	// Camera movement code is included in the library. If you want to use your own code for that you can by modifying the "origin" of the camera.
	renderer.HandleMovement(&Camera)
	// Rendering 3d objects
	renderer.Render(screen, Objects, &Camera, color.White)
}
``` 
*extract from [./examples/basic/example.go](./examples/basic/example.go)* 

Handling movement can be done using the `.HandleMovement` function (this uses default ebiten key press detection, [Keybinds](#Keybinds)), passing by the `Camera` every frame. If more customization is needed the `Camera`'s angle and position variables are exposed and can be used to adjust or remove completely the ability to move the camera.

Multiple cameras can be used to performed sequenced actions ([**check examples**](./examples/sequence/)). Even multiple `.Render` calls per frame are accepted. *Be careful to not render two objects with different cameras in the same frame, or some weird visuals may be seen* 

## Sources 

The engine was designed for the [Motorola Science Cup 2023/2024](https://science-cup.pl/) competition. For the replica of an old arcade classic *Battlezone*

This engine was originally based on a similar [Python](https://www.python.org/) project by [StanislavPetrovV](https://github.com/StanislavPetrovV/Software_3D_engine) and another project created by [tvytlx](https://github.com/tvytlx/render-py)


## Keybinds
*All movement is relative to cameras rotation and position*
| Key      | Action              |
| -------- | ------------------- |
| KeyA     | Move left           |
| KeyD     | Move right          |
| KeyW     | Move forward        |
| KeyS     | Move backwards      |
| KeySpace | Move up             |
| KeyShift | Move down           |
| KeyLeft  | Rotate camera left  |
| KeyRight | Rotate camera right |
| KeyUp    | Rotate camera up    |
| KeyDown  | Rotate camera down  |

## Development

This is a hobby project, open to contributions! **Expect progress based on my free time**.
