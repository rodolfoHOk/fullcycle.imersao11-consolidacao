package usecase

import (
	"context"
	"time"

	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/domain/entity"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/domain/repository"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/pkg/uow"
)

type MatchInput struct {
	ID string 
	Date time.Time
	TeamAID string
	TeamBID string
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func (u *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	err := u.Uow.Do(ctx, func(uow *uow.Uow) error {
		matchRepository := u.getMatchRepository(ctx)
		teamRepository := u.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}
		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)

		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (u *MatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := u.Uow.GetRepository(ctx, "TeamRepository")
	if err != nil {
		panic(err)
	}
	return teamRepository.(repository.TeamRepositoryInterface)
}
