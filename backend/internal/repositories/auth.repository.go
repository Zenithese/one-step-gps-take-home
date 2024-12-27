package repositories

import (
	"database/sql"
	"errors"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (repo *AuthRepository) GetUserByUsername(username string) (string, string, error) {
	var userID, passwordHash string
	err := repo.DB.QueryRow("SELECT user_id, password_hash FROM users WHERE username = ?", username).Scan(&userID, &passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", nil
		}
		return "", "", err
	}
	return userID, passwordHash, nil
}

func (repo *AuthRepository) CreateUser(username, passwordHash string) error {
	_, err := repo.DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, passwordHash)
	return err
}