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
	logger     logger
}

func NewChangeBalanceHandler(logger logger) *changeBalanceHandler {
	return &changeBalanceHandler{
		clientsMap: sync.Map{},
		logger:     logger,
	}
}
