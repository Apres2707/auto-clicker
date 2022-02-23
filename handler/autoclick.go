package handler

import (
	"context"
	"errors"
)

type AutoClick struct {
	Storage Storage
	Breaker Breaker
	Clicker Clicker
}

func NewAutoClick(
	Storage Storage,
	Breaker Breaker,
	Clicker Clicker,
) (AutoClick, error) {
	if Storage == nil {
		return AutoClick{}, errors.New("storage cannot be nil")
	}
	if Breaker == nil {
		return AutoClick{}, errors.New("breaker cannot be nil")
	}
	if Clicker == nil {
		return AutoClick{}, errors.New("clicker cannot be nil")
	}

	return AutoClick{
		Storage: Storage,
		Breaker: Breaker,
		Clicker: Clicker,
	}, nil
}

func (a AutoClick) Handle() error {
	ctxWithCancel, cancelFunction := context.WithCancel(context.Background())

	go func() {
		err := a.Breaker.Check()
		if err != nil {
			// TODO log
		}
		cancelFunction()
	}()

	config, err := a.Storage.GetConfig()
	if err != nil {
		return err
	}
	a.Clicker.Start(ctxWithCancel, config.Actions)

	return nil
}
