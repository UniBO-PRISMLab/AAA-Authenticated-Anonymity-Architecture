package auth

import (
	"context"

	"github.com/UniBO-PRISMLab/nip/db"
	"github.com/UniBO-PRISMLab/nip/models"
)

type Service struct {
	configuration models.Configuration
	authRepo      *db.AuthRepository
}

func NewService(
	configuration models.Configuration,
	authRepo *db.AuthRepository,
) *Service {
	return &Service{
		configuration: configuration,
		authRepo:      authRepo,
	}
}

func (s *Service) GetPAC(ctx context.Context) (*models.PACResponseModel, error) {
	return s.authRepo.GetPAC(ctx)
}

func (s *Service) GetSAC(ctx context.Context) (*models.SACResponseModel, error) {
	return s.authRepo.GetSAC(ctx)
}
