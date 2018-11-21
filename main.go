package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Define Bird shape constants
var BirdDefaultEdgeLength = 30.0
var BirdDefaultBaseLength = 20.0

func centerWindow(win *pixelgl.Window) {
	x, y := pixelgl.PrimaryMonitor().Size()
	width, height := win.Bounds().Size().XY()
	win.SetPos(
		pixel.V(
			x/2-width/2,
			y/2-height/2,
		),
	)
}

// Generates a new IMDraw object representing a bird,
// pointed at a specified direction. The direction
// starts at 0 for North, and increases clockwise to
// 359.
func generateBird(direction int) *imdraw.IMDraw {
	bird := imdraw.New(nil)
	bird.Color = colornames.Black
	bird.Push(pixel.V(0, 0))
	bird.Push(pixel.V(BirdDefaultBaseLength, 0))
	bird.Push(pixel.V(BirdDefaultBaseLength/2, BirdDefaultEdgeLength))
	bird.Polygon(1)

	return bird
}

func simulationLoop() {
	cfg := pixelgl.WindowConfig{
		Title:  "go-boids",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	centerWindow(win)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		bird := generateBird(0)
		bird.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(simulationLoop)
}
