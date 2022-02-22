package main

import (
	"fmt"
)

func main() {
	container, err := NewContainer()
	if err != nil {
		panic("container create error")
	}

	handler := container.Handler.Configurator
	err = handler.InitConfig()
	if err != nil {
		panic("save config error")
	}

	fmt.Println("exiting...")
}
