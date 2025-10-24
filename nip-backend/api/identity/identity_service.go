package identity

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/pem"

	"github.com/UniBO-PRISMLab/nip-backend/db"
	"github.com/UniBO-PRISMLab/nip-backend/models"
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
	var pemBlock *pem.Block
	var pkBytes []byte
	var err error

	// TODO: validate user personal data with public key

	pkBytes, err = base64.StdEncoding.DecodeString(req.PublicKey)
	if err != nil {
		return nil, models.ErrorPublicKeyDecoding
	}

	pemBlock, _ = pem.Decode(pkBytes)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		return nil, models.ErrorInvalidPublicKeyHeader
	}

	// compute PK digest (32 bytes)
	h := sha256.New()
	h.Write(pkBytes)
	pkSum := h.Sum(nil)

	// generate nonce (32 bytes)
	nonce := make([]byte, 32)
	rand.Read(nonce)

	// derive pid = HMAC(SK, pkSum || nonce)
	key := s.configuration.SK
	mac := hmac.New(sha256.New, key)
	mac.Write(pkSum[:])
	mac.Write(nonce)
	pid := mac.Sum(nil)

	return s.identityRepo.IssuePID(
		ctx,
		req.PublicKey,
		base64.StdEncoding.EncodeToString(pid),
		base64.StdEncoding.EncodeToString(nonce),
	)
}

func (s *Service) GetUserByPID(ctx context.Context, PID *string) (*models.User, error) {
	return s.identityRepo.GetUserByPID(ctx, PID)
}
