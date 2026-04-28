package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
)

type handler struct {
	service domain.OwnerService
	units   domain.Gettable[models.UnitBO]
}

func New(s domain.OwnerService, a domain.Gettable[models.UnitBO]) *handler {
	return &handler{service: s, units: a}
}
