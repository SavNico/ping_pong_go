package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	x, y, speed, witdh, height int32
}

type PlayerPaddle struct {
	paddle Paddle
}

type CpuPaddle struct {
	paddle Paddle
}

func (paddle *Paddle) DrawPaddle() {
	rl.DrawRectangle(paddle.x, paddle.y, paddle.witdh, paddle.height, rl.White)
}

func (paddle *Paddle) limitMovement() {
	if paddle.y <= 0 {
		paddle.y = 0
	}
	if paddle.y+paddle.height >= int32(rl.GetScreenHeight()) {
		paddle.y = int32(rl.GetScreenHeight()) - paddle.height
	}
}

// Player logic
func (pp *PlayerPaddle) Update() {

	if rl.IsKeyDown(rl.KeyUp) {
		pp.paddle.y = pp.paddle.y - pp.paddle.speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		pp.paddle.y = pp.paddle.y + pp.paddle.speed
	}

	pp.paddle.limitMovement()
}

// AI logic
func (cp *CpuPaddle) Update(ball_y int32) {
	if cp.paddle.y+cp.paddle.height/2 > ball_y {
		cp.paddle.y = cp.paddle.y - cp.paddle.speed
	}
	if cp.paddle.y+cp.paddle.height/2 <= ball_y {
		cp.paddle.y = cp.paddle.y + cp.paddle.speed
	}

	cp.paddle.limitMovement()
}
