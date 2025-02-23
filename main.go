package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
)

// Define velocity
const velocity = 0.001

// Define world constants
const canvasWidth = 1024
const canvasHeight = 768

// Define Bird shape constants
const birdDefaultEdgeLength = 20.0
const birdDefaultBaseLength = 20.0

// Boid represents a bird present on the screen.
type Boid struct {
	centroid Coordinate
	points   []Coordinate
}

// Coordinate represents an x,y position on the canvas.
type Coordinate struct {
	x, y float64
}

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

// Generates a new Boid instance comprising of the vectors
// of each point and the centroid.
// The direction starts at 0 for North, and increases
// clockwise to 359. The coordinate is comprised of
// an X and Y value signalling the bottom-left corner
// of the boid.
func generateBoid(direction float64, coords Coordinate) Boid {
	fmt.Printf("xCoord: %f, yCoord: %f\n", coords.x, coords.y)

	// Calculate vectors for each point of the triange
	bottomLeft := Coordinate{coords.x, coords.y}
	bottomRight := Coordinate{coords.x + birdDefaultBaseLength, coords.y}
	top := Coordinate{coords.x + (birdDefaultBaseLength / 2), coords.y + birdDefaultEdgeLength}

	// Calculate the coordinates for the centre of the triangle
	centre := Coordinate{(bottomLeft.x + bottomRight.x + top.x) / 3, (bottomLeft.y + bottomRight.y + top.y) / 3}

	fmt.Printf("centre: %+v\n", centre)

	// Revise the vectors for each point, taking into account the supplied
	// direction.
	points := []Coordinate{bottomLeft, bottomRight, top}

	fmt.Println(points)

	for i := range points {
		point := &points[i]

		point.x = math.Cos(direction)*point.x + math.Sin(direction)*point.y + (centre.x - centre.x*math.Cos(direction) + centre.y*math.Sin(direction))
		point.y = -math.Sin(direction)*point.x + math.Cos(direction)*point.y + (centre.y + centre.x*math.Sin(direction) - centre.y*math.Cos(direction))
	}

	fmt.Println(points)

	return Boid{centroid: centre, points: points}
}

func drawBoid(boid Boid, win *pixelgl.Window) {
	boidRender := imdraw.New(nil)
	boidRender.Color = colornames.Black

	for _, point := range boid.points {
		boidRender.Push(pixel.V(point.x, point.y))
	}

	boidRender.Polygon(1)

	boidRender.Draw(win)
}

func degreesToRadians(degree float64) float64 {
	return degree * (math.Pi / 180.0)
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

	x := float64(canvasWidth / 2)
	y := 300.0

	angle := 0.0

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		boid := generateBoid(degreesToRadians(angle), Coordinate{x: x, y: y})
		drawBoid(boid, win)

		win.Update()

		// x += velocity
		// y += velocity
		angle += velocity * 10
	}
}

func main() {
	pixelgl.Run(simulationLoop)
}
