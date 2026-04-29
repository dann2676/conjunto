package assembly

import "asamblea/internal/domain"

type service struct {
	repo domain.AssemblyRepository
}

func New(r domain.AssemblyRepository) domain.AssemblyService {
	return &service{repo: r}
}
