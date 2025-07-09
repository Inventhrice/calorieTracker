package middlewares

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/alexedwards/argon2id"
)

const numTokenBytes = 64

type ActiveToken struct {
	userID string
	Token  string
}

var activeSessions []ActiveToken // mapping emails, referenced by token

func AuthenticateUser(email string, password string) (string, error) {
	storedPassword := struct {
		UserID   string `db:"id"`
		Password string `db:"password"`
	}{"", ""}

	if err := Database.Get(&storedPassword, "SELECT id, password FROM users WHERE email=?", email); err != nil {
		return "", err
	}

	if match, err := argon2id.ComparePasswordAndHash(password, storedPassword.Password); err != nil || !match {
		if !match {
			return "", errors.New("Passwords do not match.")
		}
		return "", err
	}

	removeActiveSession(storedPassword.UserID)
	genToken, _ := generateToken()

	activeSessions = append(activeSessions, ActiveToken{userID: storedPassword.UserID, Token: genToken})

	return genToken, nil
}

// Bearer-Token-Based Authentication
// https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/
func generateToken() (string, error) {
	bytes := make([]byte, numTokenBytes)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ChangePassword(email string, password string) error {
	hashedPassword, _ := argon2id.CreateHash(password, argon2id.DefaultParams)
	if _, err := Database.Exec("UPDATE users SET password=? WHERE email=?", hashedPassword); err != nil {
		return err
	}

	return nil
}

func removeActiveSession(userID string) {
	if found, index := findActiveSession(userID); found {
		activeSessions = append(activeSessions[:index], activeSessions[index+1:]...)
	}
}

func findActiveSession(userID string) (bool, int) {
	for index, activeToken := range activeSessions {
		if activeToken.userID == userID {
			return true, index
		}
	}
	return false, -1
}

func CheckLoggedIn(token string) (string, error) {
	for _, activeToken := range activeSessions {
		if activeToken.Token == token {
			return activeToken.userID, nil
		}
	}
	return "", errors.New("User was not found in the list of active sessions.")
}
