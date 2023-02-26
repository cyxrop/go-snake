package game

import (
	"fmt"
	"image/color"
	"image/color/palette"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"go-snake/internal/coords"
)

var (
	mapKeyNameToDir = map[fyne.KeyName]direction{
		fyne.KeyDown:  downDir,
		fyne.KeyUp:    upDir,
		fyne.KeyLeft:  leftDir,
		fyne.KeyRight: rightDir,
	}
	mapReverseDir = map[direction]fyne.KeyName{
		downDir:  fyne.KeyUp,
		upDir:    fyne.KeyDown,
		leftDir:  fyne.KeyRight,
		rightDir: fyne.KeyLeft,
	}
	mapSpeedToTicks = map[int]int64{
		1:  10,
		2:  9,
		3:  8,
		4:  7,
		5:  6,
		6:  5,
		7:  4,
		8:  3,
		9:  2,
		10: 1,
	}
)

const (
	padding   = 10
	size      = 600
	cellSize  = 30
	cellCount = int(size / cellSize)
)

type Snake struct {
	mx        *sync.RWMutex
	coords    *coords.Queue
	canvas    *fyne.Container
	foodCoors *coords.Coords

	score                 int64
	speed                 int64
	direction             direction
	isDirectionChangeable bool

	eventCh chan Event
}

func NewSnake() *Snake {
	s := &Snake{
		mx:      &sync.RWMutex{},
		canvas:  container.NewWithoutLayout(),
		eventCh: make(chan Event, 1),
	}

	s.Start()
	return s
}

func (s *Snake) Start() {
	s.coords = coords.NewQueue()
	s.coords.Push(coords.Coords{X: 0, Y: 0})
	s.coords.Push(coords.Coords{X: 1, Y: 0})
	s.setScore(0)
	s.setSpeed(5)
	s.setDirection(rightDir)
	s.setIsDirectionChangeable(true)
	s.createFood()
	s.render()
}

func (s *Snake) Canvas() fyne.CanvasObject {
	return s.canvas
}

func (s *Snake) Events() <-chan Event {
	return s.eventCh
}

func (s *Snake) Score() int {
	return int(atomic.LoadInt64(&s.score))
}

func (s *Snake) Speed() int {
	return int(atomic.LoadInt64(&s.speed))
}

func (s *Snake) Run() {
	var (
		ticker       = time.NewTicker(time.Millisecond * 50)
		ticks  int64 = 0
	)

	go func() {
		for {
			<-ticker.C

			if s.getDirection() == noneDir {
				continue
			}

			atomic.AddInt64(&ticks, 1)
			ticksMax := mapSpeedToTicks[s.Speed()]
			if ticks >= ticksMax {
				s.moveSnake()
				s.render()
				atomic.StoreInt64(&ticks, 0)
			}
		}
	}()
}

func (s *Snake) moveSnake() {
	head := s.coords.PeekLast()

	var (
		x = head.X
		y = head.Y
	)

	s.mx.RLock()
	switch s.direction {
	case leftDir:
		x--
	case upDir:
		y--
	case rightDir:
		x++
	case downDir:
		y++
	}
	s.mx.RUnlock()

	newHead := coords.Coords{X: x, Y: y}
	if !s.isNextMoveValid(newHead) {
		s.setDirection(noneDir)
		s.eventCh <- CrashedEvent
		return
	}

	if newHead.IsEqual(*s.foodCoors) {
		s.eventCh <- AteFoodEvent
		s.increaseScore(1)
		s.createFood()

		if s.Score()%5 == 0 && s.Speed() < 10 {
			s.increaseSpeed(1)
		}
	} else {
		s.coords.Pop()
	}

	s.coords.Push(newHead)
	s.setIsDirectionChangeable(true)
}

func (s *Snake) render() {
	s.mx.RLock()
	defer s.mx.RUnlock()

	// Render snake
	rects := make([]fyne.CanvasObject, 0, len(s.coords.PeekAll()))
	for _, c := range s.coords.PeekAll() {
		var (
			posX = positionByCoordinate(c.X)
			posY = positionByCoordinate(c.Y)
		)
		rects = append(rects, createSquare(posX, posY, color.White))
	}

	// Render food
	rects = append(rects, createSquare(
		positionByCoordinate(s.foodCoors.X),
		positionByCoordinate(s.foodCoors.Y),
		palette.WebSafe[5],
	))

	s.canvas.Objects = rects
	s.canvas.Refresh()
}

func (s *Snake) HandleKeyEvent(event *fyne.KeyEvent) {
	if event == nil {
		return
	}

	if !s.getIsDirectionChangeable() {
		return
	}

	curDir := s.getDirection()
	if curDir == noneDir {
		return
	}

	// Ignore reversed direction
	reversedDir := mapReverseDir[curDir]
	if event.Name == reversedDir {
		return
	}

	newDir, ok := mapKeyNameToDir[event.Name]
	if !ok {
		return
	}

	s.setDirection(newDir)
	s.setIsDirectionChangeable(false)
}

func (s *Snake) increaseScore(delta int64) {
	atomic.AddInt64(&s.score, delta)
}

func (s *Snake) increaseSpeed(delta int64) {
	atomic.AddInt64(&s.speed, delta)
	fmt.Println("increase speed")
}

func (s *Snake) setScore(score int64) {
	atomic.StoreInt64(&s.score, score)
}

func (s *Snake) setSpeed(speed int64) {
	atomic.StoreInt64(&s.speed, speed)
}

func (s *Snake) getDirection() direction {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.direction
}

func (s *Snake) setDirection(dir direction) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.direction = dir
}

func (s *Snake) getIsDirectionChangeable() bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.isDirectionChangeable
}

func (s *Snake) setIsDirectionChangeable(flag bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.isDirectionChangeable = flag
}

func (s *Snake) createFood() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.foodCoors = s.generateFoodPosition()
}

func (s *Snake) generateFoodPosition() *coords.Coords {
	coordsMap := make(map[coords.Coords]struct{}, len(s.coords.PeekAll()))
	for _, c := range s.coords.PeekAll() {
		coordsMap[c] = struct{}{}
	}

	freeCoords := make([]coords.Coords, 0)
	for x := 0; x < cellCount; x++ {
		for y := 0; y < cellCount; y++ {
			c := coords.Coords{X: x, Y: y}
			if _, ok := coordsMap[c]; !ok {
				freeCoords = append(freeCoords, c)
			}
		}
	}

	randIdx := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(freeCoords))
	c := freeCoords[randIdx]
	return &c
}

func (s *Snake) isNextMoveValid(next coords.Coords) bool {
	if isOutOfArena(next.X) || isOutOfArena(next.Y) {
		return false
	}

	s.mx.RLock()
	defer s.mx.RUnlock()
	for _, c := range s.coords.PeekAll() {
		if c.IsEqual(next) {
			return false
		}
	}
	return true
}

func positionByCoordinate(c int) float32 {
	return float32(padding + cellSize*c)
}

func createSquare(x, y float32, color color.Color) *canvas.Rectangle {
	rect := canvas.NewRectangle(color)
	rect.Resize(fyne.NewSize(cellSize, cellSize))
	rect.Move(fyne.NewPos(x, y))
	return rect
}

func isOutOfArena(c int) bool {
	return c < 0 || c >= cellCount
}
