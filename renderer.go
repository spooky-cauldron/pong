package main

// #include "raylib.h"
// #include "ray.h"
import "C"

type Renderer interface {
	Render()
}

func ScreenWidth() int {
	if C.GetScreenWidth() == 0 {
		return 800
	}
	return int(C.GetScreenWidth())
}

func ScreenHeight() int {
	if C.GetScreenHeight() == 0 {
		return 450
	}
	return int(C.GetScreenHeight())
}
