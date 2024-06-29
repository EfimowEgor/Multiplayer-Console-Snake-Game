package main

import (
	"log"
	"net"
	"snake/internal/config"
	"snake/internal/network"
)

func main() {
	ADDR := config.ServerConfig.ADDR
	PORT := config.ServerConfig.PORT
	proto := config.ServerConfig.Proto

	l, err := net.Listen(proto, ADDR+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Print(conn.LocalAddr().String())
		go network.HandleConnection(conn)
	}
}
