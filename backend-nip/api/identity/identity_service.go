package identity

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/pem"

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
	var pemBlock *pem.Block
	var pkBytes []byte
	var err error

	// convert the public key from base64 to bytes
	pkBytes, err = base64.StdEncoding.DecodeString(req.PublicKey)
	if err != nil {
		return nil, models.ErrorPublicKeyDecoding
	}

	// Check if the public key is valid
	pemBlock, _ = pem.Decode(pkBytes)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		return nil, models.ErrorInvalidPublicKeyHeader
	}

	// pub, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pkBytes, err = x509.MarshalPKIXPublicKey(pub)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	h := sha256.New()
	h.Write(pkBytes)
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
