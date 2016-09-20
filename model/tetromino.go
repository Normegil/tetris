package model

import "math/rand"

type TetrominoType int

const (
	TETROMINO_I TetrominoType = 0
	TETROMINO_O TetrominoType = 1
	TETROMINO_T TetrominoType = 2
	TETROMINO_L TetrominoType = 3
	TETROMINO_J TetrominoType = 4
	TETROMINO_Z TetrominoType = 5
	TETROMINO_S TetrominoType = 6
)

func randomTetrominoType() TetrominoType {
	return TetrominoType(rand.Int() % 7)
}

func GetColor(t TetrominoType) Color {
	switch t {
	case TETROMINO_O:
		return Color{R: 255, G: 255, B: 0, A: 255}
	case TETROMINO_I:
		return Color{R: 0, G: 255, B: 255, A: 255}
	case TETROMINO_J:
		return Color{R: 0, G: 0, B: 255, A: 255}
	case TETROMINO_L:
		return Color{R: 255, G: 125, B: 0, A: 255}
	case TETROMINO_Z:
		return Color{R: 255, G: 0, B: 0, A: 255}
	case TETROMINO_S:
		return Color{R: 0, G: 255, B: 0, A: 255}
	case TETROMINO_T:
		return Color{R: 255, G: 0, B: 255, A: 255}
	default:
		return Color{}
	}
}

type Rotation int

const (
	ROTATION_CLOCK     Rotation = 0
	ROTATION_ANTICLOCK Rotation = 1
)

type Direction int

const (
	DIRECTION_LEFT  Direction = 0
	DIRECTION_RIGHT Direction = 1
	DIRECTION_DOWN  Direction = 2
)

type Angle int

const (
	ANGLE_0   = 0
	ANGLE_90  = 1
	ANGLE_180 = 2
	ANGLE_270 = 3
)

type Tetromino struct {
	Type TetrominoType

	coordinate Coordinate
	angle      Angle
}

func (t Tetromino) Rotate(rotation Rotation) Tetromino {
	switch t.angle {
	case ANGLE_0:
		if ROTATION_CLOCK == rotation {
			t.angle = ANGLE_90
		} else if ROTATION_ANTICLOCK == rotation {
			t.angle = ANGLE_270
		}
	case ANGLE_90:
		if ROTATION_CLOCK == rotation {
			t.angle = ANGLE_180
		} else if ROTATION_ANTICLOCK == rotation {
			t.angle = ANGLE_0
		}
	case ANGLE_180:
		if ROTATION_CLOCK == rotation {
			t.angle = ANGLE_270
		} else if ROTATION_ANTICLOCK == rotation {
			t.angle = ANGLE_90
		}
	case ANGLE_270:
		if ROTATION_CLOCK == rotation {
			t.angle = ANGLE_0
		} else if ROTATION_ANTICLOCK == rotation {
			t.angle = ANGLE_180
		}
	}
	return t
}

func (t Tetromino) Move(direction Direction) Tetromino {
	if DIRECTION_LEFT == direction {
		t.coordinate.X -= 1
	} else if DIRECTION_RIGHT == direction {
		t.coordinate.X += 1
	} else if DIRECTION_DOWN == direction {
		t.coordinate.Y += 1
	}
	return t
}

func (t Tetromino) Coordinates() []Coordinate {
	var absolutes []Coordinate
	switch t.Type {
	case TETROMINO_O:
		absolutes = []Coordinate{
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 0},
			{X: 1, Y: 1},
		}
	case TETROMINO_I:
		switch t.angle {
		case ANGLE_0, ANGLE_180:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
			}
		case ANGLE_90, ANGLE_270:
			absolutes = []Coordinate{
				{X: 0, Y: -1},
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
			}
		}
	case TETROMINO_T:
		switch t.angle {
		case ANGLE_0:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 0, Y: 1},
			}
		case ANGLE_90:
			absolutes = []Coordinate{
				{X: 0, Y: -1},
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 0, Y: 1},
			}
		case ANGLE_180:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 0, Y: -1},
			}
		case ANGLE_270:
			absolutes = []Coordinate{
				{X: 0, Y: -1},
				{X: 1, Y: 0},
				{X: 0, Y: 0},
				{X: 0, Y: 1},
			}
		}
	case TETROMINO_L:
		switch t.angle {
		case ANGLE_0:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: -1, Y: 1},
			}
		case ANGLE_90:
			absolutes = []Coordinate{
				{X: -1, Y: -1},
				{X: 0, Y: -1},
				{X: 0, Y: 0},
				{X: 0, Y: 1},
			}
		case ANGLE_180:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: -1},
			}
		case ANGLE_270:
			absolutes = []Coordinate{
				{X: 0, Y: -1},
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 1, Y: 1},
			}
		}
	case TETROMINO_J:
		switch t.angle {
		case ANGLE_0:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			}
		case ANGLE_90:
			absolutes = []Coordinate{
				{X: 0, Y: -1},
				{X: 0, Y: -1},
				{X: 0, Y: 1},
				{X: -1, Y: 1},
			}
		case ANGLE_180:
			absolutes = []Coordinate{
				{X: -1, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: -1, Y: -1},
			}
		case ANGLE_270:
			absolutes = []Coordinate{
				{X: 0, Y: -1},
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 1, Y: -1},
			}
		}
	case TETROMINO_Z:
		switch t.angle {
		case ANGLE_0, ANGLE_180:
			absolutes = []Coordinate{
				{X: 0, Y: 1},
				{X: 0, Y: 0},
				{X: -1, Y: 0},
				{X: 1, Y: 1},
			}
		case ANGLE_90, ANGLE_270:
			absolutes = []Coordinate{
				{X: 0, Y: 1},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: -1},
			}
		}
	case TETROMINO_S:
		switch t.angle {
		case ANGLE_0, ANGLE_180:
			absolutes = []Coordinate{
				{X: -1, Y: 1},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: -1},
			}
		case ANGLE_90, ANGLE_270:
			absolutes = []Coordinate{
				{X: 0, Y: 0},
				{X: 0, Y: -1},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			}
		}
	}

	return []Coordinate{
		{X: t.coordinate.X + absolutes[0].X, Y: t.coordinate.Y + absolutes[0].Y},
		{X: t.coordinate.X + absolutes[1].X, Y: t.coordinate.Y + absolutes[1].Y},
		{X: t.coordinate.X + absolutes[2].X, Y: t.coordinate.Y + absolutes[2].Y},
		{X: t.coordinate.X + absolutes[3].X, Y: t.coordinate.Y + absolutes[3].Y},
	}
}
