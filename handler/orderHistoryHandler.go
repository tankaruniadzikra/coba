package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func DisplayOrderHistory(db *sql.DB, user entity.User) error {
	rows, err := db.Query(`
		SELECT 
			Orders.OrderID,
			Products.ProductName,
			OrderItems.Quantity,
			Orders.OrderDate,
			Orders.TotalAmount
		FROM 
			Orders
		JOIN 
			OrderItems ON Orders.OrderID = OrderItems.OrderID
		JOIN 
			Products ON OrderItems.ProductID = Products.ProductID
		WHERE 
			Orders.UserID = ?
	`, user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var orderHistory entity.OrderHistory
	fmt.Printf("\n[OrderID] - [ProductName] - [Quantity] - [OrderDate] - [TotalAmount]\n")
	for rows.Next() {
		err := rows.Scan(
			&orderHistory.OrderID,
			&orderHistory.ProductName,
			&orderHistory.Quantity,
			&orderHistory.OrderDate,
			&orderHistory.TotalAmount,
		)
		if err != nil {
			return err
		}

		fmt.Printf("[%v] - [%v] - [%v] - [%v] - [%.2f]\n",
			orderHistory.OrderID,
			orderHistory.ProductName,
			orderHistory.Quantity,
			orderHistory.OrderDate,
			orderHistory.TotalAmount,
		)
	}
	fmt.Println()
	return nil
}
