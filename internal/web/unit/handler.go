package unit

import "asamblea/internal/domain"

type handler struct {
	service domain.UnitService
}

func New(s domain.UnitService) *handler {
	return &handler{service: s}
}
