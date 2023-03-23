package workers

import (
	"fmt"
)

type workerPool struct {
	queue  queueManager
	db     database
	logger logger
}

func NewWorkerPool(queue queueManager, db database, logger logger) *workerPool {
	return &workerPool{
		queue:  queue,
		logger: logger,
		db:     db,
	}
}

func (wp *workerPool) RunNewWorker(clientId int64) error {
	ch, err := wp.queue.ConsumeFromQueue(clientId)
	wp.logger.Info("hehhehhehhe")
	if err != nil {
		return fmt.Errorf("unable to consume from queue %d: %w", clientId, err)
	}
	go func() {
		consumeWorker := NewWorker(wp.logger, wp.db)
		consumeWorker.Consume(ch, clientId)
	}()
	return nil
}
