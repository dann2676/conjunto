package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
)

type handler struct {
	service   domain.OwnerService
	apartemts domain.Gettable[models.ApartmentBO]
}

func New(s domain.OwnerService, a domain.Gettable[models.ApartmentBO]) *handler {
	return &handler{service: s, apartemts: a}
}
