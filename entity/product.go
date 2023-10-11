package entity

type Product struct {
	ProductID   int
	ProductName string
	Description string
	Price       float64
	Category    Category
	Brand       Brand
	Size        Size
	Color       Color
}

type Category struct {
	CategoryID   int
	CategoryName string
	Description  string
}

type Brand struct {
	BrandID     int
	BrandName   string
	Description string
}

type Size struct {
	SizeID      int
	SizeName    string
	Description string
}

type Color struct {
	ColorID   int
	ColorName string
}
