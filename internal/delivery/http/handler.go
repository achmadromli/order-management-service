package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"double-dose/order-management/internal/app"
)

// OrderHandler adalah handler HTTP untuk pesanan
type OrderHandler struct {
	service *app.OrderService
}

// NewOrderHandler membuat instance OrderHandler
func NewOrderHandler(service *app.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

// CreateOrderHandler menangani permintaan pembuatan pesanan
func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil data pesanan dari permintaan HTTP
	var orderData app.Order
	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Membuat pesanan menggunakan service
	order, err := h.service.CreateOrder(&orderData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirimkan respons JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrderHandler menangani permintaan pengambilan pesanan berdasarkan ID
func (h *OrderHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("order_id")
	orderID, _ := strconv.Atoi(orderIDStr)

	order, err := h.service.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Mengirimkan respons JSON
	json.NewEncoder(w).Encode(order)
}

// ...
// Implementasi handler lainnya untuk operasi seperti UpdateOrder dan DeleteOrder
