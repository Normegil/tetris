package model

type Coordinate struct {
	X, Y int
}

type Board []Square

func (t Board) DetectCollision(tetromino Tetromino) bool {
	tetroCoord := tetromino.Coordinates()

	if t.detectSquaresCollision(tetroCoord) {
		return true
	}

	return t.detectBoardCollision(tetroCoord)
}

func (t Board) detectBoardCollision(tetroCoord []Coordinate) bool {
	for _, coord := range tetroCoord {
		if coord.Y < 0 || coord.X < 0 || coord.X >= 22 {
			return true
		}
	}
	return false
}

func (t Board) detectSquaresCollision(tetroCoord []Coordinate) bool {
	for _, square := range t {
		for _, coord := range tetroCoord {
			if square.X == coord.X && square.Y == coord.Y {
				return true
			}
		}
	}
	return false
}

func (t Board) AddSquares(tetromino Tetromino) Board {
	tetroCoord := tetromino.Coordinates()
	for _, coord := range tetroCoord {
		t = append(t, Square{
			X:     coord.X,
			Y:     coord.Y,
			Color: GetColor(tetromino.Type),
		})
	}
	return t
}

func (t Board) RemoveFullLines() int {
	var lines map[int]int
	for _, square := range t {
		lines[square.Y] += 1
	}

	var nbRemovedLine int
	for lineNb, nbSquare := range lines {
		if nbSquare == 10 {
			nbRemovedLine += 1
			t = t.removeLine(lineNb)
		}
	}

	for lineNb, nbSquare := range lines {
		if nbSquare == 10 {
			for _, square := range t {
				if square.Y > lineNb {
					square.Y -= 1
				}
			}
		}
	}

	return nbRemovedLine
}

func (t Board) removeLine(y int) Board {
	for i, square := range t {
		if square.Y == y {
			t = append(t[:i], t[i+1:]...)
		}
	}
	return t
}

func (t Board) HasLost() bool {
	for _, square := range t {
		if square.Y > 20 {
			return true
		}
	}
	return false
}
