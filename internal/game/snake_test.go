package game

import (
	"testing"

	"go-snake/internal/coords"
)

func TestSnake_generateFoodPosition(t *testing.T) {
	s := Snake{coords: coords.NewQueue()}
	s.coords.Push(coords.Coords{X: 0, Y: 0})
	s.coords.Push(coords.Coords{X: 1, Y: 0})

	s.generateFoodPosition()
}
