package model

import "github.com/Sirupsen/logrus"

type Coordinate struct {
	X, Y int
}

type Board struct {
	Squares      []Square
	RemovedLines int
}

func (t Board) DetectCollision(tetromino Tetromino) bool {
	tetroCoord := tetromino.Coordinates()

	if t.detectSquaresCollision(tetroCoord) {
		return true
	}

	return t.detectBoardCollision(tetroCoord)
}

func (t Board) detectBoardCollision(tetroCoord []Coordinate) bool {
	for _, coord := range tetroCoord {
		if coord.Y > 21 || coord.X < 0 || coord.X >= 10 {
			return true
		}
	}
	return false
}

func (t Board) detectSquaresCollision(tetroCoord []Coordinate) bool {
	for _, square := range t.Squares {
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
		t.Squares = append(t.Squares, Square{
			X:     coord.X,
			Y:     coord.Y,
			Color: GetColor(tetromino.Type),
		})
	}
	return t
}

func (t *Board) RemoveFullLines() {
	lines := make(map[int]int)
	for _, square := range t.Squares {
		lines[square.Y] += 1
	}

	logrus.Debug("TAP")

	var nbRemovedLine int
	for lineNb, nbSquare := range lines {
		if nbSquare == 10 {
			nbRemovedLine += 1
			t.Squares = t.removeLine(lineNb)
			t.RemovedLines += 1

			logrus.WithField("Line", lineNb).Debug("Line Removed")
			for i := range t.Squares {
				logrus.WithField("Square", t.Squares[i]).Debug("Before Removing")
				if t.Squares[i].Y < lineNb {
					t.Squares[i].Y += 1
				}
				logrus.WithField("Square", t.Squares[i]).Debug("After Removing")
			}
		}
	}
}

func (t Board) removeLine(y int) []Square {
	var board []Square
	for _, square := range t.Squares {
		if square.Y != y {
			board = append(board, square)
		}
	}
	return board
}

func (t Board) HasLost() bool {
	for _, square := range t.Squares {
		if square.Y < 2 {
			return true
		}
	}
	return false
}
