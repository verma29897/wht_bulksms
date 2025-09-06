package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/verma29897/bulksms/db"
)

type User struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	Username     string    `json:"username" gorm:"uniqueIndex"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

func scanUser(row *sql.Row) (*User, error) {
	var u User
	var created sql.NullTime
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Username, &u.PasswordHash, &created); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if created.Valid {
		u.CreatedAt = created.Time
	}
	return &u, nil
}

func GetUserByEmail(email string) (*User, error) {
	conn := db.GetDB()
	row := conn.QueryRow(`SELECT id, name, email, username, password_hash, created_at FROM users WHERE email = $1`, email)
	return scanUser(row)
}

func GetUserByUsername(username string) (*User, error) {
	conn := db.GetDB()
	row := conn.QueryRow(`SELECT id, name, email, username, password_hash, created_at FROM users WHERE username = $1`, username)
	return scanUser(row)
}

func CreateUser(name, email, username, passwordHash string) (*User, error) {
	conn := db.GetDB()
	row := conn.QueryRow(`INSERT INTO users (name, email, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id, created_at`, name, email, username, passwordHash)
	var id int64
	var created sql.NullTime
	if err := row.Scan(&id, &created); err != nil {
		return nil, err
	}
	var createdAt time.Time
	if created.Valid {
		createdAt = created.Time
	}
	return &User{ID: id, Name: name, Email: email, Username: username, PasswordHash: passwordHash, CreatedAt: createdAt}, nil
}
