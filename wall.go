package main

// #include "raylib.h"
import "C"

type Wall struct {
	Collider
	// Width     int
	// Height    int
	// X         int
	// Y         int
	// BounceDir Vec2
}

func (w *Wall) Render() {
	C.DrawRectangle(C.int(w.X), C.int(w.Y), C.int(w.Width), C.int(w.Height), C.BLUE)
}

// func (w *Wall) GetColliderInfo() ColliderInfo {
// 	return ColliderInfo {
// 		Width: w.Width,
// 		Height: w.Height,
// 		X: w.X,
// 		Y: w.Y,
// 		BounceDir: ,
// 	}
// }
