package network

import (
	"errors"
	"io"
	"log"
	"net"
	"snake/internal/components"
	"snake/internal/config"
	"sync"
)

func ParseConn(r io.Reader) (string, error) {
	/*
		1. empty conn -> return err
		2. conn contains command to create -> create lobby
		3. conn contains command to connect -> connect to lobby by ID
		Let the input messages be in the format:
		?5#<id> - command to create lobby
		:<id> - command to connect to lobby, if empty fuc u, do it again
		!<name> - command to set name, if empty set random name ????
	*/
	var buf []byte = make([]byte, 32)
	_, err := r.Read(buf)
	log.Printf("%s", buf)

	if err != nil {
		return "", err
	}
	switch {
	case len(buf) == 0:
		return "", errors.New("empty command")
	}
	return string(buf), nil
}

// Нужно разделять handleConnection и handleLobbyConnection
func HandleConnection(conn net.Conn, pool *Pool) {
	defer conn.Close()

	ParseConn(conn)

	pool.Lock()
	err := pool.AddConnection(conn)
	pool.Unlock()

	if err != nil {
		conn.Write([]byte(err.Error() + "\n"))
		return
	}


	// INIT GAME OBJECTS
	rows := config.GameConfig.ROWS
	cols := config.GameConfig.COLS
	speed := config.GameConfig.SPEED
	length := config.GameConfig.LEN

	// for every new connection init game
	snake := components.InitSnake(rows, cols, length)
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
