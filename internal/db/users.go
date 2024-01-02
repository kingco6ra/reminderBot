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
	UserID    int64      `db:"user_id"`
	CreatedAt *time.Time `db:"created_at"`
	Language  *string    `db:"language"`
	Timezone  *string    `db:"timezone"`
	Lat       *float64   `db:"lat"`
	Lon       *float64   `db:"lon"`
}

func (ut *UsersTable) createTable() {
	_, err := ut.connection.Exec(
		`CREATE TABLE IF NOT EXISTS users (
            id serial PRIMARY KEY, 
            user_id integer UNIQUE, 
            created_at timestamp DEFAULT current_timestamp,
			language VARCHAR NULL,
			timezone VARCHAR NULL,
			lat NUMERIC NULL,
			lon NUMERIC NULL
        )`,
	)
	if err != nil {
		log.Fatal("Creating table error: ", err)
	}
}

func (ut *UsersTable) Select(id int64) (*User, bool) {
	var user User

	err := ut.connection.QueryRow(
		`
		SELECT
			user_id,
			created_at,
			language,
			timezone,
			lat,
			lon
		FROM users WHERE user_id = $1`, id,
	).Scan(&user.UserID, &user.CreatedAt, &user.Language, &user.Timezone, &user.Lat, &user.Lon)

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
		`INSERT INTO users (user_id, language, timezone, lat, lon) VALUES ($1, $2, $3, $4, $5)`,
		user.UserID, user.Language, user.Timezone, user.Lat, user.Lon,
	)

	if err != nil {
		log.Fatal("Inserting error: ", err)
	}
}

func (ut *UsersTable) Update(user User) {
	existingUser, exists := ut.Select(user.UserID)
	if !exists {
		log.Println("User not found for update")
		return
	}

	if user.CreatedAt != nil {
		existingUser.CreatedAt = user.CreatedAt
	}
	if user.Language != nil {
		existingUser.Language = user.Language
	}
	if user.Timezone != nil {
		existingUser.Timezone = user.Timezone
	}
	if user.Lat != nil {
		existingUser.Lat = user.Lat
	}
	if user.Lon != nil {
		existingUser.Lon = user.Lon
	}

	_, err := ut.connection.Exec(
		`UPDATE users SET created_at = $1, language = $2, timezone = $3, lat = $4, lon = $5 WHERE user_id = $6`,
		existingUser.CreatedAt, existingUser.Language, existingUser.Timezone, existingUser.Lat, existingUser.Lon, existingUser.UserID,
	)

	if err != nil {
		log.Fatal("Updating error: ", err)
	}
}

func (ut *UsersTable) Drop() {}
