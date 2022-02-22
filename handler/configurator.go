package handler

import (
	"auto-clicker/model"
	"errors"
	"fmt"
)

type Configurator struct {
	Storage Storage
	ActionConfigurator Action
}

func NewConfigurator(
	Storage Storage,
	ActionConfigurator Action,
) (Configurator, error) {
	if Storage == nil {
		return Configurator{}, errors.New("storage cannot be nil")
	}
	if ActionConfigurator == nil {
		return Configurator{}, errors.New("actionConfigurator cannot be nil")
	}

	return Configurator{
		Storage: Storage,
		ActionConfigurator: ActionConfigurator,
	}, nil
}

func (c Configurator) InitConfig() error {
	actionList := c.ActionConfigurator.FillActions()
	config := model.Config{
		Actions: actionList,
	}

	err := c.Storage.Save(config)
	if err != nil {
		return fmt.Errorf("save config: %w", err)
	}

	return nil
}
