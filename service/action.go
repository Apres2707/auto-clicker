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
			addAction(actionList)
		}
		i++
	}

	return actionList
}

func addAction(actionList []model.Action) {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	actionList = append(actionList, model.Action{
		Name:             strconv.Itoa(len(actionList) + 1),
		XCoordinate:      x,
		YCoordinate:      y,
		DelayAfter:       0,
		ScrollAfterDelay: 0,
	})
}
