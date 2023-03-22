package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dataBase struct {
	conn *pgxpool.Pool
}

type DbManager interface {
	ChangeBalance(ctx context.Context, clientId int64, amount int64) (bool, error)
}
