package main

import (
	"github.com/fogleman/gg"
	"math/rand"
	"time"
)

func StaticPainting() {
	const size = 800
	const points = 100000

	dc := gg.NewContext(size, size)

	// белый фон
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// чёрные точки
	dc.SetRGB(0, 0, 0)

	// вершины треугольника
	a := Point{400, 100}
	b := Point{100, 700}
	c := Point{700, 700}

	// массив вершин
	vertices := []Point{a, b, c}

	// начальная точка (можно любую)
	current := Point{400, 400}

	// Рандомайзер
	rand.Seed(time.Now().UnixNano())

	//итерации
	for i := 0; i < points; i++ {
		v := vertices[rand.Intn(3)]
		current = Point{
			(current.X + v.X) / 2,
			(current.Y + v.Y) / 2,
		}
		dc.SetPixel(int(current.X), int(current.Y))
	}

	// сохраняем результат
	dc.SavePNG("sierpinski.png")
}
