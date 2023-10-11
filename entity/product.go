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
}

type Brand struct {
	BrandID   int
	BrandName string
}

type Size struct {
	SizeID   int
	SizeName string
}

type Color struct {
	ColorID   int
	ColorName string
}
