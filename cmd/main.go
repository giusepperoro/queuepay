package main

import (
	"context"
	"github.com/giusepperoro/queuepay.git/internal/database"
	"github.com/giusepperoro/queuepay.git/internal/handler"
	"github.com/giusepperoro/queuepay.git/internal/redis"
	"github.com/giusepperoro/queuepay.git/internal/workers"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	client := redis.RedisNew()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	db, err := database.New(ctx)
	if err != nil {
		log.Println(err)
		log.Fatal("database connect error")
	}
	workers.InitWorkers(db, client, 10)

	http.HandleFunc("/charge", handler.Withdrawal(client))
	err = http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("Server shutdown", err)
	}
}
