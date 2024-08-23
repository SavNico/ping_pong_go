package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	//* GAME VARIABLES
	const SCREEN_WIDTH = 800
	const SCREEN_HEIGHT = 450
	const RECTANGLE_WIDTH = 15
	const RECTANGLE_HEIGHT = 80
	var player_score int32
	var cpu_score int32

	ball := Ball{
		radius:  10,
		x:       SCREEN_WIDTH / 2,
		y:       SCREEN_HEIGHT / 2,
		speed_x: 3,
		speed_y: 3,
	}

	playerPaddle := PlayerPaddle{
		paddle: Paddle{
			witdh:  RECTANGLE_WIDTH,
			height: RECTANGLE_HEIGHT,
			x:      SCREEN_WIDTH - RECTANGLE_WIDTH - 10,
			y:      SCREEN_HEIGHT/2 - RECTANGLE_HEIGHT/2,
			speed:  6,
		},
	}

	cpuPaddle := CpuPaddle{
		paddle: Paddle{
			witdh:  RECTANGLE_WIDTH,
			height: RECTANGLE_HEIGHT,
			x:      10,
			y:      SCREEN_HEIGHT/2 - RECTANGLE_HEIGHT/2,
			speed:  6,
		},
	}

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "PING-PONG en GO")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Update
		ball.Update(&cpu_score, &player_score)
		playerPaddle.Update()
		cpuPaddle.Update(ball.y)

		// Checking collisions
		if rl.CheckCollisionCircleRec(
			rl.Vector2{X: float32(ball.x), Y: float32(ball.y)},
			ball.radius,
			rl.Rectangle{
				X:      float32(playerPaddle.paddle.x),
				Y:      float32(playerPaddle.paddle.y),
				Width:  float32(playerPaddle.paddle.witdh),
				Height: float32(playerPaddle.paddle.height)},
		) {
			ball.speed_x *= -1
		}
		if rl.CheckCollisionCircleRec(
			rl.Vector2{X: float32(ball.x), Y: float32(ball.y)},
			ball.radius,
			rl.Rectangle{
				X:      float32(cpuPaddle.paddle.x),
				Y:      float32(cpuPaddle.paddle.y),
				Width:  float32(cpuPaddle.paddle.witdh),
				Height: float32(cpuPaddle.paddle.height)},
		) {
			ball.speed_x *= -1
		}

		// Draw
		rl.ClearBackground(rl.Black) // Con esto limpiamos la pantalla de lo que se renderizó anteriormente

		rl.DrawLine(SCREEN_WIDTH/2, 0, SCREEN_WIDTH/2, SCREEN_HEIGHT, rl.White)
		ball.DrawBall()
		cpuPaddle.paddle.DrawPaddle()
		playerPaddle.paddle.DrawPaddle()

		rl.DrawText(fmt.Sprint(cpu_score), (SCREEN_WIDTH/4 - 20), 20, 80, rl.White)
		rl.DrawText(fmt.Sprint(player_score), (3*SCREEN_WIDTH/4 - 20), 20, 80, rl.White)

		rl.EndDrawing()
	}
}
