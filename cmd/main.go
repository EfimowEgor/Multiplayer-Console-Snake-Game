package main

import (
	"fmt"
	"net"
	"os"
	"snake/internal/network"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
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
		go network.HandleConnection(conn)
	}
}
