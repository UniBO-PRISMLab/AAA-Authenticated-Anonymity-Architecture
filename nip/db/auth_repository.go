package db

import (
	"context"
	_ "embed"

	"github.com/UniBO-PRISMLab/nip/models"
)

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) GetPAC(ctx context.Context) (*models.PACResponseModel, error) {
	return &models.PACResponseModel{}, nil
}

func (r *AuthRepository) GetSAC(ctx context.Context) (*models.SACResponseModel, error) {
	return &models.SACResponseModel{}, nil
}
