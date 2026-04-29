package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) RegisterVote(ctx context.Context, vote models.VoteBO) error {
	return s.repo.RegisterVote(ctx, vote)
}

func (s *service) GetVotes(ctx context.Context, agendaItemID int) ([]models.VoteBO, error) {
	return s.repo.GetVotes(ctx, agendaItemID)
}
