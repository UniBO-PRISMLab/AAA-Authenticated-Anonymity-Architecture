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
	var resp *models.SACResponseModel

	_, pkBytes, err := s.uip.GetSIDRecord(ctx, req.SID)
	if err != nil {
		return nil, fmt.Errorf("failed to get SID record: %w", err)
	}

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

	sac := make([]byte, 8)
	rand.Read(sac)
	expiration := time.Now().Add(2 * time.Minute).UTC()

	if resp, err = s.authRepo.IssueSAC(
		ctx,
		base64.StdEncoding.EncodeToString(sac),
		&req.SID,
		expiration,
	); err != nil {
		return nil, err
	}

	if err = s.uip.SubmitSAC(ctx, sac); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) VerifyPAC(
	ctx context.Context,
	req *models.PACVerificationRequestModel,
) (*models.PACVerificationResponseModel, error) {
	user, err := s.identityService.GetUserByPID(ctx, &req.PID)
	if err != nil {
		return nil, models.ErrorUserWithPIDNotFound
	}

	return s.authRepo.VerifyPAC(ctx, &user.PID, req.PAC)
}

func (s *Service) VerifySAC(
	ctx context.Context,
	req *models.SACVerificationRequestModel,
) (*models.SACVerificationResponseModel, error) {
	var err error
	var sac []byte
	var pkBytes []byte
	var pemBlock *pem.Block

	if sac, err = s.uip.GetSACRecord(ctx, []byte(req.PublicKey)); err != nil {
		return nil, err
	}

	if base64.StdEncoding.EncodeToString(sac) != req.SAC {
		return nil, models.ErrorSACMismatch
	}

	pkBytes, err = base64.StdEncoding.DecodeString(req.PublicKey)
	if err != nil {
		return nil, models.ErrorPublicKeyDecoding
	}

	pemBlock, _ = pem.Decode(pkBytes)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		return nil, models.ErrorInvalidPublicKeyHeader
	}

	pub, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, models.ErrorInvalidPublicKey
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, models.ErrorInvalidPublicKey
	}

	sig, err := base64.StdEncoding.DecodeString(req.SignedSAC)
	if err != nil {
		return nil, models.ErrorInvalidSignatureEncoding
	}

	sacB64 := base64.StdEncoding.EncodeToString(sac)
	h := crypto.SHA256.New()
	h.Write([]byte(sacB64))
	sacHash := h.Sum(nil)

	if err := rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, sacHash, sig); err != nil {
		return nil, models.ErrorSACSignatureVerification
	}

	return &models.SACVerificationResponseModel{
		Valid: true,
	}, nil
}
