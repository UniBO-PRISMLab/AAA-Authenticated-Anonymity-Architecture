package identity

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"

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

func (s *Service) IssuePID(ctx context.Context, req *models.PIDRequestModel) (*models.PIDResponseModel, error) {
	h := sha256.New()
	h.Write([]byte(req.PublicKey))
	pkSum := h.Sum(nil)

	nonce := make([]byte, 32)
	rand.Read(nonce)

	key := s.configuration.SK

	// Compute f_k(r)
	mac := hmac.New(sha256.New, key)
	mac.Write(nonce)
	macKR := mac.Sum(nil)

	// Compute pid as
	// pid = f_k(r) ^ pk
	pid := make([]byte, 32)
	for i := range macKR {
		pid[i] = macKR[i] ^ pkSum[i]
	}

	return s.identityRepo.IssuePID(ctx, req.PublicKey, base64.StdEncoding.EncodeToString(pid), base64.StdEncoding.EncodeToString(nonce))
}
