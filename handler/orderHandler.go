package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func PlaceOrder(db *sql.DB, user entity.User, ProductID, Quantity int) error {
	// Step 1: Buat order baru untuk pengguna
	result, err := db.Exec("INSERT INTO Orders (UserID, OrderDate, TotalAmount) VALUES (?, NOW(), 0)", user.Id)
	if err != nil {
		return err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Step 2: Dapatkan nama produk dari ID produk
	var productName string
	err = db.QueryRow("SELECT ProductName FROM Products WHERE ProductID = ?", ProductID).Scan(&productName)
	if err != nil {
		return err
	}

	// Dapatkan harga per unit dari produk
	var pricePerUnit float64
	err = db.QueryRow("SELECT Price FROM Products WHERE ProductID = ?", ProductID).Scan(&pricePerUnit)
	if err != nil {
		return err
	}

	// Hitung total harga dari pesanan
	totalPrice := pricePerUnit * float64(Quantity)

	// Tambahkan item ke OrderItems
	_, err = db.Exec("INSERT INTO OrderItems (OrderID, ProductID, Quantity, PricePerUnit, TotalPrice) VALUES (?, ?, ?, ?, ?)",
		orderID, ProductID, Quantity, pricePerUnit, totalPrice)
	if err != nil {
		return err
	}

	// Step 3: Perbarui TotalAmount di Orders
	_, err = db.Exec("UPDATE Orders SET TotalAmount = ? WHERE OrderID = ?", totalPrice, orderID)
	if err != nil {
		return err
	}

	// Step 4: Kurangi stok produk
	_, err = db.Exec("UPDATE Stock SET Quantity = Quantity - ? WHERE ProductID = ?", Quantity, ProductID)
	if err != nil {
		return err
	}

	// Menampilkan rincian pesanan
	fmt.Printf("Produk '%s' (Quantity: %d) berhasil dipesan. Total Harga: %.2f\n", productName, Quantity, totalPrice)

	return nil
}
