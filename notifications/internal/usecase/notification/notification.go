package notification

import (
	"context"

	"github.com/TitusW/notifications/internal/entity"
)

func (uc Usecase) CreateNotification(ctx context.Context, input entity.CreateNotification) error {
	err := uc.Uow.Do(ctx, func(txCtx context.Context) error {
		notification, err := uc.NotificationRepo.CreateNotification(txCtx, input)

		if err != nil {
			return err
		}

		var dispatches []entity.CreateNotificationDispatch

		for i := 0; i < len(input.Channels); i++ {
			dispatch := entity.CreateNotificationDispatch{
				NotificationKsuid: notification.Ksuid,
				Channel:           input.Channels[i],
				RecipientTarget:   "test",
				Status:            entity.StatusPending,
			}

			dispatches = append(dispatches, dispatch)
		}

		err = uc.NotificationDispatchRepo.BulkCreateNotificationDispatches(txCtx, dispatches)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
