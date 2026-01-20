package main

import (
	"log"
	"net/http"

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

	log.Printf("Server listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
