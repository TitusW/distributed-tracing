package notification

import (
	"context"
	"time"

	"github.com/TitusW/notifications/internal/entity"
	unitofwork "github.com/TitusW/unit-of-work"
	"github.com/segmentio/ksuid"
)

func convertCreateToModel(input entity.CreateNotification) Notification {
	return Notification{
		Ksuid:            ksuid.New().String(),
		IdempotencyKey:   input.IdempotencyKey,
		UserKsuid:        input.UserKsuid,
		EventType:        input.EventType,
		TransactionKsuid: input.TransactionKsuid,
		ContextData:      input.ContextData,
		InsertedAt:       time.Now().UTC(),
	}
}

func convertEntityFromModel(model Notification) entity.Notification {
	return entity.Notification{
		Ksuid:            model.Ksuid,
		IdempotencyKey:   model.IdempotencyKey,
		UserKsuid:        model.UserKsuid,
		EventType:        model.EventType,
		TransactionKsuid: model.TransactionKsuid,
		ContextData:      model.ContextData,
	}
}

func (m Module) CreateNotification(ctx context.Context, input entity.CreateNotification) (entity.Notification, error) {
	tx := unitofwork.GetTX(ctx, m.db)

	notification := convertCreateToModel(input)

	err := tx.WithContext(ctx).Create(notification).Error
	if err != nil {
		return entity.Notification{}, err
	}

	ent := convertEntityFromModel(notification)

	return ent, nil
}
