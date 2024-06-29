package components

import (
	"errors"
	"net"
	"snake/internal/config"
	"sync"
	"time"
	"unicode"
)

func GameLoop(g Grid, snake *Snake, conn net.Conn,
	STPLSCH chan struct{}, STPRDCH chan struct{}, MVCH chan rune, wg *sync.WaitGroup, freq time.Duration) {
	defer wg.Done()

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
			err := snake.Move(config.GameConfig.ROWS-1, config.GameConfig.COLS-1, &g)
			if errors.Is(err, CustomError{Info: "OUT OF BOUNDS OF GRID", Data: *snake.Body[0]}) ||
				errors.Is(err, CustomError{Info: "HIT THE BODY", Data: *snake.Body[0]}) {
				mat := config.ReturnClearScreen + "YOU'VE LOST" + config.CRLF +"PRESS ANY KEY TO CONTINUE\n"
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
			mat := g.DisplayGrid()
			conn.Write([]byte(mat))
			time.Sleep(time.Millisecond * freq)
		}
	}
}
