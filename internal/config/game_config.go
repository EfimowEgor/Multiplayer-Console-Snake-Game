package config

import (
	"log"
	"strconv"
	"time"
)

const (
	// CONTROL
	ReturnClearScreen string = "\033[H\033[J"
	ReturnCarriage    string = "\033[H"
	CRLF              string = "\r\n"
	// GAME ELEMENTS
	BodySumbol  byte = 'O'
	FieldSymbol byte = '#'
	FoodSymbol  byte = '*'
	// COLORS
	FoodColor  string = "\033[0;31m"
	SnakeColor string = "\033[0;32m"
	GridColor  string = "\033[0;37m"
)

type configGame struct {
	ROWS       int
	COLS       int
	LEN        int
	SPEED      time.Duration
	FieldSpace int
}

var GameConfig *configGame

func init() {
	rowCount, err := strconv.Atoi(getEnv("ROWS"))
	if err != nil {
		log.Fatal("cannot parse number of rows")
	}
	colCount, err := strconv.Atoi(getEnv("COLS"))
	if err != nil {
		log.Fatal("cannot parse number of cols")
	}
	snakeLen, err := strconv.Atoi(getEnv("COLS"))
	if err != nil {
		log.Fatal("cannot parse snake length")
	}
	speed, err := strconv.Atoi(getEnv("SPEED"))
	if err != nil {
		log.Fatal("cannot parse snake speed")
	}
	space, err := strconv.Atoi(getEnv("SPACE"))
	if err != nil {
		log.Fatal("cannot parse field space")
	}
	GameConfig = &configGame{
		ROWS:       rowCount,
		COLS:       colCount,
		LEN:        snakeLen,
		SPEED:      time.Duration(speed),
		FieldSpace: space,
	}
}
