package handler

import (
	"auto-clicker/model"
	"context"
)

type ActionService interface {
	FillActions(ctx context.Context) []model.Action
}
