package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/takokun778/oreoreddd/internal/domain/model/sample"
	"github.com/takokun778/oreoreddd/internal/usecase/port"
)

const SamplePath = "samples"

type Sample struct {
	create port.SampleCreateUsecase
	read   port.SampleReadUsecase
	update port.SampleUpdateUsecase
}

func NewSample(
	create port.SampleCreateUsecase,
	read port.SampleReadUsecase,
	update port.SampleUpdateUsecase,
) *Sample {
	return &Sample{
		create: create,
		read:   read,
		update: update,
	}
}

func (s *Sample) Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		s.get(w, r)
	case http.MethodPost:
		s.post(w, r)
	case http.MethodPut:
		s.put(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Sample) get(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/"+SamplePath)
	_, id := filepath.Split(sub)
	if id == "" {
		log.Printf("Failed to get sample: %s", r.URL.Path)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input := port.SampleReadInput{
		ID: sample.ID(id),
	}

	output, err := s.read.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Failed to read sample: %s", err)

		w.WriteHeader(http.StatusNotFound)

		return
	}

	type Res struct {
		Sample struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"sample"`
	}

	res := &Res{
		Sample: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   output.Sample.ID.String(),
			Name: output.Sample.Name.String(),
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("%+v", res)

	w.Header().Set("Content-Type", "application/json")
}

func (s *Sample) post(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		Name string `json:"name"`
	}

	var req Req

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	input := port.SampleCreateInput{
		Name: sample.Name(req.Name),
	}

	output, err := s.create.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Failed to create sample: %s", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	type Res struct {
		Sample struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"sample"`
	}

	res := &Res{
		Sample: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   output.Sample.ID.String(),
			Name: output.Sample.Name.String(),
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (s *Sample) put(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/"+SamplePath)
	_, id := filepath.Split(sub)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	type Req struct {
		Name string `json:"name"`
	}

	var req Req

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	input := port.SampleUpdateInput{
		ID:   sample.ID(id),
		Name: sample.Name(req.Name),
	}

	output, err := s.update.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Failed to update sample: %s", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	type Res struct {
		Sample struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"sample"`
	}

	res := &Res{
		Sample: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   output.Sample.ID.String(),
			Name: output.Sample.Name.String(),
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}
