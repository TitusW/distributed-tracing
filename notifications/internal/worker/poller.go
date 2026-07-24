package worker

import (
	"context"
	"time"

	"github.com/TitusW/notifications/internal/entity"
)

type NotificationDispatchRepoItf interface {
	GetDispatchesForPoller(ctx context.Context) ([]entity.NotificationDispatch, error)
	UpdateNotificationDispatch(ctx context.Context, input entity.UpdateNotificationDispatch) error
}

type DBPoller struct {
	repo       NotificationDispatchRepoItf
	workerPool *WorkerPool
}

func NewDBPoller(repo NotificationDispatchRepoItf, workerPool *WorkerPool) *DBPoller {
	return &DBPoller{
		repo:       repo,
		workerPool: workerPool,
	}
}

// StartPolling runs on app startup in cmd/server/main.go
func (p *DBPoller) StartPolling(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// 1. Fetch pending notifications from DB
			dispatches, _ := p.repo.GetDispatchesForPoller(ctx)

			// 2. TRIGGER THE WORKERS with real payloads from the DB
			for _, dispatch := range dispatches {
				p.repo.UpdateNotificationDispatch(ctx,
					entity.UpdateNotificationDispatch{
						Ksuid:  dispatch.Ksuid,
						Status: entity.StatusProcessing,
					},
				)

				job := Job{
					ID:         dispatch.Ksuid,
					Type:       dispatch.Channel,
					RetryCount: dispatch.RetryCount,
					MaxRetries: MaxRetries,
				}

				p.workerPool.Submit(job)
			}
		}
	}
}
