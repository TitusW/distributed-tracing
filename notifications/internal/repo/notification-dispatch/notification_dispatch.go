package notificationdispatch

import (
	"context"
	"time"

	"github.com/TitusW/notifications/internal/entity"
	unitofwork "github.com/TitusW/unit-of-work"
	"github.com/segmentio/ksuid"
)

func convertCreateToModel(input entity.CreateNotificationDispatch) NotificationDispatch {
	return NotificationDispatch{
		Ksuid:             ksuid.New().String(),
		NotificationKsuid: input.NotificationKsuid,
		Channel:           input.Channel,
		RecipientTarget:   input.RecipientTarget,
		Status:            entity.StatusPending,
		RetryCount:        0,
		InsertedAt:        time.Now().UTC(),
	}
}

func convertCreatesToModels(input []entity.CreateNotificationDispatch) []NotificationDispatch {
	var notificationDispatches []NotificationDispatch

	for i := 0; i < len(input); i++ {
		notificationDispatch := convertCreateToModel(input[i])
		notificationDispatches = append(notificationDispatches, notificationDispatch)
	}
	return notificationDispatches
}

func convertUpdateToModel(input entity.UpdateNotificationDispatch) NotificationDispatch {
	return NotificationDispatch{
		Ksuid:            input.Ksuid,
		Status:           input.Status,
		ProviderResponse: input.ProviderResponse,
		RetryCount:       input.RetryCount,
		SentAt:           input.SentAt,
	}
}

func convertModelToEntity(model NotificationDispatch) entity.NotificationDispatch {
	return entity.NotificationDispatch{
		Ksuid:             model.Ksuid,
		NotificationKsuid: model.NotificationKsuid,
		Channel:           model.Channel,
		RecipientTarget:   model.RecipientTarget,
		Status:            model.Status,
		ProviderResponse:  model.ProviderResponse,
		RetryCount:        model.RetryCount,
		SentAt:            model.SentAt,
	}
}

func convertModelsToEntities(models []NotificationDispatch) []entity.NotificationDispatch {
	var notificationDispatches []entity.NotificationDispatch

	for i := 0; i < len(models); i++ {
		notificationDispatch := convertModelToEntity(models[i])
		notificationDispatches = append(notificationDispatches, notificationDispatch)
	}

	return notificationDispatches
}

func (m Module) BulkCreateNotificationDispatches(ctx context.Context, input []entity.CreateNotificationDispatch) error {
	tx := unitofwork.GetTX(ctx, m.db)

	notificationDispatchers := convertCreatesToModels(input)

	err := tx.WithContext(ctx).Create(notificationDispatchers).Error

	if err != nil {
		return err
	}

	return nil
}

func (m Module) GetDispatchesForPoller(ctx context.Context) ([]entity.NotificationDispatch, error) {
	tx := unitofwork.GetTX(ctx, m.db)

	var notificationDispatches []NotificationDispatch

	query := tx.Model(notificationDispatches)

	query.Where("status = ?", entity.StatusPending).
		Or(
			tx.Where("status = ?", entity.StatusFailed).Where("retry_count <= ?", 3),
		).Find(notificationDispatches)

	return convertModelsToEntities(notificationDispatches), nil
}

func (m Module) UpdateNotificationDispatch(ctx context.Context, input entity.UpdateNotificationDispatch) error {
	tx := unitofwork.GetTX(ctx, m.db)

	updateInput := convertUpdateToModel(input)

	var notificationDispatcher NotificationDispatch

	err := tx.Model(notificationDispatcher).WithContext(ctx).Updates(updateInput).Error

	if err != nil {
		return err
	}

	return nil
}
