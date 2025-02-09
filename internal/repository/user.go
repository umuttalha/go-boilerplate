package repository

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" means this field won't be included in JSON
	Status   string `json:"status"`
}

func (r *Repository) CreateUser(user *User) error {
	query := `
		INSERT INTO users (name, email, password, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Status).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, password, status FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
