package handler

import (
	"database/sql"
	"fmt"
)

func ViewUserReport(db *sql.DB, userID int) error {
	rows, err := db.Query(`
        SELECT 
            u.UserID, 
            u.Email, 
            p.ProductName, 
            oi.Quantity
        FROM 
            Users u
        JOIN 
            Orders o ON u.UserID = o.UserID
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

	fmt.Println("User Report:")
	fmt.Println("UserID\tEmail\tProductName\tQuantity")

	for rows.Next() {
		var userID, quantity int
		var email, productName string
		err := rows.Scan(&userID, &email, &productName, &quantity)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t%s\t%s\t%d\n", userID, email, productName, quantity)
	}

	return nil
}
