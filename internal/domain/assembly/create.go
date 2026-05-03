package assembly

import (
	"asamblea/internal/models"
	"context"
	"math/rand"
)

func (s *service) Create(ctx context.Context, assembly models.AssemblyBO) error {
	assembly.Status = "draft"
	assembly.Slug = generateSlug()
	return s.repo.Save(ctx, assembly)
}

func generateSlug() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
