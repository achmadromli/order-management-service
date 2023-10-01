package app

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) CreateOrder(order *Order) (*Order, error) {
	// Validasi pesanan dan hitung total harga, lalu simpan ke database
	// ...
	return s.repo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(orderID int) (*Order, error) {
	return s.repo.GetOrderByID(orderID)
}

func (s *OrderService) UpdateOrderStatus(orderID int, status string) error {
	return s.repo.UpdateOrderStatus(orderID, status)
}

func (s *OrderService) DeleteOrder(orderID int) error {
	return s.repo.DeleteOrder(orderID)
}
