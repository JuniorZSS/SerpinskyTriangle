package main

import (
	"github.com/fogleman/gg"
	"math/rand"
	"time"
)

func BransliPaporotnik() {
	const size = 800
	const points = 100000

	dc := gg.NewContext(size, size)

	// белый фон
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// чёрные точки
	dc.SetRGB(0, 0, 0)

	x := 0.0
	y := 0.0

	// Рандомайзер
	rand.Seed(time.Now().UnixNano())

	// Цикл Брансли
	for i := 0; i < points; i++ {
		r := rand.Float64()

		var nextX, nextY float64

		switch {
		case r < 0.01:
			nextX = 0.0
			nextY = 0.16 * y
		case r < 0.86:
			nextX = 0.85*x + 0.04*y
			nextY = -0.04*x + 0.85*y + 1.6
		case r < 0.93:
			nextX = 0.20*x - 0.26*y
			nextY = 0.23*x + 0.22*y + 1.6
		default:
			nextX = -0.15*x + 0.28*y
			nextY = 0.26*x + 0.24*y + 0.44
		}

		x = nextX
		y = nextY

		// Преобразование координат в экранные
		screenX := int(400 + x*60)
		screenY := int(800 - y*60)

		dc.SetPixel(screenX, screenY)
	}

	// сохраняем результат
	dc.SavePNG("branslipaporotnik.png")
}
