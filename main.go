package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
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
	status string `json:"status"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		log.Println(err)
		log.Fatal("redis connect error")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	db, err := New(ctx)
	if err != nil {
		log.Println(err)
		log.Fatal("database connect error")
	}
	http.HandleFunc("/charge", Withdrawal(db, client))
	http.ListenAndServe("0.0.0.0:80", nil)
}

func Withdrawal(db *dataBase, client *redis.Client) func(w http.ResponseWriter, r *http.Request) {
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
		client.LPush(string(req.id), req.amount)

		response := WithdrawalResponse{
			status: "transaction in que",
		}
		rawData, err := json.Marshal(response)
		_, err = w.Write(rawData)
		if err != nil {
			log.Printf("unable to write data: %v", err)
		}
	}
}
