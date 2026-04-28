package apartment

import "asamblea/internal/domain"

type service struct {
	repo domain.ApartmentRepository
}

func New(r domain.ApartmentRepository) domain.ApartmentService {
	return &service{repo: r}
}
