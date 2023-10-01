package app

// OrderService adalah service yang berisi logika bisnis untuk pesanan
type OrderService struct {
	repo OrderRepository
}

// NewOrderService membuat instance OrderService
func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo}
}

// CreateOrder membuat pesanan baru
func (s *OrderService) CreateOrder(order *Order) (*Order, error) {
	// Validasi pesanan dan hitung total harga, lalu simpan ke database
	// ...
	return s.repo.CreateOrder(order)
}

// GetOrderByID mengambil pesanan berdasarkan ID
func (s *OrderService) GetOrderByID(orderID int) (*Order, error) {
	return s.repo.GetOrderByID(orderID)
}

// UpdateOrderStatus mengubah status pesanan
func (s *OrderService) UpdateOrderStatus(orderID int, status string) error {
	return s.repo.UpdateOrderStatus(orderID, status)
}

// DeleteOrder menghapus pesanan
func (s *OrderService) DeleteOrder(orderID int) error {
	return s.repo.DeleteOrder(orderID)
}
