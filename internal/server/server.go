package server

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
	"github.com/viveksingh-01/ginger-root/internal/address"
	"github.com/viveksingh-01/ginger-root/internal/auth"
	"github.com/viveksingh-01/ginger-root/internal/cart"
	"github.com/viveksingh-01/ginger-root/internal/config"
	"github.com/viveksingh-01/ginger-root/internal/health"
	"github.com/viveksingh-01/ginger-root/internal/menu"
	"github.com/viveksingh-01/ginger-root/internal/order"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
	"github.com/viveksingh-01/ginger-root/internal/restaurantmenu"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config, db *mongo.Database) *Server {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.AllowedOrigin},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Guest-Id"},
		ExposeHeaders:    []string{"X-Guest-Id"},
		AllowCredentials: true,
	}))

	// Use request-logger
	r.Use(RequestLogger())

	// Add health-check handler
	healthHandler := health.NewHandler()
	r.GET("/health", healthHandler.Check)

	// Create route group
	apiPath := r.Group("/api/v1")

	// Integrate user-auth repository, service and handler
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	auth.RegisterRoutes(apiPath, authHandler)

	// Integrate restaurant's repository, service and handler
	restaurantRepo := restaurant.NewRepository(db)
	restaurantService := restaurant.NewService(restaurantRepo)
	restaurantHandler := restaurant.NewHandler(restaurantService)
	// Register route for restaurant-handler
	restaurant.RegisterRoutes(apiPath, restaurantHandler)

	// Integrate menu's repository, service and handler and register route
	menuRepo := menu.NewRepository(db)
	menuService := menu.NewService(menuRepo)
	menuHandler := menu.NewHandler(menuService)
	menu.RegisterRoutes(apiPath, menuHandler)

	restaurantMenuService := restaurantmenu.NewService(restaurantService, menuService)
	restaurantMenuHandler := restaurantmenu.NewHandler(restaurantMenuService)
	restaurantmenu.RegisterRoutes(apiPath, restaurantMenuHandler)

	cartRepo := cart.NewRepository(db)
	cartService := cart.NewService(cartRepo, menuService, restaurantService, authService)
	cartHandler := cart.NewHandler(cartService)
	cart.RegisterRoutes(apiPath, cartHandler)

	apiPath.Use(AuthMiddleware())
	{
		auth.RegisterMeHandler(apiPath, authHandler)

		// Register Address route
		addressRepo := address.NewRepository(db)
		addressService := address.NewService(addressRepo)
		addressHandler := address.NewHandler(addressService)
		address.RegisterRoutes(apiPath, addressHandler)

		orderRepo := order.NewRepository(db)
		orderService := order.NewService(orderRepo, cartRepo, menuService, restaurantService, addressService)
		orderHandler := order.NewHandler(orderService)
		order.RegisterRoutes(apiPath, orderHandler)
	}

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	return &Server{httpServer: server}
}

func (s *Server) Start() {
	// Channel to listen for OS signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		log.Printf("Server listening on %s", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
	s.httpServer.Shutdown(ctx)
	log.Println("Server exited successfully.")
}
