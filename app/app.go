package app

import (
	"errors"
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) < 2 {
		panic("command cannot be empty")
	}
	handler, err := getHandler(os.Args[1])
	if err != nil {
		panic(err)
	}
	err = handler.Handle()
	if err != nil {
		panic(err)
	}

	fmt.Println("exiting...")
}

func getHandler(cmd string) (Handler, error) {
	container, err := NewContainer()
	if err != nil {
		panic("container create error")
	}

	switch cmd {
	case "cfg":
		return container.Handler.Configurator, nil
	case "run":
		return container.Handler.AutoClick, nil
	default:
		return nil, errors.New("unknown command")
	}
}
