package main

//func (g *Game) Update() error {
//	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
//		x, y := ebiten.CursorPosition()
//		if g.mode == 1 {
//			g.points = append(g.points, Point{x, y})
//		} else {
//			for i := 0; i < 10; i++ {
//				g.points = append(g.points, Point{x + i, y + i})
//			}
//		}
//	}
//
//	// переключение режима — нажми клавишу 1 или 2
//	if ebiten.IsKeyPressed(ebiten.Key1) {
//		g.mode = 1
//	}
//	if ebiten.IsKeyPressed(ebiten.Key2) {
//		g.mode = 10
//	}
//
//	return nil
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
//	screen.Fill(color.White)
//
//	for _, p := range g.points {
//		screen.Set(p.X, p.Y, color.Black)
//	}
//
//	ebitenutil.DebugPrint(screen, "Нажми 1 — рисовать по 1 точке\nНажми 2 — по 10 точек\nКликай мышью!")
//}
//
//func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
//	return 800, 800
//}
//
//func DinamicPaitingGame() {
//	ebiten.SetWindowSize(800, 800)
//	ebiten.SetWindowTitle("Рисуй точечки мышкой")
//
//	game := &Game{mode: 1}
//	if err := ebiten.RunGame(game); err != nil {
//		log.Fatal(err)
//	}
//}
