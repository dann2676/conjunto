package assembly

import (
	"asamblea/internal/domain"
	"context"
)

func (s *service) UpdateStatus(ctx context.Context, id int, newStatus string) error {
	assembly, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	valid := map[string]string{
		"draft": "open",
		"open":  "closed",
	}
	if valid[assembly.Status] != newStatus {
		return domain.DuplicatedErr("transición de estado inválida")
	}
	assembly.Status = newStatus
	if err = s.repo.Save(ctx, assembly); err != nil {
		return err
	}
	// generar códigos al abrir
	if newStatus == "open" {
		return s.GenerateCodes(ctx, id)
	}
	return nil
}
