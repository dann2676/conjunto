package assembly

import "asamblea/internal/domain"

type handler struct {
	service domain.AssemblyService
}

func New(s domain.AssemblyService) *handler {
	return &handler{service: s}
}
