package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"double-dose/order-management/config"
	"double-dose/order-management/internal/app"
	orderHttp "double-dose/order-management/internal/delivery/http"
	"double-dose/order-management/internal/repository/postgres"
)

func main() {
	cfg := config.LoadConfig()

	var db *sql.DB

	orderRepo := postgres.NewPostgreSQLRepository(db)
	orderService := app.NewOrderService(orderRepo)
	orderHandler := orderHttp.NewOrderHandler(orderService)

	http.HandleFunc("/", orderHandler.CreateOrderHandler)
	http.HandleFunc("/:order_id", orderHandler.GetOrderHandler)

	go func() {
		log.Printf("Server started on port %s...", cfg.Server.Port)
		if err := http.ListenAndServe(":"+cfg.Server.Port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
}
