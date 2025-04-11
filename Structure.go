package main

type Point struct {
	X int
	Y int
}

type Game struct {
	points []Point
	mode   int // 1 = по одной точке, 10 = пачкой
}
