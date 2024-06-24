package internal

type Dir byte

const (
	UP    Dir = 'W'
	DOWN  Dir = 'S'
	LEFT  Dir = 'A'
	RIGHT Dir = 'D'
)

var HEAD map[Dir]byte = map[Dir]byte{
	UP:    '^',
	DOWN:  'v',
	LEFT:  '<',
	RIGHT: '>',
}

var MVMAT map[Dir][2]int = map[Dir][2]int{
	'D': {1, 0},
	'W': {0, 1},
	'A': {-1, 0},
	'S': {0, -1},
}

type Snake struct {
	Body      []*Cell // [HEAD_COO, BODY_COOs]
	Direction Dir
}

func InitSnake(matX, matY int) *Snake {
	var length int = 5
	var initX, initY int
	initX = (matX - 1) / 2
	initY = (matY - 1) / 2
	snake := &Snake{
		Body:      make([]*Cell, length),
		Direction: RIGHT,
	}
	for i := 0; i < length; i++ {
		var val byte
		if i == 0 {
			val = HEAD[snake.Direction]
		} else {
			val = 'O'
		}
		snake.Body[i] = &Cell{
			Value:   val,
			X:       initX,
			Y:       initY - i,
			CanWalk: false,
			CanEat:  false,
		}
	}
	return snake
}

func (s *Snake) Move(height, width int, g *Grid) error {
	// MOVE BODY
	for i := len(s.Body) - 1; i > 0; i-- {
		s.Body[i].X = s.Body[i-1].X
		s.Body[i].Y = s.Body[i-1].Y
	}

	newHeadPosY := s.Body[0].Y + MVMAT[s.Direction][0]
	newHeadPosX := s.Body[0].X - MVMAT[s.Direction][1]

	// CHECK OUT OF BOUNDS
	if newHeadPosY < 0 || newHeadPosY > height || newHeadPosX > width || newHeadPosX < 0 {
		return CustomError{
			Info: "OUT OF BOUNDS OF GRID",
			Data: *s.Body[0],
		}
	}

	// CHECK COLLISION WITH BODY
	if !g.Mat[newHeadPosX][newHeadPosY].CanWalk {
		return CustomError{
			Info: "HIT THE BODY",
			Data: *s.Body[0],
		}
	}

	// CHECK IF NEXT CELL - FOOD
	var newTailX, newTailY int =  s.Body[0].X,  s.Body[0].Y

	if g.Mat[newHeadPosX][newHeadPosY].CanEat {
		s.Body = append(s.Body, &Cell{
			Value:   'O',
			X:       newTailX,
			Y:       newTailY,
			CanWalk: false,
			CanEat:  false,
		})
		g.Mat[newHeadPosX][newHeadPosY].CanEat = false
		g.Mat[newHeadPosX][newHeadPosY].Value = '#'
	}

	// MOVE HEAD
	s.Body[0].Y += MVMAT[s.Direction][0]
	s.Body[0].X -= MVMAT[s.Direction][1]

	s.Body[0].Value = HEAD[s.Direction]

	return nil
}

func (s *Snake) CheckHeadDir(dir Dir) bool {
	if (s.Direction == UP && dir == DOWN) ||
		(s.Direction == DOWN && dir == UP) ||
		(s.Direction == LEFT && dir == RIGHT) ||
		(s.Direction == RIGHT && dir == LEFT) {
		return false
	}
	return true
}
