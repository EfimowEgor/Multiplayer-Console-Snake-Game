package services

import (
	"fmt"
	"math/rand"
	"strings"
)

type Cell struct {
	Value   byte
	X       int
	Y       int
	CanWalk bool
	CanEat  bool
}

type Grid struct {
	Mat  [][]*Cell
	Food *Cell
}

func CreateEmptyField(rows, cols int) Grid {
	// Create 2D slice
	newGrid := Grid{
		Mat: make([][]*Cell, rows),
	}
	for i := range newGrid.Mat {
		newGrid.Mat[i] = make([]*Cell, cols)
	}
	// Fill 2D slice with '#'
	for i := range newGrid.Mat {
		for j := range newGrid.Mat[i] {
			newGrid.Mat[i][j] = &Cell{
				Value:   '#',
				X:       i,
				Y:       j,
				CanWalk: true,
				CanEat:  false,
			}
		}
	}
	return newGrid
}

func (g *Grid) GetSnake(s Snake) {
	// NEED TO CHANGE
	for i := 0; i < len(g.Mat); i++ {
		for j := 0; j < len(g.Mat[i]); j++ {
			for _, snakeCell := range s.Body {
				if i == snakeCell.X && j == snakeCell.Y {
					g.Mat[i][j] = snakeCell
					break
				} else {
					g.Mat[i][j] = &Cell{
						Value:   '#',
						X:       i,
						Y:       j,
						CanWalk: true,
						CanEat:  false,
					}
				}
			}
		}
	}
}

func (g *Grid) GetFood() {
	g.Mat[g.Food.X][g.Food.Y].Value = '*'
	g.Mat[g.Food.X][g.Food.Y].CanEat = true
}

func (g *Grid) GenerateFood() *Cell {
	var emptyCells []*Cell = make([]*Cell, 0)
	for i := range g.Mat {
		for j := range g.Mat[i] {
			if g.Mat[i][j].CanWalk {
				emptyCells = append(emptyCells, g.Mat[i][j])
			}
		}
	}
	rndPos := rand.Intn(len(emptyCells))

	g.Mat[emptyCells[rndPos].X][emptyCells[rndPos].Y].CanEat = true
	g.Mat[emptyCells[rndPos].X][emptyCells[rndPos].Y].Value = '*'

	return g.Mat[emptyCells[rndPos].X][emptyCells[rndPos].Y]
}

func (g *Grid) DisplayGrid() string {
	var mat string
	for i := range g.Mat {
		for j := range g.Mat[i] {
			switch {
			case g.Mat[i][j].Value == '*':
				mat += fmt.Sprintf("\033[0;31m %*c", 1, g.Mat[i][j].Value)
			case g.Mat[i][j].Value == 'O' || strings.ContainsRune("^<>v", rune(g.Mat[i][j].Value)):
				mat += fmt.Sprintf("\033[0;32m %*c", 1, g.Mat[i][j].Value)
			case g.Mat[i][j].Value == '#':
				mat += fmt.Sprintf("\033[0;37m %*c", 1, g.Mat[i][j].Value)
			}
		}
		mat += "\r\n"
	}
	mat += "\033[H"
	
	return mat
}
