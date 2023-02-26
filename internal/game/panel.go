package game

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Panel struct {
	scoreLabel *widget.Label
	retryBtn   *widget.Button
	infoLabel  *widget.Label

	snake *Snake
}

func NewPanel(snake *Snake) *Panel {
	p := &Panel{
		snake:     snake,
		infoLabel: widget.NewLabel(""),
	}
	p.retryBtn = widget.NewButton("Retry", p.handleRetry)
	p.scoreLabel = widget.NewLabel(formatScore(p.snake.Score()))
	return p
}

func (p *Panel) Canvas() *fyne.Container {
	return container.NewGridWithColumns(3, p.scoreLabel, p.retryBtn, p.infoLabel)
}

func (p *Panel) Run(events <-chan Event) {
	go func() {
		for e := range events {
			switch e {
			case CrashedEvent:
				p.handleCrashed()
			case AteFoodEvent:
				p.handleAteFound()
			}
		}
	}()
}

func (p *Panel) handleCrashed() {
	p.infoLabel.SetText("Boom!!!")
}

func (p *Panel) handleAteFound() {
	p.scoreLabel.SetText(formatScore(p.snake.Score()))
}

func (p *Panel) handleRetry() {
	p.infoLabel.SetText("")
	p.snake.Start()
	p.scoreLabel.SetText(formatScore(p.snake.Score()))
}

func formatScore(score int) string {
	return fmt.Sprintf("Score: %d", score)
}
