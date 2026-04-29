package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func (r *repository) RegisterVote(ctx context.Context, vote models.VoteBO) error {
	entity := mapVoteToEntity(vote)
	if err := r.db.Save(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.DuplicatedErr("voto")
		}
		derr := domain.SavingErr("voto")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}

func (r *repository) GetVotes(ctx context.Context, agendaItemID int) ([]models.VoteBO, error) {
	var entities []models.VoteEntity
	err := r.db.Preload("Unit").Where("agenda_item_id = ?", agendaItemID).Find(&entities).Error
	if err != nil {
		derr := domain.NotFounErr("votos")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	result := make([]models.VoteBO, len(entities))
	for i, e := range entities {
		result[i] = mapEntityToVote(e)
	}
	return result, nil
}
