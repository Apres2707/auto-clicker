package main

import (
	"auto-clicker/handler"
	"auto-clicker/service"
	"auto-clicker/storage"
)

type Container struct {
	Storage struct{
		File storage.File
	}
	Service struct{
		Configurator service.Action
	}
	Handler struct{
		Configurator handler.Configurator
	}
}

func NewContainer() (*Container, error) {
	var container *Container
	container.registerStorage()
	container.registerService()
	err := container.registerHandler()
	if err != nil {
		return nil, err
	}

	return container, nil
}

func (c *Container) registerStorage()  {
	c.Storage.File = storage.File{}
}

func (c *Container) registerService()  {
	c.Service.Configurator = service.Action{}
}

func (c *Container) registerHandler() error {
	configurator, err := handler.NewConfigurator(c.Storage.File, c.Service.Configurator)
	if err != nil {
		return err
	}

	c.Handler.Configurator = configurator

	return nil
}
