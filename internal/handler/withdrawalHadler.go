package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (h *HandleWithdrawal) Withdrawal() func(w http.ResponseWriter, r *http.Request) {
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
		response := WithdrawalResponse{
			Status: "transaction in queue",
		}
		s := h.ForWorkers[req.Id]
		log.Println(s)
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
