package usecase

import (
	"double-dose/order-management/internal/app"
)

type OrderUseCase struct {
	service *app.OrderService
}

func NewOrderUseCase(service *app.OrderService) *OrderUseCase {
	return &OrderUseCase{service}
}

func (uc *OrderUseCase) CreateOrderUseCase(orderData *app.Order) (*app.Order, error) {
	// Validasi data pesanan dan logika bisnis lainnya
	// ...
	return uc.service.CreateOrder(orderData)
}

func (uc *OrderUseCase) GetOrderUseCase(orderID int) (*app.Order, error) {
	return uc.service.GetOrderByID(orderID)
}
