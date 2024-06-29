package network

import (
	"errors"
	"fmt"
	"net"
	"snake/pkg/structs"
)

// Pool implements collection to store incoming connetion data (addr, etc)
// It is required to check if connection unique.
type Pool struct {
	ConnPool structs.Set[string]
}

func InitConnPool() *Pool {
	data := make([]string, 0)
	return &Pool{
		ConnPool: structs.NewSet(data),
	}
}

func (p *Pool) AddConnection(conn net.Conn) error {
	if p.ConnPool.Find(conn.LocalAddr().String()) {
		return errors.New("connection already added")
	}
	p.ConnPool.Add(conn.LocalAddr().String())
	return nil
}

func (p *Pool) DeleteConnection(conn net.Conn) {
	p.ConnPool.Remove(conn.LocalAddr().String())
}

func (p *Pool) String() string {
	return fmt.Sprint(p.ConnPool)
}