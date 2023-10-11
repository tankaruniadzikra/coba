package handler

import (
	"database/sql"
	"fmt"
)

func ViewStockReport(db *sql.DB) error {
	rows, err := db.Query("SELECT Products.ProductID, Products.ProductName, Stock.Quantity, OrderItems.PricePerUnit FROM Products JOIN Stock ON Products.ProductID = Stock.ProductID JOIN OrderItems ON Products.ProductID = OrderItems.ProductID")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Stock Report:")
	fmt.Println("ProductID\tProductName\tStock\tPricePerUnit")

	for rows.Next() {
		var productID, quantity int
		var productName string
		var pricePerUnit float64

		err := rows.Scan(&productID, &productName, &quantity, &pricePerUnit)
		if err != nil {
			return err
		}

		fmt.Printf("%d\t%s\t%d\t%.2f\n", productID, productName, quantity, pricePerUnit)
	}

	return nil
}
