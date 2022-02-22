package handler

import "auto-clicker/model"

type Action interface {
	FillActions() []model.Action
}
