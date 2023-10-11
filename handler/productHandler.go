package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func DisplayProducts(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT 
			p.ProductID, 
			p.ProductName, 
			p.Description, 
			p.Price, 
			c.CategoryID, 
			c.CategoryName, 
			b.BrandID, 
			b.BrandName, 
			sz.SizeID, 
			sz.SizeName, 
			co.ColorID, 
			co.ColorName 
		FROM Products p
		JOIN Categories c ON p.CategoryID = c.CategoryID
		JOIN Brands b ON p.BrandID = b.BrandID
		JOIN Sizes sz ON p.SizeID = sz.SizeID
		JOIN Colors co ON p.ColorID = co.ColorID
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var product entity.Product
	fmt.Printf("\n[ID] - [Name] - [Description] - [Price] - [CategoryID] - [CategoryName] - [BrandID] - [BrandName] - [SizeID] - [SizeName] - [ColorID] - [ColorName]\n")
	for rows.Next() {
		err := rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.Category.CategoryID,
			&product.Category.CategoryName,
			&product.Brand.BrandID,
			&product.Brand.BrandName,
			&product.Size.SizeID,
			&product.Size.SizeName,
			&product.Color.ColorID,
			&product.Color.ColorName,
		)
		if err != nil {
			return err
		}

		fmt.Printf("[%v] - [%v] - [%v] - [%.2f] - [%v] - [%v] - [%v] - [%v] - [%v] - [%v] - [%v] - [%v]\n",
			product.ProductID,
			product.ProductName,
			product.Description,
			product.Price,
			product.Category.CategoryID,
			product.Category.CategoryName,
			product.Brand.BrandID,
			product.Brand.BrandName,
			product.Size.SizeID,
			product.Size.SizeName,
			product.Color.ColorID,
			product.Color.ColorName,
		)
	}
	fmt.Println()
	return nil
}
