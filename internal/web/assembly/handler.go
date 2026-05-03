package assembly

import (
	"asamblea/internal/domain"
)

type handler struct {
	service domain.AssemblyService
	units   domain.UnitService
}

func New(s domain.AssemblyService, u domain.UnitService) *handler {
	return &handler{service: s, units: u}
}
