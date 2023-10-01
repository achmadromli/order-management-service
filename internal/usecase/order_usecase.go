package usecase

import (
	"double-dose/order-management/internal/app"
)

// OrderUseCase adalah use case atau bisnis logic untuk pesanan
type OrderUseCase struct {
	service *app.OrderService
}

// NewOrderUseCase membuat instance OrderUseCase
func NewOrderUseCase(service *app.OrderService) *OrderUseCase {
	return &OrderUseCase{service}
}

// CreateOrderUseCase adalah use case untuk membuat pesanan
func (uc *OrderUseCase) CreateOrderUseCase(orderData *app.Order) (*app.Order, error) {
	// Validasi data pesanan dan logika bisnis lainnya
	// ...
	return uc.service.CreateOrder(orderData)
}

// GetOrderUseCase adalah use case untuk mengambil pesanan berdasarkan ID
func (uc *OrderUseCase) GetOrderUseCase(orderID int) (*app.Order, error) {
	return uc.service.GetOrderByID(orderID)
}

// UpdateOrderStatusUseCase adalah
