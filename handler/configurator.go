package handler

import (
	"auto-clicker/model"
	"context"
	"errors"
	"fmt"
)

type Configurator struct {
	Storage            Storage
	ActionConfigurator ActionService
	Breaker            Breaker
}

func NewConfigurator(
	Storage Storage,
	ActionConfigurator ActionService,
	Breaker Breaker,
) (Configurator, error) {
	if Storage == nil {
		return Configurator{}, errors.New("storage cannot be nil")
	}
	if ActionConfigurator == nil {
		return Configurator{}, errors.New("actionConfigurator cannot be nil")
	}
	if Breaker == nil {
		return Configurator{}, errors.New("breaker cannot be nil")
	}

	return Configurator{
		Storage:            Storage,
		ActionConfigurator: ActionConfigurator,
		Breaker:            Breaker,
	}, nil
}

func (c Configurator) Handle() error {
	ctxWithCancel, cancelFunction := context.WithCancel(context.Background())

	go func() {
		err := c.Breaker.Check()
		if err != nil {
			// TODO log
		}
		cancelFunction()
	}()

	actionList := c.ActionConfigurator.FillActions(ctxWithCancel)
	config := model.Config{
		Actions: actionList,
	}

	err := c.Storage.SaveConfig(config)
	if err != nil {
		return fmt.Errorf("save config: %w", err)
	}

	return nil
}
