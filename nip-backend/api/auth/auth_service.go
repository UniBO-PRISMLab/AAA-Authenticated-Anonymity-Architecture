package auth

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa"
	"github.com/UniBO-PRISMLab/nip-backend/api/identity"
	"github.com/UniBO-PRISMLab/nip-backend/db"
	"github.com/UniBO-PRISMLab/nip-backend/models"
)

type Service struct {
	configuration   models.Configuration
	identityService *identity.Service
	authRepo        *db.AuthRepository
	uip             *aaa.Service
}

func NewService(
	configuration models.Configuration,
	authRepo *db.AuthRepository,
	identityService *identity.Service,
	uip *aaa.Service,
) *Service {
	return &Service{
		configuration:   configuration,
		authRepo:        authRepo,
		identityService: identityService,
		uip:             uip,
	}
}

// https://www.rfc-editor.org/rfc/rfc8017
func (s *Service) IssuePAC(
	ctx context.Context,
	req *models.PACRequestModel,
) (*models.PACResponseModel, error) {
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
	rsaPub, ok := pk.(*rsa.PublicKey)
	if !ok {
		return nil, models.ErrorInvalidPublicKey
	}

	h := crypto.SHA256.New()
	h.Write([]byte(user.PID))
	pidHash := h.Sum(nil)

	sig, err := base64.StdEncoding.DecodeString(req.SignedPID)
	if err != nil {
		return nil, models.ErrorInvalidSignatureEncoding
	}

	if err := rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, pidHash, sig); err != nil {
		return nil, models.ErrorPIDSignatureVerification
	}

	expiration := time.Now().Add(2 * time.Minute).UTC()

	a, _ := rand.Int(rand.Reader, big.NewInt(900000))
	pac := a.Int64() + 100000

	return s.authRepo.IssuePAC(ctx, &user.PID, pac, expiration)
}

func (s *Service) IssueSAC(
	ctx context.Context,
	req *models.SACRequestModel,
) (*models.SACResponseModel, error) {
	var err error
	var sid string
	var resp *models.SACResponseModel

	_, pkBytes, err := s.uip.GetSIDRecord(ctx, req.SID)
	if err != nil {
		return nil, fmt.Errorf("failed to get SID record: %w", err)
	}

	// Print the retrieved public key in a readable format
	fmt.Printf("Retrieved public key (base64): %s\n", base64.StdEncoding.EncodeToString(pkBytes))

	publicKeyPemBlock, _ := pem.Decode(pkBytes)
	if publicKeyPemBlock == nil || publicKeyPemBlock.Type != "PUBLIC KEY" {
		return nil, models.ErrorInvalidPublicKeyHeader
	}

	pub, err := x509.ParsePKIXPublicKey(publicKeyPemBlock.Bytes)
	if err != nil {
		return nil, models.ErrorInvalidPublicKey
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, models.ErrorInvalidPublicKey
	}

	sig, err := base64.StdEncoding.DecodeString(req.SignedSID)
	if err != nil {
		return nil, models.ErrorInvalidSignatureEncoding
	}

	h := crypto.SHA256.New()
	h.Write([]byte(req.SID))
	sidHash := h.Sum(nil)

	if err := rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, sidHash, sig); err != nil {
		return nil, models.ErrorSIDSignatureVerification
	}

	a, _ := rand.Int(rand.Reader, big.NewInt(900000))
	sac := a.Int64() + 100000
	expiration := time.Now().Add(2 * time.Minute).UTC()

	resp, err = s.authRepo.IssueSAC(ctx, nil, sac, &sid, expiration)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) VerifyPAC(ctx context.Context, req *models.PACVerificationRequestModel) (*models.PACVerificationResponseModel, error) {
	user, err := s.identityService.GetUserByPID(ctx, &req.PID)
	if err != nil {
		return nil, models.ErrorUserWithPIDNotFound
	}

	return s.authRepo.VerifyPAC(ctx, &user.PID, req.PAC)
}
