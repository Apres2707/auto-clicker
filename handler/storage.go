package handler

import "auto-clicker/model"

type Storage interface {
	Save(config model.Config) error
}
