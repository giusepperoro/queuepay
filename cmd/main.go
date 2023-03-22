package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/giusepperoro/queuepay.git/internal/config"
	"github.com/giusepperoro/queuepay.git/internal/handler"
)

const configFileNameEnv = "CONFIG_FILE_NAME"

func main() {
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	//db, err := database.New(ctx)
	//if err != nil {
	//	log.Println(err)
	//	log.Fatal("database connect error")
	//}
	//_ = db

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

	var handlerWithdrawal handler.HandleWithdrawal
	http.HandleFunc("/charge", handlerWithdrawal.Withdrawal())
	err = http.ListenAndServe(cfg.ServerAddressUrl, nil)
	if err != nil {
		log.Fatal("Server shutdown", err)
	}
}
