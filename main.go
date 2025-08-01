package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
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
	origin_x:   100,
	origin_y:   0,
	width:      100,
	height:     100,
	velocity_x: float32(rand.Intn(5) + 1),
	velocity_y: float32(rand.Intn(5) + 1),
	color:      randomColors[rand.Intn(20)],
	antialias:  true,
}

func (g *Game) Update() error {
	// now := time.Now().UnixMilli()
	// if now-timer > ms && IND < len(toShow) {
	// 	msg += string(toShow[IND])
	// 	IND++
	// 	timer = now
	// }
	if cube.origin_x+cube.width+cube.velocity_x >= 900 || cube.origin_x+cube.velocity_x <= 0 {
		cube.color = randomColors[rand.Intn(20)]
		cube.velocity_x = -cube.velocity_x
	}

	if cube.origin_y+cube.height+cube.velocity_y >= 600 || cube.origin_y+cube.velocity_y <= 0 {
		cube.color = randomColors[rand.Intn(20)]
		cube.velocity_y = -cube.velocity_y
	}

	cube.origin_x += cube.velocity_x
	cube.origin_y += cube.velocity_y
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, msg)
	vector.DrawFilledRect(screen, cube.origin_x, cube.origin_y, cube.width, cube.height, cube.color, cube.antialias)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 900, 600
}

func main() {
	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowTitle("Bouncer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
