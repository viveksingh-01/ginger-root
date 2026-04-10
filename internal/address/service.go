package address

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateAddress(ctx context.Context, req *CreateAddressRequest, userId string) (*Address, error) {
	uid, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	address := &Address{
		UserID:     uid,
		Name:       req.Name,
		Phone:      req.Phone,
		Annotation: req.Annotation,
		Address:    req.Address,
		House:      req.House,
		Area:       req.Area,
		City:       req.City,
		Landmark:   req.Landmark,
		Lat:        req.Lat,
		Lng:        req.Lng,
	}

	return s.repo.Create(ctx, address)
}

func (s *Service) GetAddresses(ctx context.Context, userId string) ([]*Address, error) {
	uid, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByUserID(ctx, uid)
}
