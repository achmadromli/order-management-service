package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"double-dose/order-management/internal/app"
)

type OrderHandler struct {
	service *app.OrderService
}

func NewOrderHandler(service *app.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var orderData app.Order
	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	order, err := h.service.CreateOrder(&orderData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("order_id")
	orderID, _ := strconv.Atoi(orderIDStr)

	order, err := h.service.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// ...
// Implementasi handler lainnya untuk operasi seperti UpdateOrder dan DeleteOrder
