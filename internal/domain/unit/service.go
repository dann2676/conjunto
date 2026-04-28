package unit

import "asamblea/internal/domain"

type service struct {
	repo domain.UnitRepository
}

func New(r domain.UnitRepository) domain.UnitService {
	return &service{repo: r}
}
