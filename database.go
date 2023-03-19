package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type dataBase struct {
	conn *pgxpool.Pool
}

type DbManager interface {
	ChangeBalance(ctx context.Context, clientId int64, amount int64) bool
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

func (d *dataBase) ChangeBalance(ctx context.Context, clientId int64, amount int64) bool {
	var balance, id int64
	var query = "SELECT balance, client_id FROM accounts WHERE client_id = $1"
	row := d.conn.QueryRow(ctx, query, clientId)
	err := row.Scan(&balance, &id)
	if err != nil {
		log.Println(err)
		return false
	}
	if id != clientId {
		log.Println("client does not exist")
		return false
	}
	if balance+amount < 0 {
		log.Println("negative balance")
		return false
	}

	query = "UPDATE accounts SET balance = balance + $1 WHERE client_id = $2"
	d.conn.QueryRow(ctx, query, amount, clientId)
	if err != nil {
		log.Println(err)
		return false
	}

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
