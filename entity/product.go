package entity

// Product represents a product entity.
type Product struct {
	ProductID   int
	ProductName string
	Description string
	Price       float64
	Material    string
	Weight      float64
	Brand       Brand
	Size        Size
	Color       Color
	Stock       Stock
	Category    Category
}

// Brand represents a brand entity.
type Brand struct {
	BrandID     int
	BrandName   string
	Description string
}

// Size represents a size entity.
type Size struct {
	SizeID      int
	SizeName    string
	Description string
}

// Color represents a color entity.
type Color struct {
	ColorID   int
	ColorName string
}

// Stock represents a stock entity.
type Stock struct {
	StockID   int
	ProductID int
	Quantity  int
}

// Category represents a category entity.
type Category struct {
	CategoryID   int
	CategoryName string
	Description  string
}
