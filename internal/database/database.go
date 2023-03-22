package database

import (
	"context"
	"fmt"
	"github.com/giusepperoro/queuepay.git/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New(ctx context.Context, cfg config.ServiceConfiguration) (*DataBase, error) {
	connection, err := pgxpool.Connect(ctx, cfg.PostgresConnectUrl)
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &DataBase{Conn: connection}, nil
}

func (d *DataBase) ChangeBalance(ctx context.Context, clientId int64, amount int64) (bool, error) {
	opts := pgx.TxOptions{
		IsoLevel: "serializable",
	}
	err := d.Conn.BeginTxFunc(ctx, opts,
		func(tx pgx.Tx) error {
			var balance, id int64
			var query = "SELECT balance, client_id FROM accounts WHERE client_id = $1 FOR UPDATE"

			row := d.Conn.QueryRow(ctx, query, clientId)
			err := row.Scan(&balance, &id)
			if err != nil {
				return fmt.Errorf("error in get client from database: %v", err)
			}
			if id != clientId {
				return fmt.Errorf("client does not exist")
			}
			if balance+amount < 0 {
				return fmt.Errorf("negative balance")
			}

			query = "UPDATE accounts SET balance = balance + $1 WHERE client_id = $2"
			err = d.Conn.QueryRow(ctx, query, amount, clientId).Scan()
			if err != nil && err != pgx.ErrNoRows {
				return fmt.Errorf("error update client data: %v", err)
			}
			return nil
		},
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
