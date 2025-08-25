package identity

import (
	"context"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"

	"github.com/UniBO-PRISMLab/nip/db"
	"github.com/UniBO-PRISMLab/nip/models"
	"github.com/google/uuid"
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

func (s *Service) IssuePID(ctx context.Context, req *models.PIDRequestModel) (*models.PIDResponseModel, error) {
	h := sha256.New()
	h.Write([]byte(req.PublicKey))
	bs := h.Sum(nil)
	uuid := uuid.New()

	pidBytes := append(bs, uuid[:]...)
	pid := hex.EncodeToString(pidBytes)

	return s.identityRepo.IssuePID(ctx, req.PublicKey, pid)
}
