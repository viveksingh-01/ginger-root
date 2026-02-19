package restaurant

import "context"

// This code defines a Service layer that:
// * Sits between controllers/handlers and the database
// * Delegates data access to a Repository
// * Provides a clean API for restaurant-related operations
// * The service layer is “The gatekeeper of business rules”

// Service has a dependency on Repository
// The service knows exactly how data is accessed
// Note: The dependency is a concrete type, not an interface
type Service struct {
	repository *Repository
}

// 1. Accepts a pointer to a Repository
// 2. Stores it inside the service
// 3. Returns a ready-to-use Service
// This is dependency injection, but with a concrete dependency.
func NewService(repo *Repository) *Service {
	return &Service{repository: repo}
}

const (
	defaultLimit int64 = 20
	maxLimit     int64 = 100
)

// 1. The caller (e.g., an HTTP handler) calls Service.List
// 2. The service forwards the call to Repository.List
// 3. The repository queries the database
// 4. Results are returned back up the stack
func (s *Service) List(ctx context.Context, skip, limit int64, filter Filter, sort Sort) ([]Restaurant, error) {
	if limit < 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if skip <= 0 {
		skip = 1
	}
	return s.repository.List(ctx, skip, limit, filter, sort)
}

func (s *Service) GetRestaurant(ctx context.Context, restaurantId string) (*Restaurant, error) {
	return s.repository.FindByID(ctx, restaurantId)
}

func (s *Service) Search(ctx context.Context, query string) ([]Restaurant, error) {
	return s.repository.Search(ctx, query)
}
