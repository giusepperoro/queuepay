package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (h *changeBalanceHandler) HandleBalanceChange(w http.ResponseWriter, r *http.Request) {
	var req changeBalanceRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error reading request body", zap.Error(err))
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error("error unmarshalling data", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// тут будет поход в RabbitMQ
	response := changeBalanceResponse{
		Status: "transaction in queue",
	}
	if err != nil {
		response.Status = "error adding transaction in queue"
	}

	h.addCLientInfo(req.Id)
	rawData, err := json.Marshal(response)
	_, err = w.Write(rawData)
	if err != nil {
		h.logger.Error("error unmarshalling data", zap.Error(err), zap.Int64("clientId", req.Id))
	}
}

func (h *changeBalanceHandler) addCLientInfo(clientId int64) {
	h.clientsMap.Store(clientId, struct{}{})
}