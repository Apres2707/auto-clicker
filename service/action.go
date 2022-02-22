package service

import (
	"auto-clicker/model"
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"time"
)

type Action struct {}

func (c Action) FillActions() []model.Action {
	i := 1
	max := 3

	var actionList []model.Action
	for i <= max {
		time.Sleep(5 * time.Second)
		x, y := robotgo.GetMousePos()
		fmt.Println("pos:", x, y)
		actionList = append(actionList, model.Action{
			Name:             strconv.Itoa(i),
			XCoordinate:      x,
			YCoordinate:      y,
			DelayAfter:       0,
			ScrollAfterDelay: 0,
		})
		i++
	}

	return actionList
}
