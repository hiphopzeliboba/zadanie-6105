package tender

import (
	"zadanie-6105/internal/repository"
	def "zadanie-6105/internal/service"
)

var _ def.TenderService = (*serv)(nil)

type serv struct {
	tenderRepository repository.TenderRepository
}

//func (s serv) Create(ctx context.Context, info *models.Tender) (*models.Tender, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewService(tenderRepository repository.TenderRepository) *serv {
	return &serv{tenderRepository: tenderRepository}
}
