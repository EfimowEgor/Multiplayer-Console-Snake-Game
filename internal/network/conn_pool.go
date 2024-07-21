package network

import (
	"errors"
	"fmt"
	"net"
	"snake/pkg/structs"
	"sync"
)

// Pool implements collection to store incoming connetion data (addr, etc)
// It is required to check if connection unique.
type Pool struct {
	LobbyID  string
	
	sync.Mutex
	ConnPool structs.Set[string]
}

func InitConnPool() *Pool {
	data := make([]string, 0)
	return &Pool{
		ConnPool: structs.NewSet(data),
	}
}

func (p *Pool) AddConnection(conn net.Conn) error {
	if p.ConnPool.Find(conn.RemoteAddr().String()) {
		return fmt.Errorf("already connected")
	}
	p.ConnPool.Add(conn.RemoteAddr().String())
	return nil
}

func (p *Pool) DeleteConnection(conn net.Conn) error {
	if !p.ConnPool.Find(conn.RemoteAddr().String()) {
		return errors.New("connection not in the pool")
	}
	p.ConnPool.Remove(conn.RemoteAddr().String())
	return nil
}

func (p *Pool) String() string {
	return fmt.Sprint(p.ConnPool)
}
