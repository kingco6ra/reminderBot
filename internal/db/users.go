package db

import (
	"log"
	"time"

	"github.com/jackc/pgx"
)

type UsersTable struct {
	connection *pgx.Conn
}

type User struct {
	UserID    uint64    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (ut *UsersTable) createTable() {
	// TODO: add tzinfo
	_, err := ut.connection.Exec(
		`CREATE TABLE IF NOT EXISTS users (
            id serial PRIMARY KEY, 
            user_id integer UNIQUE, 
            created_at timestamp DEFAULT current_timestamp
        )`,
	)
	if err != nil {
		log.Fatal("Creating table error: ", err)
	}
}

func (ut *UsersTable) Select(id uint64) (*User, bool) {
	var user User

	err := ut.connection.QueryRow(
		`SELECT user_id, created_at FROM users WHERE user_id = $1`, id,
	).Scan(&user.UserID, &user.CreatedAt)

	if err != nil {
		return nil, false
	}

	return &user, true
}

func (ut *UsersTable) Insert(user User) {
	if _, exists := ut.Select(user.UserID); exists {
		return
	}

	_, err := ut.connection.Exec(
		`INSERT INTO users (user_id) VALUES ($1)`, user.UserID,
	)

	if err != nil {
		log.Fatal("Inserting error: ", err)
	}
}

func (ut *UsersTable) Update() {}
func (ut *UsersTable) Drop()   {}
