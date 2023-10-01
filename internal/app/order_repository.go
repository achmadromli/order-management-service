package app

// OrderRepository adalah antarmuka repository untuk pesanan
type OrderRepository interface {
	CreateOrder(order *Order) (*Order, error)
	GetOrderByID(orderID int) (*Order, error)
	UpdateOrderStatus(orderID int, status string) error
	DeleteOrder(orderID int) error
}
