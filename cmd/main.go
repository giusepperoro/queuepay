package main

import (
	"github.com/giusepperoro/queuepay.git/internal/handler"
	"log"
	"net/http"
)

func main() {
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	//db, err := database.New(ctx)
	//if err != nil {
	//	log.Println(err)
	//	log.Fatal("database connect error")
	//}
	//_ = db
	var handlerWithdrawal handler.HandleWithdrawal
	http.HandleFunc("/charge", handlerWithdrawal.Withdrawal())
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("Server shutdown", err)
	}
}
