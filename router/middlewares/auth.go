package middlewares

import (
	"github.com/alexedwards/argon2id"
	"errors"
)

func AuthenticateUser(email string, password string) error {
	storedPassword := ""

	if err := Database.Get(&storedPassword, "SELECT password FROM users WHERE email=?", email); err != nil{
		return err
	}

	if match, err := argon2id.ComparePasswordAndHash(password, storedPassword); err != nil || !match {
		if !match{
			return errors.New("Passwords do not match.") 
		}
		return err
	}

	return nil
}

func changePassword(email string, password string) error {
	hashedPassword, _ := argon2id.CreateHash(password, argon2id.DefaultParams)
	if _, err := Database.Exec("UPDATE users SET password=? WHERE email=?", hashedPassword); err != nil {
		return err
	}

	return nil
}
