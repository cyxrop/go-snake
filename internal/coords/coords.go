package coords

type Coords struct {
	X, Y int
}

func (c Coords) IsEqual(other Coords) bool {
	return c.X == other.X && c.Y == other.Y
}
