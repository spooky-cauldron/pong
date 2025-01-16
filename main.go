package main

// #cgo LDFLAGS: -lraylib
// #include "raylib.h"
// #include "ray.h"
import "C"

var ball *Ball = &Ball{Radius: 20, X: 100, Y: 100, Vx: 300, Vy: 300}

const paddleHeight = 100

var paddle *Paddle = &Paddle{
	Collider: Collider{
		X:         20,
		Y:         ScreenHeight()/2 - paddleHeight/2,
		Width:     10,
		Height:    paddleHeight,
		BounceDir: Vec2{-1, 1},
	},
	Speed: 400,
}

var walls []*Wall

var Renderers []Renderer

var GameOver = false

func main() {
	walls = BuildWalls()
	Renderers = []Renderer{ball, paddle}
	for _, wall := range walls {
		Renderers = append(Renderers, wall)
	}
	C.renderLoop()
}

//export GoLogicLoop
func GoLogicLoop() {
	for _, wall := range walls {
		CheckCollision(ball, wall.Collider)
	}
	CheckCollision(ball, paddle.Collider)
	ball.Move()
	paddle.Move()

	if ball.X < 0 {
		GameOver = true
	}
}

//export GoRenderLoop
func GoRenderLoop() {
	for _, renderer := range Renderers {
		renderer.Render()
	}

	if GameOver {
		C.DrawText(C.CString("Game Over"), C.int(ScreenWidth()/2), C.int(ScreenHeight()/2), 20, C.GREEN)
	}
}

func BuildWalls() []*Wall {
	wallThickness := 10
	upWall := &Wall{
		Collider: Collider{
			X:         0,
			Y:         0,
			Width:     ScreenWidth(),
			Height:    wallThickness,
			BounceDir: Vec2{X: 1, Y: -1},
		},
	}
	downWall := &Wall{
		Collider: Collider{
			X:         0,
			Y:         ScreenHeight() - wallThickness,
			Width:     ScreenWidth(),
			Height:    wallThickness,
			BounceDir: Vec2{X: 1, Y: -1},
		},
	}
	rightWall := &Wall{
		Collider: Collider{
			X:         ScreenWidth() - wallThickness,
			Y:         0,
			Width:     wallThickness,
			Height:    ScreenHeight(),
			BounceDir: Vec2{X: -1, Y: 1},
		},
	}
	walls := []*Wall{upWall, downWall, rightWall}
	return walls
}
