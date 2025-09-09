package auth

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"time"

	"github.com/UniBO-PRISMLab/nip/api/identity"
	"github.com/UniBO-PRISMLab/nip/db"
	"github.com/UniBO-PRISMLab/nip/models"
)

type Service struct {
	configuration   models.Configuration
	identityService *identity.Service
	authRepo        *db.AuthRepository
}

func NewService(
	configuration models.Configuration,
	authRepo *db.AuthRepository,
	identityService *identity.Service,
) *Service {
	return &Service{
		configuration:   configuration,
		authRepo:        authRepo,
		identityService: identityService,
	}
}

// https://www.rfc-editor.org/rfc/rfc8017
func (s *Service) IssuePAC(ctx context.Context, req *models.PACRequestModel) (*models.PACResponseModel, error) {
	user, err := s.identityService.GetUserByPID(ctx, &req.PID)
	if err != nil {
		return nil, models.ErrorUserWithPIDNotFound
	}

	pkBytes, err := base64.StdEncoding.DecodeString(user.PublicKey)
	if err != nil {
		return nil, models.ErrorPublicKeyDecoding
	}

	publicKeyPemBlock, _ := pem.Decode(pkBytes)
	if publicKeyPemBlock == nil || publicKeyPemBlock.Type != "PUBLIC KEY" {
		return nil, models.ErrorInvalidPublicKeyHeader
	}

	pk, err := x509.ParsePKIXPublicKey(publicKeyPemBlock.Bytes)
	if err != nil {
		return nil, models.ErrorInvalidPublicKey
	}

	h := crypto.SHA256.New()
	h.Write([]byte(user.PID))
	pidHash := h.Sum(nil)

	sig, err := base64.StdEncoding.DecodeString(req.SignedPID)
	if err != nil {
		return nil, models.ErrorInvalidSignatureEncoding
	}

	if err := rsa.VerifyPKCS1v15(pk.(*rsa.PublicKey), crypto.SHA256, pidHash, sig); err != nil {
		return nil, models.ErrorPIDSignatureVerification
	}

	expiration := time.Now().Add(2 * time.Minute).UTC()

	a, _ := rand.Int(rand.Reader, big.NewInt(900000))
	pac := a.Int64() + 100000

	return s.authRepo.IssuePAC(ctx, &user.PID, pac, expiration)
}

func (s *Service) GetSAC(ctx context.Context) (*models.SACResponseModel, error) {
	return s.authRepo.GetSAC(ctx)
}

func (s *Service) VerifyPAC(ctx context.Context, req *models.PACVerificationRequestModel) (*models.PACVerificationResponseModel, error) {
	user, err := s.identityService.GetUserByPID(ctx, &req.PID)
	if err != nil {
		return nil, models.ErrorUserWithPIDNotFound
	}

	return s.authRepo.VerifyPAC(ctx, &user.PID, req.PAC)
}
