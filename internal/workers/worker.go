package workers

import (
	"context"
	"fmt"
	"github.com/giusepperoro/queuepay.git/internal/database"
	"github.com/giusepperoro/queuepay.git/internal/redis"
	"log"
	"time"
)

type Worker interface {
	Run()
}

type Work struct {
	db     database.DbManager
	client redis.Queue
}

func InitWorkers(db database.DbManager, client redis.Queue, amount int64) {
	var i int64
	for i = 0; i < amount; i++ {
		w := NewWorker(db, client)
		w.Run()
	}
}

func NewWorker(db database.DbManager, q redis.Queue) *Work {
	return &Work{
		db:     db,
		client: q,
	}
}

func (w *Work) Run() {
	fmt.Println("worker added...")
	go func() {
		for {
			err := w.process()
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func (w *Work) process() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ok, amount, clientId := w.client.Get()
	if !ok {
		return nil
	}
	_, err := w.db.ChangeBalance(ctx, clientId, amount)
	return err
}
