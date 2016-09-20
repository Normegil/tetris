package model

import (
	"time"
)

type Tetris struct {
	Board            Board
	DeletedLines     int
	CurrentTetromino Tetromino
	NextTetromino    Tetromino

	lastUpdated time.Time
}

func NewTetris() *Tetris {
	return &Tetris{
		Board: Board{},

		CurrentTetromino: Tetromino{
			Type: randomTetrominoType(),
			coordinate: Coordinate{
				X: 5,
				Y: 0,
			},
		},
		NextTetromino: Tetromino{
			Type: randomTetrominoType(),
			coordinate: Coordinate{
				X: 5,
				Y: 0,
			},
		},
		lastUpdated: time.Now(),
	}
}

func (t *Tetris) Update() {
	if t.needDownMove() {
		tempTetromino := t.CurrentTetromino.Move(DIRECTION_DOWN)
		if t.Board.DetectCollision(tempTetromino) {
			t.Board.AddSquares(t.CurrentTetromino)
			t.DeletedLines += t.Board.RemoveFullLines()
			t.CurrentTetromino = t.NextTetromino
			t.NextTetromino = Tetromino{
				Type: randomTetrominoType(),
				coordinate: Coordinate{
					X: 5,
					Y: 0,
				},
			}
		} else {
			t.CurrentTetromino = tempTetromino
		}
	}
}

func (t Tetris) needDownMove() bool {
	gap := time.Duration(time.Now().UnixNano() - t.lastUpdated.UnixNano())
	var gapBetweenUpdate time.Duration
	if 0 == t.Level() {
		gapBetweenUpdate = time.Second
	} else {
		gapBetweenUpdate = time.Second / time.Duration(t.Level())
	}

	return gap > gapBetweenUpdate
}

func (t Tetris) Level() int {
	return t.DeletedLines / 10
}

func (t Tetris) HasLost() bool {
	return t.Board.HasLost()
}
