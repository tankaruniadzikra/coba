package handler

import (
	"database/sql"
	"fmt"
)

func ViewOrderReport(db *sql.DB, userID int) error {
	rows, err := db.Query(`
		SELECT 
			u.Email, 
			o.OrderDate, 
			p.ProductName, 
			oi.Quantity, 
			oi.PricePerUnit, 
			(oi.Quantity * oi.PricePerUnit) AS TotalAmount
		FROM 
			Orders o
		JOIN 
			Users u ON o.UserID = u.UserID
		JOIN 
			OrderItems oi ON o.OrderID = oi.OrderID
		JOIN 
			Products p ON oi.ProductID = p.ProductID
		WHERE
			u.UserID = ?
	`, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Order Report:")
	fmt.Println("Email\tOrderDate\tProductName\tQuantity\tPricePerUnit\tTotalAmount")

	for rows.Next() {
		var email, orderDate, productName string
		var quantity int
		var pricePerUnit, totalAmount float64
		err := rows.Scan(&email, &orderDate, &productName, &quantity, &pricePerUnit, &totalAmount)
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%s\t%s\t%d\t%.2f\t%.2f\n", email, orderDate, productName, quantity, pricePerUnit, totalAmount)
	}

	return nil
}
