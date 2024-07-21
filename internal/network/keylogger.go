package network

import (
	"log"
	"net"
	"snake/internal/config"
	"strings"
	"sync"
)

func HandleUserInput(conn net.Conn, STPLSCH chan struct{}, STPRDCH chan struct{}, MVCH chan rune, wg *sync.WaitGroup, pool *Pool) {
	defer wg.Done()
	for {
		var buf []byte = make([]byte, 1)
		_, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		log.Printf("%s", buf)
		char := rune(buf[0])
		log.Printf("READ: %c\n", buf)
		select {
		case <-STPRDCH:
			pool.Lock()
			err := pool.DeleteConnection(conn)
			pool.Unlock()

			if err != nil {
				log.Printf("Tried to delete connection %s from pool, but it is not in the %s. Conn closed%s",
							conn.RemoteAddr().String(), pool, config.CRLF)
				return
			}
			
			log.Printf("ConnPool after lose: %s", pool)
			return
		default:
			if char == 'q' {
				pool.Lock()
				pool.DeleteConnection(conn)
				pool.Unlock()

				log.Printf("ConnPool after manually closed: %s", pool)

				close(STPLSCH)
				close(MVCH)

				conn.Write([]byte(config.ReturnClearScreen + "GAME STOPPED\n"))

				return
			}
		}
		if strings.ContainsRune("wasd", char) {
			MVCH <- char
		}
	}
}
