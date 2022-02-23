package service

import (
	"auto-clicker/model"
	"context"
	"github.com/go-vgo/robotgo"
	"time"
)

const maxClickIteration = 30000

type Clicker struct{}

func (c Clicker) Start(ctx context.Context, actionList []model.Action) {
	i := 1
	for i <= maxClickIteration {
		for _, actionItem := range actionList {
			select {
			case <-ctx.Done():
				return
			default:
				doAction(actionItem)
			}
		}
		i++
	}

	return
}

func doAction(actionItem model.Action) {
	robotgo.MoveClick(actionItem.XCoordinate, actionItem.YCoordinate)
	time.Sleep(time.Duration(actionItem.DelayAfter) * time.Millisecond)
	if actionItem.ScrollAfterDelay != 0 {
		robotgo.Scroll(0, actionItem.ScrollAfterDelay)
		time.Sleep(500 * time.Millisecond)
	}
}
