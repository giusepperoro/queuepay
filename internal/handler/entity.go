package handler

import (
	"github.com/giusepperoro/queuepay.git/internal/database"
	"github.com/giusepperoro/queuepay.git/internal/rabbit"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rabbitmq/amqp091-go"
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
	db         *pgxpool.Pool
	rbt        *amqp091.Connection
	logger     logger
}

func NewChangeBalanceHandler(logger logger, db *database.DataBase, rbt *rabbit.QueueRabbit) *changeBalanceHandler {
	return &changeBalanceHandler{
		clientsMap: sync.Map{},
		db:         db.Conn,
		rbt:        rbt.Client,
		logger:     logger,
	}
}
