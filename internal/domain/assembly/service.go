package assembly

import "asamblea/internal/domain"

type service struct {
	repo     domain.AssemblyRepository
	unitRepo domain.UnitRepository
}

func New(r domain.AssemblyRepository, u domain.UnitRepository) domain.AssemblyService {
	return &service{repo: r, unitRepo: u}
}
