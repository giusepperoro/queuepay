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

	err = h.publisher.PutToQueue(r.Context(), req.Id, req.Amount)
	if err != nil {
		h.logger.Error("error publishing", zap.Error(err), zap.Int64("clientId", req.Id))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.makeConsumer(req.Id)
	if err != nil {
		h.logger.Error("unable to create worker", zap.Error(err), zap.Int64("clientId", req.Id))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := changeBalanceResponse{
		Status: "transaction in queue",
	}
	if err != nil {
		response.Status = "error adding transaction in queue"
	}

	rawData, err := json.Marshal(response)
	_, err = w.Write(rawData)
	if err != nil {
		h.logger.Error("error unmarshalling data", zap.Error(err), zap.Int64("clientId", req.Id))
	}
}

func (h *changeBalanceHandler) makeConsumer(clientId int64) error {
	_, ok := h.clientsMap.Load(clientId)
	if ok {
		return nil
	}

	h.clientsMap.Store(clientId, struct{}{})
	return h.workerPool.RunNewWorker(clientId)
}
