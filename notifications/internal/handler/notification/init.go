package notification

import (
	"context"

	"github.com/TitusW/notifications/internal/entity"
)

type UsecaseItf interface {
	CreateNotification(ctx context.Context, input entity.CreateNotification) error
}

type Handler struct {
	uc UsecaseItf
}

func New(uc UsecaseItf) Handler {
	return Handler{
		uc: uc,
	}
}
