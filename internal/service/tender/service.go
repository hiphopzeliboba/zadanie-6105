package tender

import (
	"zadanie-6105/internal/repository"
	def "zadanie-6105/internal/service"
)

var _ def.TenderService = (*serv)(nil)

type serv struct {
	tenderRepository repository.TenderRepository
}

func NewService(tenderRepository repository.TenderRepository) *serv {
	return &serv{tenderRepository: tenderRepository}
}
