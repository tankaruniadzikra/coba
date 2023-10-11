package entity

type Order struct {
	OrderID     int
	UserID      int
	OrderDate   string
	TotalAmount float64
}

type OrderItem struct {
	OrderItemID  int
	OrderID      int
	ProductID    int
	Quantity     int
	PricePerUnit float64
	TotalPrice   float64
}

type OrderHistory struct {
	OrderID     int
	ProductName string
	Quantity    int
	OrderDate   string
	TotalAmount float64
}

type OrderReport struct {
	OrderDay      string
	TotalQuantity int
	TotalSales    float64
}
