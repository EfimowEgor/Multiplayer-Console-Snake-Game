package main

import (
	"errors"
	"fmt"
	"snake/internal"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/eiannone/keyboard"
)

func main() {
	var cols, rows int
	var updateFreq time.Duration
	// PARAMS
	// ----------
	cols = 17
	rows = 17
	updateFreq = 200
	// GAME OBJECTS
	// ----------
	snake := internal.InitSnake(rows, cols)
	g := internal.CreateEmptyField(rows, cols)
	g.Food = g.GenerateFood()
	g.GetSnake(*snake)
	g.DisplayGrid()
	// CONTROL CHANNELS
	// ----------
	STPLSCH := make(chan struct{})
	STPRDCH := make(chan struct{})
	MVCH := make(chan rune)

	var wg sync.WaitGroup
	wg.Add(2)

	// PUT TO A SEPARATE FUNC
	go func() {
		defer wg.Done()
		for {
			select {
			case <-STPLSCH:
				return
			case v := <-MVCH:
				// CHANGE HEAD DIR
				if snake.CheckHeadDir(internal.Dir(unicode.ToUpper(v))) {
					snake.Direction = internal.Dir(unicode.ToUpper(v))
				}
			default:
				// MOVE SNAKE
				err := snake.Move(cols-1, rows-1, &g)
				if errors.Is(err, internal.CustomError{Info: "OUT OF BOUNDS OF GRID", Data: *snake.Body[0]}) ||
					errors.Is(err, internal.CustomError{Info: "HIT THE BODY", Data: *snake.Body[0]}) {
					fmt.Printf("\033[H\033[JYOU'VE LOST\nPRESS ANY KEY TO CONTINUE\n")
					close(STPRDCH)
					return
				}
				g.GetSnake(*snake)
				if g.Food != nil {
					g.GetFood()
				} else {
					g.Food = g.GenerateFood()
				}
				g.DisplayGrid()
				time.Sleep(time.Millisecond * updateFreq)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			char, _, err := keyboard.GetSingleKey()
			select {
			case <-STPRDCH:
				return
			default:
				if err != nil {
					panic(err)
				}
				if char == 'q' {
					close(STPLSCH)
					close(MVCH)
					fmt.Printf("\033[H\033[JGAME STOPPED\n")
					return
				}
			}
			if strings.ContainsRune("wasd", char) {
				MVCH <- char
			}
		}
	}()

	wg.Wait()
}
