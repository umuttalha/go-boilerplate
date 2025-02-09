package repository

import (
	"database/sql"
)

type Session struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

func (r *Repository) CreateSession(session *Session) error {
	query := `
        INSERT INTO sessions (token, user_id)
        VALUES ($1, $2)`

	_, err := r.db.Exec(query, session.Token, session.UserID)
	return err
}

func (r *Repository) GetSession(token string) (*Session, error) {
	session := &Session{}
	query := `
        SELECT token, user_id 
        FROM sessions 
        WHERE token = $1`

	err := r.db.QueryRow(query, token).Scan(&session.Token, &session.UserID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *Repository) DeleteSession(token string) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE token = $1", token)
	return err
}

func (r *Repository) DeleteUserSessions(userID int) error {
	query := `DELETE FROM sessions WHERE user_id = $1`
	_, err := r.db.Exec(query, userID)
	return err
}
