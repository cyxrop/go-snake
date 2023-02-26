package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"go-snake/internal/game"
)

func main() {
	a := app.New()
	window := a.NewWindow("Snake")
	window.Resize(fyne.NewSize(630, 670))
	snake := game.NewSnake()
	panel := game.NewPanel(snake)

	window.SetContent(container.NewVBox(
		panel.Canvas(),
		container.NewWithoutLayout(
			game.RenderArena(),
			snake.Canvas(),
		),
	))

	window.Canvas().SetOnTypedKey(snake.HandleKeyEvent)
	snake.Run()
	panel.Run(snake.Events())

	window.ShowAndRun()
}
