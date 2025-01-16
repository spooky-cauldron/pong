package main

// #include "raylib.h"
import "C"

type Paddle struct {
	Collider
	Speed int
}

func (p *Paddle) Render() {
	C.DrawRectangle(C.int(p.X), C.int(p.Y), C.int(p.Width), C.int(p.Height), C.GREEN)
}

func (p *Paddle) Move() {
	if C.IsKeyDown(C.KEY_UP) {
		p.Collider.Y -= int(DeltaTime() * float64(p.Speed))
	}

	if C.IsKeyDown(C.KEY_DOWN) {
		p.Collider.Y += int(DeltaTime() * float64(p.Speed))
	}

	if p.Collider.Y < 0 {
		p.Collider.Y = 0
	} else if p.Collider.Y > ScreenHeight()-p.Height {
		p.Collider.Y = ScreenHeight() - p.Height
	}
}
