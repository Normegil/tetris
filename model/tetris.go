package model

import "time"

type Tetris struct {
	Board            Board
	CurrentTetromino Tetromino
	NextTetromino    Tetromino

	lastUpdated time.Time
}

func NewTetris() *Tetris {
	return &Tetris{
		Board: Board{},

		CurrentTetromino: Tetromino{
			Type: randomTetrominoType(),
			Center: Coordinate{
				X: 5,
				Y: 0,
			},
		},
		NextTetromino: Tetromino{
			Type: randomTetrominoType(),
			Center: Coordinate{
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
			t.handleDownCollision()
		} else {
			t.CurrentTetromino = tempTetromino
		}
		t.lastUpdated = time.Now()
	}
}

func (t Tetris) needDownMove() bool {
	gap := time.Duration(time.Now().UnixNano() - t.lastUpdated.UnixNano())
	var gapBetweenUpdate time.Duration
	if 0 == t.Level() {
		gapBetweenUpdate = time.Second / 2
	} else {
		gapBetweenUpdate = time.Second / time.Duration(t.Level())
	}

	return gap > gapBetweenUpdate
}

func (t *Tetris) Rotate(rotation Rotation) {
	temp := t.CurrentTetromino.Rotate(rotation)
	if !t.Board.DetectCollision(temp) {
		t.CurrentTetromino = temp
	}
}

func (t *Tetris) Move(direction Direction) {
	temp := t.CurrentTetromino.Move(direction)
	if !t.Board.DetectCollision(temp) {
		t.CurrentTetromino = temp
	} else if DIRECTION_DOWN == direction {
		t.handleDownCollision()
	}
}

func (t *Tetris) handleDownCollision() {
	t.Board = t.Board.AddSquares(t.CurrentTetromino)
	t.Board.RemoveFullLines()
	t.CurrentTetromino = t.NextTetromino
	t.NextTetromino = Tetromino{
		Type: randomTetrominoType(),
		Center: Coordinate{
			X: 5,
			Y: 0,
		},
	}
}

func (t Tetris) Level() int {
	return t.Board.RemovedLines / 10
}

func (t Tetris) HasLost() bool {
	return t.Board.HasLost()
}
