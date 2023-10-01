package app

type Order struct {
	ID         int
	CustomerID int
	MenuItemID int
	Quantity   uint
	TotalPrice float64
	OrderDate  string
	Status     string
}
