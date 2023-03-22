package main

import (
	"context"
	"fmt"
	"github.com/giusepperoro/queuepay.git/internal/database"
	"github.com/giusepperoro/queuepay.git/internal/rabbit"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/giusepperoro/queuepay.git/internal/config"
	"github.com/giusepperoro/queuepay.git/internal/handler"
)

const configFileNameEnv = "CONFIG_FILE_NAME"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Fatalf("unable to sync zap logger: %v", err)
		}
	}()

	cfg, err := config.GetConfigFromFile("config.yaml")
	if err != nil {
		logger.Fatal("unable to get config file name from env", zap.Error(err))
	}
	fmt.Println("cfg:", cfg)
	db, err := database.New(ctx, cfg)
	if err != nil {
		log.Println(err)
		log.Fatal("database connect error")
	}
	rbt, err := rabbit.ConnectRabbit(cfg)
	if err != nil {
		log.Println(err)
		log.Fatal("rabbitMQ connect error")
	}

	changeBalanceHandle := handler.NewChangeBalanceHandler(logger, db, rbt)

	http.HandleFunc("/charge", changeBalanceHandle.HandleBalanceChange)
	err = http.ListenAndServe(cfg.ServerAddressUrl, nil)
	if err != nil {
		log.Fatal("Server shutdown", err)
	}
}
