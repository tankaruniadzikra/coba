package handler

import (
	"database/sql"
	"pair-project/entity"
)

func TopSellingDayReport(db *sql.DB) ([]entity.OrderReport, error) {
	rows, err := db.Query(`
		SELECT DATE(OrderDate) AS OrderDay, 
		       SUM(Quantity) AS TotalQuantity, 
		       COALESCE(SUM(TotalAmount), 0) AS TotalSales
		FROM Orders o
		JOIN OrderItems oi ON o.OrderID = oi.OrderID
		GROUP BY DATE(OrderDate)
		ORDER BY TotalSales DESC
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderReports []entity.OrderReport

	for rows.Next() {
		var orderReport entity.OrderReport
		err := rows.Scan(&orderReport.OrderDay, &orderReport.TotalQuantity, &orderReport.TotalSales)
		if err != nil {
			return nil, err
		}
		orderReports = append(orderReports, orderReport)
	}

	return orderReports, nil
}
