package handler

import "auto-clicker/model"

type Storage interface {
	SaveConfig(config model.Config) error
	GetConfig() (model.Config, error)
}
