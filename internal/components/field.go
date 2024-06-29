package services

import (
	"fmt"
	"math/rand"
	"snake/internal/config"
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
	newGrid := Grid{
		Mat: make([][]*Cell, rows),
	}
	for i := range newGrid.Mat {
		newGrid.Mat[i] = make([]*Cell, cols)
	}
	// Fill grid
	for i := range rows {
		for j := range cols {
			newGrid.Mat[i][j] = &Cell{
				Value:   config.FieldSymbol,
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
	for i := 0; i < len(g.Mat); i++ {
		for j := 0; j < len(g.Mat[i]); j++ {
			for _, snakeCell := range s.Body {
				if i == snakeCell.X && j == snakeCell.Y {
					g.Mat[i][j] = snakeCell
					break
				} else {
					g.Mat[i][j] = &Cell{
						Value:   config.FieldSymbol,
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
	g.Mat[g.Food.X][g.Food.Y].Value = config.FoodSymbol
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
	g.Mat[emptyCells[rndPos].X][emptyCells[rndPos].Y].Value = config.FoodSymbol

	return g.Mat[emptyCells[rndPos].X][emptyCells[rndPos].Y]
}

func (g *Grid) DisplayGrid() string {
	var mat string
	var spaceBetween = config.GameConfig.FieldSpace
	mat += config.ReturnClearScreen
	for i := range g.Mat {
		for j := range g.Mat[i] {
			switch {
			case g.Mat[i][j].Value == config.FoodSymbol:
				mat += fmt.Sprintf("%s%*c", config.FoodColor, spaceBetween, g.Mat[i][j].Value)
			case g.Mat[i][j].Value == config.BodySumbol || strings.ContainsRune(HEADS, rune(g.Mat[i][j].Value)):
				mat += fmt.Sprintf("%s%*c", config.SnakeColor, spaceBetween, g.Mat[i][j].Value)
			case g.Mat[i][j].Value == config.FieldSymbol:
				mat += fmt.Sprintf("%s%*c", config.GridColor, spaceBetween, g.Mat[i][j].Value)
			}
		}
		mat += config.CRLF
	}

	return mat
}
