package network

import (
	"log"
	"sync"
	"net"
	"strings"
)

func HandleUserInput(conn net.Conn, STPLSCH chan struct{}, STPRDCH chan struct{}, MVCH chan rune, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		var buf []byte = make([]byte, 1)
		_, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		char := rune(buf[0])
		log.Printf("READ: %c\n", buf)
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
}