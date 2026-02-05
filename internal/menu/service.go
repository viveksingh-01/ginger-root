package menu

import "context"

type Service struct {
	repository Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repository: *repo}
}

func (s *Service) GetMenu(ctx context.Context, restaurantID string) (*Menu, error) {
	return s.repository.FindByRestaurantID(ctx, restaurantID)
}
