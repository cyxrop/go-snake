package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func RenderArena() fyne.CanvasObject {
	var (
		arenaTop = &canvas.Line{
			StrokeColor: theme.ForegroundColor(),
			StrokeWidth: 2,
			Position1:   fyne.NewPos(padding, padding),
			Position2:   fyne.NewPos(size+padding, padding),
		}
		arenaLeft = &canvas.Line{
			StrokeColor: theme.ForegroundColor(),
			StrokeWidth: 2,
			Position1:   fyne.NewPos(padding, padding),
			Position2:   fyne.NewPos(padding, size+padding),
		}
		arenaBottom = &canvas.Line{
			StrokeColor: theme.ForegroundColor(),
			StrokeWidth: 2,
			Position1:   fyne.NewPos(padding, size+padding),
			Position2:   fyne.NewPos(size+padding, size+padding),
		}
		arenaRight = &canvas.Line{
			StrokeColor: theme.ForegroundColor(),
			StrokeWidth: 2,
			Position1:   fyne.NewPos(size+padding, padding),
			Position2:   fyne.NewPos(size+padding, size+padding),
		}
	)

	return container.NewWithoutLayout(arenaTop, arenaLeft, arenaBottom, arenaRight)
}
