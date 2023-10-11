package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func AddCart(db *sql.DB, user entity.User, ProductID, Quantity int) error {
	_, err := db.Exec("INSERT INTO ShoppingCart (UserID, ProductID, Quantity) VALUES (?,?,?)", user.Id, ProductID, Quantity)
	fmt.Printf("Product ID %d dengan Quantity %d berhasil dimasukkan ke dalam cart\n", ProductID, Quantity)
	return err
}
