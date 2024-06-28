package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"snake/internal"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	snake := internal.InitSnake(17, 17)
	g := internal.CreateEmptyField(17, 17)
	g.GetSnake(*snake)
	g.Food = g.GenerateFood()
	mat := g.DisplayGrid()

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
				err := snake.Move(17-1, 17-1, &g)
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
				mat = g.DisplayGrid()
				conn.Write([]byte(mat))
				time.Sleep(time.Millisecond * 250)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			var buf []byte = make([]byte, 1)
			_, err := conn.Read(buf)
			if err != nil {
				panic(err)
			}
			char := rune(buf[0])
			fmt.Printf("READ: %c\n", buf)
			select {
			case <-STPRDCH:
				return
			default:
				if char == 'q' {
					close(STPLSCH)
					close(MVCH)
					conn.Write([]byte("\033[H\033[JGAME STOPPED\n"))
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

func main() {
	ADDR, _ := os.LookupEnv("ADDR")
	PORT, _ := os.LookupEnv("PORT")
	fmt.Println(ADDR + ":" + PORT)
	l, err := net.Listen("tcp", ADDR+":"+PORT)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
