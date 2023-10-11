package handler

import (
	"database/sql"
	"pair-project/entity"
)

func TopPurchasingUsers(db *sql.DB) ([]entity.User, error) {
	rows, err := db.Query(`
		SELECT 
			u.UserID, 
			u.Email, 
			u.FirstName, 
			u.LastName, 
			COALESCE(SUM(o.TotalAmount), 0) AS TotalPurchase
		FROM Users u
		LEFT JOIN Orders o ON u.UserID = o.UserID
		GROUP BY u.UserID
		ORDER BY TotalPurchase DESC
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topUsers []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.TotalPurchase)
		if err != nil {
			return nil, err
		}
		topUsers = append(topUsers, user)
	}

	return topUsers, nil
}
