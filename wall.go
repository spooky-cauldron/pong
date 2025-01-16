package main

// #include "raylib.h"
import "C"

type Wall struct {
	Collider
}

func (w *Wall) Render() {
	C.DrawRectangle(C.int(w.X), C.int(w.Y), C.int(w.Width), C.int(w.Height), C.BLUE)
}
