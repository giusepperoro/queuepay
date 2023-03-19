package handler

import (
	"encoding/json"
	"github.com/giusepperoro/queuepay.git/internal/redis"
	"io"
	"log"
	"net/http"
)

type WithdrawalRequest struct {
	Id     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

type WithdrawalResponse struct {
	Status string `json:"status"`
}

func Withdrawal(q *redis.RedisQueue) func(w http.ResponseWriter, r *http.Request) {
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
		err = q.Set(req.Id, req.Amount)
		response := WithdrawalResponse{
			Status: "transaction in queue",
		}
		if err != nil {
			response.Status = "error adding transaction in queue"
		}
		rawData, err := json.Marshal(response)
		_, err = w.Write(rawData)
		if err != nil {
			log.Printf("unable to write data: %v", err)
		}
	}
}
