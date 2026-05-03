package assembly

import (
	"asamblea/internal/domain"
)

type handler struct {
	service domain.AssemblyService
	units   domain.UnitService
	owners  domain.OwnerService
}

func New(s domain.AssemblyService, u domain.UnitService, o domain.OwnerService) *handler {
	return &handler{service: s, units: u, owners: o}
}
