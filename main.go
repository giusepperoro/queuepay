package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

type WithdrawalRequest struct {
	id     int64 `json:"id"`
	amount int64 `json:"balance"`
}

type WithdrawalResponse struct {
	status bool `json:"status"`
	//err    string `json:"error,omitempty"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	db, err := New(ctx)
	if err != nil {
		log.Println(err)
		log.Fatal("database connect error")
	}
	http.HandleFunc("/charge", Withdrawal(db))
	http.ListenAndServe("0.0.0.0:80", nil)
}

func Withdrawal(db *dataBase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var req WithdrawalRequest

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		ok := db.ChangeBalance(ctx, req.id, req.amount)
		response := WithdrawalResponse{
			status: ok,
		}
		rawData, err := json.Marshal(response)
		w.WriteHeader(http.StatusOK)

		_, err = w.Write(rawData)
		if err != nil {
			log.Printf("unable to write data: %v", err)
		}
	}
}
