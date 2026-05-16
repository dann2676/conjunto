package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
)

func (s *service) GenerateCodes(ctx context.Context, assemblyID int) error {
	// obtener todas las unidades activas
	units, err := s.unitRepo.GetAll(ctx, false)
	if err != nil {
		return err
	}
	return s.repo.GenerateCodes(ctx, assemblyID, units)
}

func (s *service) ValidateCode(ctx context.Context, code string, assemblyID int) (models.AssemblyCodeBO, error) {
	ac, err := s.repo.GetCode(ctx, code)
	if err != nil {
		return models.AssemblyCodeBO{}, domain.NotFounErr("código inválido")
	}
	if ac.AssemblyID != assemblyID {
		return models.AssemblyCodeBO{}, domain.AsociationErr("código no pertenece a esta asamblea")
	}
	if ac.Used {
		return models.AssemblyCodeBO{}, domain.DuplicatedErr("código ya utilizado")
	}
	return ac, nil
}

func (s *service) GetCodes(ctx context.Context, assemblyID int) ([]models.AssemblyCodeBO, error) {
	return s.repo.GetCodes(ctx, assemblyID)
}

func (s *service) UseCode(ctx context.Context, code string) error {
	return s.repo.MarkCodeUsed(ctx, code)
}
