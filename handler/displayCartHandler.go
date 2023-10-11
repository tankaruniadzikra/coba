package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func DisplayCart(db *sql.DB, user entity.User) error {
	rows, err := db.Query("SELECT products.ProductName, products.Description, ShoppingCart.Quantity FROM ShoppingCart JOIN Products ON products.ProductID = ShoppingCart.ProductID WHERE ShoppingCart.UserID = ?", user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var displayCart entity.DisplayCart
	fmt.Printf("\n[Name] - [Description] - [Quantity]\n")
	for rows.Next() {
		err := rows.Scan(&displayCart.Name, &displayCart.Description, &displayCart.Quantity)
		if err != nil {
			return err
		}

		fmt.Printf("[%v] - [%v] - [%v]\n", displayCart.Name, displayCart.Description, displayCart.Quantity)
	}
	fmt.Println()
	return nil
}
