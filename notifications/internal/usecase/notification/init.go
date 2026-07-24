package notification

import (
	"context"

	"github.com/TitusW/notifications/internal/entity"
	unitofwork "github.com/TitusW/unit-of-work"
)

type NotificationRepoItf interface {
	CreateNotification(ctx context.Context, input entity.CreateNotification) (entity.Notification, error)
}

type NotificationDispatchRepoItf interface {
	BulkCreateNotificationDispatches(ctx context.Context, input []entity.CreateNotificationDispatch) error
	UpdateNotificationDispatch(ctx context.Context, input entity.UpdateNotificationDispatch) error
}

type Usecase struct {
	NotificationRepo         NotificationRepoItf
	NotificationDispatchRepo NotificationDispatchRepoItf
	Uow                      unitofwork.UnitOfWorkItf
}

func New(notificationRepo NotificationRepoItf, notificationDispatchRepo NotificationDispatchRepoItf, uow unitofwork.UnitOfWorkItf) Usecase {
	return Usecase{
		NotificationRepo:         notificationRepo,
		NotificationDispatchRepo: notificationDispatchRepo,
		Uow:                      uow,
	}
}
