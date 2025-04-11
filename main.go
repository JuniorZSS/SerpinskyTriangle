package main

import "fmt"

func main() {
	fmt.Println("Выберите статическое создание треугольника Серпинского, или динамическое, сами будете контроллировать ввод")
	fmt.Print("Введите - 1, для статического выбора и 2 для динамического: ")
	var a int
	fmt.Scan(&a)
	if a == 1 {
		StaticPainting()
	} else if a == 2 {
		//DinamicPaitingGame()
	}
}
