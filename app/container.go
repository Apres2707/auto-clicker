package app

import (
	"auto-clicker/handler"
	"auto-clicker/service"
	"auto-clicker/storage"
)

type Container struct {
	Storage struct {
		File storage.File
	}
	Service struct {
		Configurator service.Action
		Breaker      service.Breaker
		Clicker      service.Clicker
	}
	Handler struct {
		Configurator handler.Configurator
		AutoClick    handler.AutoClick
	}
}

func NewContainer() (*Container, error) {
	container := Container{}
	container.registerStorage()
	container.registerService()
	err := container.registerHandler()
	if err != nil {
		return nil, err
	}

	return &container, nil
}

func (c *Container) registerStorage() {
	c.Storage.File = storage.File{}
}

func (c *Container) registerService() {
	c.Service.Configurator = service.Action{}
	c.Service.Breaker = service.Breaker{}
	c.Service.Clicker = service.Clicker{}
}

func (c *Container) registerHandler() error {
	configurator, err := handler.NewConfigurator(c.Storage.File, c.Service.Configurator, c.Service.Breaker)
	if err != nil {
		return err
	}
	autoClick, err := handler.NewAutoClick(c.Storage.File, c.Service.Breaker, c.Service.Clicker)
	if err != nil {
		return err
	}

	c.Handler.Configurator = configurator
	c.Handler.AutoClick = autoClick

	return nil
}
