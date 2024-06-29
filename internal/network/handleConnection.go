package network

import (
	"net"
	"snake/internal/config"
	"snake/internal/components"
	"sync"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	rows := config.GameConfig.ROWS
	cols := config.GameConfig.COLS
	speed := config.GameConfig.SPEED

	// for every new connection init game
	snake := services.InitSnake(rows, cols)
	g := services.CreateEmptyField(rows, cols)
	g.GetSnake(*snake)
	g.Food = g.GenerateFood()
	mat := g.DisplayGrid()

	_, err := conn.Write([]byte(mat))
	if err != nil {
		panic(err)
	}

	STPLSCH := make(chan struct{})
	STPRDCH := make(chan struct{})
	MVCH := make(chan rune)

	var wg sync.WaitGroup
	wg.Add(2)

	go services.GameLoop(g, snake, conn, STPLSCH, STPRDCH, MVCH, &wg, speed)

	go HandleUserInput(conn, STPLSCH, STPRDCH, MVCH, &wg)

	wg.Wait()
}
