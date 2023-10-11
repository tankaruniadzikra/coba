package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func DeleteCart(db *sql.DB, user entity.User, ProductID int) error {
	_, err := db.Exec("DELETE FROM ShoppingCart WHERE UserID = ? AND ProductID = ?", user.Id, ProductID)
	if err != nil {
		return err
	}

	fmt.Printf("Product ID %d berhasil dihapus dari cart\n", ProductID)
	return nil
}
