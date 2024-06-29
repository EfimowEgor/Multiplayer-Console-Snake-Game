package main

import (
	"log"
	"net"
	"os"
	"snake/internal/config"
	"snake/internal/network"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	ADDR, _ := os.LookupEnv("ADDR")
	PORT, _ := os.LookupEnv("PORT")

	l, err := net.Listen("tcp", ADDR+":"+PORT)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		log.Print(conn.LocalAddr().String())
		go network.HandleConnection(conn)
	}
}
