package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func DeleteOrderHistory(db *sql.DB, user entity.User, orderID int) error {
	// Periksa apakah orderID milik pengguna
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Orders WHERE OrderID = ? AND UserID = ?", orderID, user.Id).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Tidak dapat menemukan pesanan dengan ID %d", orderID)
	}

	// Hapus catatan terkait dari OrderItems terlebih dahulu
	_, err = db.Exec("DELETE FROM OrderItems WHERE OrderID = ?", orderID)
	if err != nil {
		return err
	}

	// Hapus pesanan dari tabel Orders
	_, err = db.Exec("DELETE FROM Orders WHERE OrderID = ?", orderID)
	if err != nil {
		return err
	}

	fmt.Printf("Pesanan dengan ID %d berhasil dihapus\n", orderID)
	return nil
}
