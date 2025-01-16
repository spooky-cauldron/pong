package main

// #include "raylib.h"
import "C"

type Vec2 struct {
	X int
	Y int
}

// type Collider interface {
// 	GetColliderInfo() ColliderInfo
// }

type Collider struct {
	X         int
	Y         int
	Height    int
	Width     int
	BounceDir Vec2
}

func Time() float64 {
	return float64(C.GetTime())
}

func DeltaTime() float64 {
	return float64(C.GetFrameTime())
}

// func CheckCollision(ball *Ball, wall *Wall) {
func CheckCollision(ball *Ball, collider Collider) {
	center := C.Vector2{x: C.float(ball.X), y: C.float(ball.Y)}
	rect := C.Rectangle{
		x:      C.float(collider.X),
		y:      C.float(collider.Y),
		height: C.float(collider.Height),
		width:  C.float(collider.Width),
	}
	colliding := C.CheckCollisionCircleRec(center, C.float(ball.Radius), rect)
	if colliding {
		ball.Vx *= collider.BounceDir.X
		ball.Vy *= collider.BounceDir.Y
	}
}
