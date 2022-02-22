package main

import (
	"auto-clicker/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/mattn/go-tty"
	"io/ioutil"
	"log"
	"time"
)

const configPath = "./config.json"

func main() {
	ctxWithCancel, cancelFunction := context.WithCancel(context.Background())
	ttyItem, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = ttyItem.Close() }()

	go func() {
		checkQuit()
		cancelFunction()
	}()
	start(ctxWithCancel)
}

func start(ctx context.Context) {
	actionList := getActionList()

	maxIteration := 30000
	i := 1
	for i <= maxIteration {
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
}

func doAction(actionItem model.Action) {
	robotgo.MoveClick(actionItem.XCoordinate, actionItem.YCoordinate)
	time.Sleep(time.Duration(actionItem.DelayAfter) * time.Millisecond)
	if actionItem.ScrollAfterDelay != 0 {
		robotgo.Scroll(0, actionItem.ScrollAfterDelay)
		time.Sleep(500 * time.Millisecond)
	}
}

func getActionList() []model.Action {
	plan, _ := ioutil.ReadFile(configPath)
	var conf model.Config
	_ = json.Unmarshal(plan, &conf)

	return conf.Actions
}

func checkQuit() {
	ttyItem, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = ttyItem.Close() }()

	for {
		r, err := ttyItem.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if string(r) == "q" {
			fmt.Println("press q")
			return
		}
	}
}
