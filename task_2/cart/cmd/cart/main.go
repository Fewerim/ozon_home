package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/app"
	"github.com/Fewerim/ozon_home/task_2/cart/internal/repository/memory"
	productsgrpc "github.com/Fewerim/ozon_home/task_2/cart/internal/repository/products-grpc"
	cartTransportHTTP "github.com/Fewerim/ozon_home/task_2/cart/internal/transport/http"
	"github.com/Fewerim/ozon_home/task_2/cart/pkg/middleware"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
)

const (
	productCatalogTimeout = time.Second * 15
)

func main() {
	cartRepository := memory.NewCartRepository()

	grpcConn, err := productsgrpc.CreateGRPCConnection("localhost:50051")
	if err != nil {
		log.Fatalf("grpc connection: %v", err)
	}
	defer grpcConn.Close()
	client := pb.NewProductServiceClient(grpcConn)
	productsCatalog := productsgrpc.NewProductRepository(productCatalogTimeout, client)

	cartService := app.NewCartService(cartRepository, productsCatalog)

	handler := cartTransportHTTP.NewCartHTTPHandler(cartService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", handler.AddItem)
	mux.HandleFunc("GET /user/{user_id}/cart", handler.GetCart)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.LoggingMiddleware(mux),
	}

	log.Println("cart listening :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
