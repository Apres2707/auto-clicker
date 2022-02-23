package service

import (
	"auto-clicker/model"
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"time"
)

const maxConfigureIteration = 10

type Action struct{}

func (c Action) FillActions(ctx context.Context) []model.Action {
	i := 1
	var actionList []model.Action
	for i <= maxConfigureIteration {
		time.Sleep(5 * time.Second)
		select {
		case <-ctx.Done():
			return actionList
		default:
			actionList = append(actionList, addAction(i))
		}
		i++
	}

	return actionList
}

func addAction(i int) model.Action {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	return model.Action{
		Name:             strconv.Itoa(i),
		XCoordinate:      x,
		YCoordinate:      y,
		DelayAfter:       0,
		ScrollAfterDelay: 0,
	}
}
