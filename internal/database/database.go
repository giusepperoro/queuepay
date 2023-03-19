package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dataBase struct {
	conn *pgxpool.Pool
}

type DbManager interface {
	ChangeBalance(ctx context.Context, clientId int64, amount int64) (bool, error)
}

func New(ctx context.Context) (*dataBase, error) {

	connection, err := pgxpool.Connect(ctx, "postgres://postgres:postgres@database:5432/master")
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &dataBase{conn: connection}, nil
}

func (d *dataBase) ChangeBalance(ctx context.Context, clientId int64, amount int64) (bool, error) {
	var balance, id int64
	var query = "SELECT balance, client_id FROM accounts WHERE client_id = $1"

	row := d.conn.QueryRow(ctx, query, clientId)
	err := row.Scan(&balance, &id)
	if err != nil {
		return false, fmt.Errorf("error in get client from database: %v", err)
	}
	if id != clientId {
		return false, errors.New("client does not exist")
	}
	if balance+amount < 0 {
		return false, errors.New("negative balance")
	}

	query = "UPDATE accounts SET balance = balance + $1 WHERE client_id = $2"
	err = d.conn.QueryRow(ctx, query, amount, clientId).Scan()
	if err != nil && err != pgx.ErrNoRows {
		return false, fmt.Errorf("error update client data: %v", err)
	}
	return true, nil
}
