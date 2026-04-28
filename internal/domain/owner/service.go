package owner

import "asamblea/internal/domain"

type service struct {
	repo domain.OwnerRepository
}

func New(r domain.OwnerRepository) domain.OwnerService {
	return &service{repo: r}
}
