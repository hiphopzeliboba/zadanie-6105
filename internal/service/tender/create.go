package tender

import (
	"context"
	"zadanie-6105/internal/models"
)

const (
	tableName = " tenders"

	idColumn    = "id"
	titleColumn = ""
)

// Ошибка не тот слой
func (s *serv) Create(ctx context.Context, info *models.Tender) (*models.Tender, error) {
	id, err := s.tenderRepository.Create(ctx, info)
	if err != nil {
		t := &models.Tender{}
		return t, err
	}
	return id, nil
}
