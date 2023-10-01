package main

import (
	"fmt"
	"log"

	"double-dose/order-management/config"
	"double-dose/order-management/internal/app"
	"double-dose/order-management/internal/delivery/http"
	"double-dose/order-management/internal/repository/postgres"
)

func main() {
	// Membaca konfigurasi aplikasi
	cfg := config.LoadConfig()

	// Membuat instance OrderRepository (implementasinya dapat berbeda, seperti PostgreSQL atau MySQL)
	orderRepo := postgres.NewPostgreSQLRepository()

	// Membuat instance OrderService
	orderService := app.NewOrderService(orderRepo)

	// Membuat router HTTP
	router := http.NewRouter(orderService)

	// Mengatur server HTTP
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: router,
	}

	// Menjalankan server
	log.Printf("Server started on port %s...", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
