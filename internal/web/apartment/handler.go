package apartment

import "asamblea/internal/domain"

type handler struct {
	service domain.ApartmentService
}

func New(s domain.ApartmentService) *handler {
	return &handler{service: s}
}
