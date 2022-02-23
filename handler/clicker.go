package handler

import (
	"auto-clicker/model"
	"context"
)

type Clicker interface {
	Start(ctx context.Context, actionList []model.Action)
}
