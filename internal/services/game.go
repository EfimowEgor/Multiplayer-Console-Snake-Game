package services

import (
	"errors"
	"net"
	"sync"
	"time"
	"unicode"
)

func GameLoop(g Grid, snake *Snake, conn net.Conn,
	STPLSCH chan struct{}, STPRDCH chan struct{}, MVCH chan rune, wg *sync.WaitGroup, freq time.Duration) {
	defer wg.Done()
	mat := g.DisplayGrid()

	_, err := conn.Write([]byte(mat))
	if err != nil {
		panic(errors.New("cannot pass grid to connection"))
	}

	for {
		select {
		case <-STPLSCH:
			return
		case v := <-MVCH:
			// CHANGE HEAD DIR
			if snake.CheckHeadDir(Dir(unicode.ToUpper(v))) {
				snake.Direction = Dir(unicode.ToUpper(v))
			}
		default:
			// MOVE SNAKE
			err := snake.Move(len(g.Mat)-1, len(g.Mat[0])-1, &g)
			if errors.Is(err, CustomError{Info: "OUT OF BOUNDS OF GRID", Data: *snake.Body[0]}) ||
				errors.Is(err, CustomError{Info: "HIT THE BODY", Data: *snake.Body[0]}) {
				mat = "\033[H\033[JYOU'VE LOST\r\nPRESS ANY KEY TO CONTINUE\n"
				_, err = conn.Write([]byte(mat))
				if err != nil {
					panic(err)
				}
				close(STPRDCH)
				return
			}
			g.GetSnake(*snake)
			if g.Food != nil {
				g.GetFood()
			} else {
				g.Food = g.GenerateFood()
			}
			mat = g.DisplayGrid()
			conn.Write([]byte(mat))
			time.Sleep(time.Millisecond * freq)
		}
	}
}
