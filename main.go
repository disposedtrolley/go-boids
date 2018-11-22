package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Define velocity
const velocity = 0.1

// Define world constants
const canvasWidth = 1024
const canvasHeight = 768

// Define Bird shape constants
const birdDefaultEdgeLength = 30.0
const birdDefaultBaseLength = 20.0

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
// pointed at a specified direction and positioned at
// a coordinate on the canvas.
// The direction starts at 0 for North, and increases
// clockwise to 359. The coordinate is comprised of
// an X and Y value, bounded to the size of the canvas.
func generateBird(direction int, xCoord float64, yCoord float64) *imdraw.IMDraw {
	fmt.Printf("xCoord: %f, yCoord: %f\n", xCoord, yCoord)
	bird := imdraw.New(nil)
	bird.Color = colornames.Black
	bird.Push(pixel.V(xCoord, yCoord))
	bird.Push(pixel.V(xCoord+birdDefaultBaseLength, yCoord))
	bird.Push(pixel.V(xCoord+(birdDefaultBaseLength/2), yCoord+birdDefaultEdgeLength))
	bird.Polygon(1)

	return bird
}

func simulationLoop() {
	cfg := pixelgl.WindowConfig{
		Title:  "go-boids",
		Bounds: pixel.R(0, 0, canvasWidth, canvasHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	centerWindow(win)

	x := 0.0
	y := 0.0

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		bird := generateBird(0, x, y)
		bird.Draw(win)
		win.Update()

		x += velocity
		y += velocity
	}
}

func main() {
	pixelgl.Run(simulationLoop)
}
