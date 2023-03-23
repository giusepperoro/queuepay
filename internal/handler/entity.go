package handler

import (
	"sync"
)

type changeBalanceRequest struct {
	Id     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

type changeBalanceResponse struct {
	Status string `json:"status"`
}

type changeBalanceHandler struct {
	clientsMap sync.Map
	publisher  queueProducer
	workerPool workerPool
	logger     logger
}

func NewChangeBalanceHandler(logger logger, publisher queueProducer, workerPool workerPool) *changeBalanceHandler {
	return &changeBalanceHandler{
		clientsMap: sync.Map{},
		publisher:  publisher,
		workerPool: workerPool,
		logger:     logger,
	}
}
