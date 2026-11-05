package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KolesnikDmitriy/item/internal/analytics"
	"github.com/KolesnikDmitriy/item/internal/app"
	"github.com/KolesnikDmitriy/item/internal/storage"
	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dsn := os.Getenv("ITEM_DB_URL")
	if dsn == "" {
		log.Fatalf("ITEM_DB_URL is required")
	}

	db, err := storage.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)

	an := analytics.New()
	go an.Run(ctx)

	svc := app.NewItemService(db, an)
	pb.RegisterItemServer(server, svc)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to Serve: %v", err)
		}
	}()

	mux := runtime.NewServeMux()
	err = pb.RegisterItemHandlerServer(ctx, mux, svc)
	if err != nil {
		log.Fatalf("failed to RegisterItemHandlerServer: %v", err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("GET /docs/{$}", serveBytes("text/html; charset=utf-8", docsHTML))
	httpMux.Handle("GET /swagger.json", serveBytes("application/json", pb.SwaggerJSON))
	httpMux.Handle("/", mux)

	if err := http.ListenAndServe(":50052", httpMux); err != nil {
		log.Fatalf("failed to ListenAndServe: %v", err)
	}
}

func serveBytes(contentType string, body []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(body)
	})
}
