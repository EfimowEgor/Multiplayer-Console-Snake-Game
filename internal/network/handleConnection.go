package network

import (
	"net"
	"snake/internal/components"
	"snake/internal/config"
	"sync"
)

func HandleConnection(conn net.Conn, pool *Pool) {
	defer conn.Close()

	pool.Lock()
	err := pool.AddConnection(conn)
	pool.Unlock()

	if err != nil {
		conn.Write([]byte(err.Error() + "\n"))
		return
	}

	rows := config.GameConfig.ROWS
	cols := config.GameConfig.COLS
	speed := config.GameConfig.SPEED

	// for every new connection init game
	snake := components.InitSnake(rows, cols)
	g := components.CreateEmptyField(rows, cols)
	g.GetSnake(*snake)
	g.Food = g.GenerateFood()
	mat := g.DisplayGrid()

	_, err = conn.Write([]byte(mat))
	if err != nil {
		panic(err)
	}

	STPLSCH := make(chan struct{})
	STPRDCH := make(chan struct{})
	MVCH := make(chan rune)

	var wg sync.WaitGroup
	wg.Add(2)

	go components.GameLoop(g, snake, conn, STPLSCH, STPRDCH, MVCH, &wg, speed)

	go HandleUserInput(conn, STPLSCH, STPRDCH, MVCH, &wg, pool)

	wg.Wait()
}
