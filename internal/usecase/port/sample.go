package port

import (
	"github.com/takokun778/oreoreddd/internal/domain/model"
	"github.com/takokun778/oreoreddd/internal/domain/model/sample"
	"github.com/takokun778/oreoreddd/internal/usecase"
)

type SampleCreateInput struct {
	usecase.Input
	Name sample.Name
}

type SampleCreateOutput struct {
	usecase.Output
	Sample model.Sample
}

type SampleCreateUsecase interface {
	usecase.Usecase[SampleCreateInput, SampleCreateOutput]
}

type SampleReadInput struct {
	usecase.Input
	ID sample.ID
}

type SampleReadOutput struct {
	usecase.Output
	Sample model.Sample
}

type SampleReadUsecase interface {
	usecase.Usecase[SampleReadInput, SampleReadOutput]
}

type SampleUpdateInput struct {
	usecase.Input
	ID   sample.ID
	Name sample.Name
}

type SampleUpdateOutput struct {
	usecase.Output
	Sample model.Sample
}

type SampleUpdateUsecase interface {
	usecase.Usecase[SampleUpdateInput, SampleUpdateOutput]
}
