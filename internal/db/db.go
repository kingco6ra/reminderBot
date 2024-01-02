package db

import (
	"github.com/jackc/pgx"
)

type Database struct {
	connection *pgx.Conn
	usersTable *UsersTable
}

func New(cfg pgx.ConnConfig) (*Database, error) {
	conn, err := pgx.Connect(cfg)
	if err != nil {
		return nil, err
	}

	db := &Database{
		connection: conn,
		usersTable: &UsersTable{conn},
	}
	db.usersTable.createTable()
	return db, nil
}
