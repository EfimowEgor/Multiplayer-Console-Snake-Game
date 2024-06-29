package components

import "fmt"

type CustomError struct {
	Info string
	Data Cell
}

func (e CustomError) Error() string {
	return "INFO:" + e.Info +
		" CELL:" +
		fmt.Sprintf("VAL:%c, X:%d, Y:%d, WALKABLE: %t, EATABLE: %t",
			e.Data.Value,
			e.Data.X,
			e.Data.Y,
			e.Data.CanWalk,
			e.Data.CanEat,
		)
}

