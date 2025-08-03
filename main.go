package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

type Cube struct {
	origin_x   float32
	origin_y   float32
	width      float32
	height     float32
	velocity_x float32
	velocity_y float32
	color      color.Color
	antialias  bool
}
type Rect struct {
	Top, Bottom, Left, Right float32
}

func (c Cube) getBoundingClientRect() Rect {
	return Rect{
		Top:    c.origin_y,
		Bottom: c.origin_y + c.height,
		Left:   c.origin_x,
		Right:  c.origin_x + c.width,
	}
}

func (c Cube) doesIntersect(c1 Cube) bool {
	a := c.getBoundingClientRect()
	b := c1.getBoundingClientRect()
	if a.Left < b.Right &&
		a.Right > b.Left &&
		a.Top < b.Bottom &&
		a.Bottom > b.Top {
		return true
	}
	return false
}

// var msg = ""
// var toShow = "Hello World"
// var IND = 0
// var timer int64 = 0
// var ms int64 = 5000
func getRandomColors(n int) []color.Color {
	colors := make([]color.Color, n)

	for i := 0; i < n; i++ {
		colors[i] = color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}
	}

	return colors
}

var randomColors = getRandomColors(20)
var cube = Cube{
	origin_x:   0,
	origin_y:   550,
	width:      50,
	height:     50,
	velocity_x: 15,
	velocity_y: float32(rand.Intn(5) + 1),
	color:      randomColors[rand.Intn(20)],
	antialias:  true,
}

var obstacle = Cube{
	origin_x:   float32(rand.Intn(900)),
	origin_y:   0,
	width:      50,
	height:     100,
	velocity_x: 20,
	velocity_y: float32(rand.Intn(5) + 1),
	color:      randomColors[rand.Intn(20)],
	antialias:  true,
}
var gameOver bool
var obstacles []Cube
var score int

func createRandomObstacle() Cube {
	return Cube{
		origin_x:   float32(rand.Intn(850)),
		origin_y:   0,
		width:      50,
		height:     100,
		velocity_y: float32(rand.Intn(3) + 2),
		color:      randomColors[rand.Intn(len(randomColors))],
		antialias:  true,
	}
}
func (g *Game) Update() error {
	// now := time.Now().UnixMilli()
	// if now-timer > ms && IND < len(toShow) {
	// 	msg += string(toShow[IND])
	// 	IND++
	// 	timer = now
	// }
	if gameOver {
		return nil
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if cube.origin_x+cube.width+cube.velocity_x > 900 {
			cube.origin_x = 850
		} else {
			cube.origin_x += cube.velocity_x
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cube.origin_x -= cube.velocity_x
		if cube.origin_x < 0 {
			cube.origin_x = 0
		}
	}
	for i := range obstacles {
		obstacles[i].origin_y += obstacles[i].velocity_y
		if obstacles[i].origin_y > 600 {
			obstacles[i] = createRandomObstacle()
			score += 1
		}
		if cube.doesIntersect(obstacles[i]) {
			gameOver = true
		}
	}
	obstacle.origin_y += obstacle.velocity_y
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, msg)
	if gameOver {
		ebitenutil.DebugPrint(screen, "Game over")
		ebitenutil.DebugPrintAt(screen, strconv.Itoa(score), 870, 0)
		return
	}
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(score), 870, 0)
	vector.DrawFilledRect(screen, cube.origin_x, cube.origin_y, cube.width, cube.height, cube.color, cube.antialias)
	for _, o := range obstacles {
		vector.DrawFilledRect(screen, o.origin_x, o.origin_y, o.width, o.height, o.color, o.antialias)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 900, 600
}

func main() {
	for i := 0; i < 5; i++ {
		obstacles = append(obstacles, createRandomObstacle())
	}
	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowTitle("Stars falling")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
