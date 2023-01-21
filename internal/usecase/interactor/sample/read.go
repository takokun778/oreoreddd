package sample

import (
	"context"
	"log"

	"github.com/takokun778/oreoreddd/internal/domain/repository"
	"github.com/takokun778/oreoreddd/internal/usecase/port"
)

var _ port.SampleReadUsecase = (*ReadInteractor)(nil)

type ReadInteractor struct {
	repository readRepository
}

type readRepository interface {
	repository.SampleFinder
}

func NewReadInteractor(
	repository readRepository,
) *ReadInteractor {
	return &ReadInteractor{
		repository: repository,
	}
}

func (ri *ReadInteractor) Execute(ctx context.Context, input port.SampleReadInput) (port.SampleReadOutput, error) {
	model, err := ri.repository.Find(ctx, input.ID)
	if err != nil {
		log.Printf("Failed to find sample: %v", err)

		return port.SampleReadOutput{}, err
	}

	return port.SampleReadOutput{
		Sample: model,
	}, nil
}
