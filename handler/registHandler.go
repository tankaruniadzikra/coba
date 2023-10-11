package handler

import (
	"database/sql"
	"pair-project/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User, db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO Users (Email, Password, FirstName, LastName) VALUES (?,?,?,?)", user.Email, hashedPassword, user.FirstName, user.LastName)
	return err
}
