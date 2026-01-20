package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/viveksingh-01/ginger-root/internal/config"
	"github.com/viveksingh-01/ginger-root/internal/database"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
)

func main() {
	log.Println("Welcome to Ginger API.")
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Load configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err.Error())
	}

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.AllowedOrigin},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Create route group
	apiPath := r.Group("/api/v1")

	// Connect to DB
	mongoClient, err := database.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err.Error())
	}
	db := mongoClient.Database(cfg.Database)

	// Integrate restaurant's repository, service and handler
	restaurantRepo := restaurant.NewRepository(db)
	restaurantService := restaurant.NewService(restaurantRepo)
	restaurantHandler := restaurant.NewHandler(restaurantService)
	// Register route for restaurant-handler
	restaurant.RegisterRoutes(apiPath, restaurantHandler)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Channel to listen for OS signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		log.Printf("Server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start the server: %s", err)
		}
	}()

	// Block until signal received
	<-shutdown
	log.Println("Shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown with timeout
	server.Shutdown(ctx)
	log.Println("Server exited successfully.")
}
