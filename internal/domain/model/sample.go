package model

import "github.com/takokun778/oreoreddd/internal/domain/model/sample"

type Sample struct {
	ID   sample.ID
	Name sample.Name
}

func NewSample(
	id sample.ID,
	name sample.Name,
) Sample {
	return Sample{
		ID:   id,
		Name: name,
	}
}

func (s Sample) Update(
	name sample.Name,
) Sample {
	return Sample{
		ID:   s.ID,
		Name: name,
	}
}
