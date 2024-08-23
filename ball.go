package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	x, y, speed_x, speed_y int32
	radius                 float32
}

func (ball *Ball) DrawBall() {
	rl.DrawCircle(ball.x, ball.y, ball.radius, rl.White)
}

func (ball *Ball) Update(cpu_score *int32, player_score *int32) {
	ball.x += ball.speed_x
	ball.y += ball.speed_y

	if float32(ball.y)+ball.radius >= float32(rl.GetScreenHeight()) || float32(ball.y)-ball.radius <= 0 {
		ball.speed_y *= -1
	}
	if float32(ball.x)+ball.radius >= float32(rl.GetScreenWidth()) {
		*cpu_score++
		ball.ResetBall()
	}
	if float32(ball.x)-ball.radius <= 0 {
		*player_score++
		ball.ResetBall()
	}
}

func (ball *Ball) ResetBall() {
	ball.x = int32(rl.GetScreenWidth() / 2)
	ball.y = int32(rl.GetScreenHeight() / 2)

	speed_choices := [2]int32{-1, 1}
	randomIndex := rand.Intn(len(speed_choices)) // generate a random int between -1 and 1
	ball.speed_x *= speed_choices[randomIndex]
	ball.speed_y *= speed_choices[randomIndex]
}
