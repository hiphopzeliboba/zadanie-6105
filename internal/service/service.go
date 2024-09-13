package service

import (
	"context"
	"zadanie-6105/internal/models"
)

type TenderService interface {
	Create(ctx context.Context, info *models.Tender) (*models.Tender, error)
}
