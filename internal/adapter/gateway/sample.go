package gateway

import (
	"context"
	"errors"
	"sync"

	"github.com/takokun778/oreoreddd/internal/domain/model"
	"github.com/takokun778/oreoreddd/internal/domain/model/sample"
	"github.com/takokun778/oreoreddd/internal/domain/repository"
)

var _ repository.SampleFinder = (*Sample)(nil)
var _ repository.SampleSaver = (*Sample)(nil)

type Sample struct {
	lock   sync.Mutex
	memory map[sample.ID]model.Sample
}

func NewSample() *Sample {
	return &Sample{
		memory: make(map[sample.ID]model.Sample),
	}
}

func (s *Sample) Find(ctx context.Context, id sample.ID) (model.Sample, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.memory[id]
	if !ok {
		return model.Sample{}, errors.New("")
	}

	return v, nil
}

func (s *Sample) Save(ctx context.Context, sample model.Sample) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.memory[sample.ID] = sample

	return nil
}
