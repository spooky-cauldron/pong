package main

// #include "raylib.h"
import "C"

type Ball struct {
	X      int
	Y      int
	Vx     int
	Vy     int
	Radius int
}

func (b *Ball) Move() {
	b.X += int(float64(b.Vx) * DeltaTime())
	b.Y += int(float64(b.Vy) * DeltaTime())
}

func (b *Ball) Render() {
	C.DrawCircle(C.int(b.X), C.int(b.Y), C.float(b.Radius), C.RED)
}
