package db

import (
	"github.com/jackc/pgx"
)

type Database struct {
	connection *pgx.Conn
	UsersTable *UsersTable
}

func New(cfg pgx.ConnConfig) (*Database, error) {
	conn, err := pgx.Connect(cfg)
	if err != nil {
		return nil, err
	}

	db := &Database{
		connection: conn,
		UsersTable: &UsersTable{conn},
	}
	db.UsersTable.createTable()
	return db, nil
}
