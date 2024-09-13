package handlers

import "zadanie-6105/internal/service"

type Implementation struct {
	tenderService service.TenderService
}

func NewImplementation(tenderService service.TenderService) *Implementation {
	return &Implementation{tenderService: tenderService}
}
