package unit

import "asamblea/internal/domain"

type handler struct {
	service domain.UnitService
	owners  domain.OwnerService
}

func New(s domain.UnitService, o domain.OwnerService) *handler {
	return &handler{service: s, owners: o}
}
