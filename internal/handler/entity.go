package handler

import (
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/giusepperoro/queuepay.git/internal/database"
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
	db         *pgxpool.Pool
	publisher  queueProducer
	logger     logger
}

func NewChangeBalanceHandler(logger logger, db *database.DataBase, publisher queueProducer) *changeBalanceHandler {
	return &changeBalanceHandler{
		clientsMap: sync.Map{},
		db:         db.Conn,
		publisher:  publisher,
		logger:     logger,
	}
}
