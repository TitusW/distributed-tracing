package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TitusW/notifications/internal/entity"
)

const (
	MaxRetries int = 3
)

type Job struct {
	ID          string
	Type        entity.Channel
	Recipient   string
	ContextData map[string]any
	RetryCount  int
	MaxRetries  int
}

type WorkerPool struct {
	maxWorkers int
	jobQueue   chan Job
	wg         sync.WaitGroup
	repo       NotificationDispatchRepoItf
}

func NewWorkerPool(maxWorkers int, queueSize int, repo NotificationDispatchRepoItf) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		jobQueue:   make(chan Job, queueSize),
		repo:       repo,
	}
}

func (p *WorkerPool) Start(ctx context.Context) {
	for i := 1; i <= p.maxWorkers; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}
}

func (p *WorkerPool) Submit(job Job) {
	p.jobQueue <- job
}

func (p *WorkerPool) Stop() {
	close(p.jobQueue)
	p.wg.Wait()
}

func (p *WorkerPool) worker(ctx context.Context, id int) {
	defer p.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-p.jobQueue:
			if !ok {
				return
			}

			// Process the notification with timeout safety
			if err := p.processNotification(ctx, id, job); err != nil {
				return
			}
		}
	}
}

func (p *WorkerPool) processNotification(ctx context.Context, workerID int, job Job) error {
	// Create a hard execution deadline per notification (e.g., 5 seconds)
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var err error
	switch job.Type {
	case entity.InApp:
		// Simulate external network latency (API call to Twilio/SendGrid)
		time.Sleep(1 * time.Second)
		fmt.Println("In App")
	case entity.Email:
		// Simulate external network latency (API call to Twilio/SendGrid)
		time.Sleep(1 * time.Second)
		fmt.Println("Email")
	case entity.Sms:
		// Simulate external network latency (API call to Twilio/SendGrid)
		time.Sleep(1 * time.Second)
		fmt.Println("SMS")
	case entity.Push:
		// Simulate external network latency (API call to Twilio/SendGrid)
		time.Sleep(1 * time.Second)
		fmt.Println("Push")
	}

	p.finalizeJobStatus(ctx, job, err)

	return nil
}

func (p *WorkerPool) finalizeJobStatus(ctx context.Context, job Job, err error) {
	if err != nil {
		p.repo.UpdateNotificationDispatch(ctx, entity.UpdateNotificationDispatch{
			Ksuid:  job.ID,
			Status: entity.StatusSent,
			ProviderResponse: map[string]any{
				"data": "success",
			},
			SentAt: time.Now().UTC(),
		})
		return
	}

	nextRetryCount := job.RetryCount + 1

	if nextRetryCount >= job.MaxRetries {
		p.repo.UpdateNotificationDispatch(ctx, entity.UpdateNotificationDispatch{
			Ksuid:  job.ID,
			Status: entity.StatusDead,
		})
	} else {
		p.repo.UpdateNotificationDispatch(ctx, entity.UpdateNotificationDispatch{
			Ksuid:  job.ID,
			Status: entity.StatusFailed,
		})
	}
}
