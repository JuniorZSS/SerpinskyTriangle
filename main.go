package main

import "fmt"

func main() {
	fmt.Println("Выберите статическое создание треугольника Серпинского, или динамическое, сами будете контроллировать ввод")
	fmt.Print("Выберите какую картинку математическим вычислением хотите создать, 1 - Треугольник Серпинского, 2 - Папоротник Брансли: ")
	var a int
	fmt.Scan(&a)
	if a == 1 {
		StaticPainting()
	} else if a == 2 {
		BransliPaporotnik()
	}
}
