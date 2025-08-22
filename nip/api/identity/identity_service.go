package identity

import (
	"context"

	"github.com/UniBO-PRISMLab/nip/db"
	"github.com/UniBO-PRISMLab/nip/models"
)

type Service struct {
	configuration models.Configuration
	identityRepo  *db.IdentityRepository
}

func NewService(
	configuration models.Configuration,
	identityRepo *db.IdentityRepository,
) *Service {
	return &Service{
		configuration: configuration,
		identityRepo:  identityRepo,
	}
}

func (s *Service) GetPID(ctx context.Context) (*models.PIDResponseModel, error) {
	return s.identityRepo.GetPID(ctx)
}
