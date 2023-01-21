package repository

import (
	"context"

	"github.com/takokun778/oreoreddd/internal/domain/model"
	"github.com/takokun778/oreoreddd/internal/domain/model/sample"
)

type SampleFinder interface {
	Find(context.Context, sample.ID) (model.Sample, error)
}

type SampleSaver interface {
	Save(context.Context, model.Sample) error
}
