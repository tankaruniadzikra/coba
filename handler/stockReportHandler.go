package handler

import (
	"database/sql"
	"pair-project/entity"
)

func GetMaxStockProduct(db *sql.DB) (entity.Product, error) {
	var maxProduct entity.Product

	// Query untuk mendapatkan produk dengan stok terbanyak
	err := db.QueryRow(`
		SELECT p.ProductID, p.ProductName, p.Description, p.Price, p.Material, p.Weight,
			b.BrandID, b.BrandName, sz.SizeID, sz.SizeName, co.ColorID, co.ColorName,
			s.Quantity
		FROM Products p
		JOIN Brands b ON p.BrandID = b.BrandID
		JOIN Sizes sz ON p.SizeID = sz.SizeID
		JOIN Colors co ON p.ColorID = co.ColorID
		JOIN Stock s ON p.ProductID = s.ProductID
		ORDER BY s.Quantity DESC
		LIMIT 1;
	`).Scan(
		&maxProduct.ProductID, &maxProduct.ProductName, &maxProduct.Description,
		&maxProduct.Price, &maxProduct.Material, &maxProduct.Weight,
		&maxProduct.Brand.BrandID, &maxProduct.Brand.BrandName,
		&maxProduct.Size.SizeID, &maxProduct.Size.SizeName,
		&maxProduct.Color.ColorID, &maxProduct.Color.ColorName,
		&maxProduct.Stock.Quantity,
	)
	if err != nil {
		return entity.Product{}, err
	}

	return maxProduct, nil
}

func GetMinStockProduct(db *sql.DB) (entity.Product, error) {
	var minProduct entity.Product

	// Query untuk mendapatkan produk dengan stok paling sedikit
	err := db.QueryRow(`
		SELECT p.ProductID, p.ProductName, p.Description, p.Price, p.Material, p.Weight,
			b.BrandID, b.BrandName, sz.SizeID, sz.SizeName, co.ColorID, co.ColorName,
			s.Quantity
		FROM Products p
		JOIN Brands b ON p.BrandID = b.BrandID
		JOIN Sizes sz ON p.SizeID = sz.SizeID
		JOIN Colors co ON p.ColorID = co.ColorID
		JOIN Stock s ON p.ProductID = s.ProductID
		ORDER BY s.Quantity ASC
		LIMIT 1;
	`).Scan(
		&minProduct.ProductID, &minProduct.ProductName, &minProduct.Description,
		&minProduct.Price, &minProduct.Material, &minProduct.Weight,
		&minProduct.Brand.BrandID, &minProduct.Brand.BrandName,
		&minProduct.Size.SizeID, &minProduct.Size.SizeName,
		&minProduct.Color.ColorID, &minProduct.Color.ColorName,
		&minProduct.Stock.Quantity,
	)
	if err != nil {
		return entity.Product{}, err
	}

	return minProduct, nil
}
