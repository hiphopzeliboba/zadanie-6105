package repository

import (
	"context"
	"zadanie-6105/internal/models"
)

type TenderRepository interface {
	Create(ctx context.Context, t *models.Tender) (*models.Tender, error)
}
