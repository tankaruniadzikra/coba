package handler

import (
	"database/sql"
	"pair-project/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string, db *sql.DB) (entity.User, bool, error) {
	var user entity.User
	row := db.QueryRow("SELECT UserID, Email, Password, FirstName, LastName FROM Users WHERE Email = ?", email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName)

	if err == sql.ErrNoRows {
		return user, false, nil // Email tidak ditemukan
	} else if err != nil {
		return user, false, err // Error lainnya
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false, nil // Kata sandi salah
	}

	return user, true, nil // Berhasil login
}
