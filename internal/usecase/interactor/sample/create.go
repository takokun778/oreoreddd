package sample

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/takokun778/oreoreddd/internal/domain/model"
	"github.com/takokun778/oreoreddd/internal/domain/model/sample"
	"github.com/takokun778/oreoreddd/internal/domain/repository"
	"github.com/takokun778/oreoreddd/internal/usecase/port"
)

var _ port.SampleCreateUsecase = (*CreateInteractor)(nil)

type CreateInteractor struct {
	repository createRepository
}

type createRepository interface {
	repository.SampleSaver
}

func NewCreateInteractor(
	repository createRepository,
) *CreateInteractor {
	return &CreateInteractor{
		repository: repository,
	}
}

func (ci *CreateInteractor) Execute(ctx context.Context, input port.SampleCreateInput) (port.SampleCreateOutput, error) {
	rand.Seed(time.Now().UnixNano())

	id := strconv.Itoa(rand.Intn(10))

	model := model.NewSample(sample.ID(id), input.Name)

	if err := ci.repository.Save(ctx, model); err != nil {
		log.Printf("Failed to save sample: %s", err)

		return port.SampleCreateOutput{}, err
	}

	return port.SampleCreateOutput{
		Sample: model,
	}, nil
}
