package sample

import (
	"context"
	"log"

	"github.com/takokun778/oreoreddd/internal/domain/repository"
	"github.com/takokun778/oreoreddd/internal/usecase/port"
)

var _ port.SampleUpdateUsecase = (*UpdateInteractor)(nil)

type UpdateInteractor struct {
	repository updateRepository
}

type updateRepository interface {
	repository.SampleFinder
	repository.SampleSaver
}

func NewUpdateInteractor(
	repository updateRepository,
) *UpdateInteractor {
	return &UpdateInteractor{
		repository: repository,
	}
}

func (ui *UpdateInteractor) Execute(ctx context.Context, input port.SampleUpdateInput) (port.SampleUpdateOutput, error) {
	src, err := ui.repository.Find(ctx, input.ID)
	if err != nil {
		log.Printf("Failed to find sample: %s", err)

		return port.SampleUpdateOutput{}, err
	}

	dst := src.Update(input.Name)

	if err := ui.repository.Save(ctx, dst); err != nil {
		log.Printf("Failed to save sample: %s", err)

		return port.SampleUpdateOutput{}, err
	}

	return port.SampleUpdateOutput{
		Sample: dst,
	}, nil
}
