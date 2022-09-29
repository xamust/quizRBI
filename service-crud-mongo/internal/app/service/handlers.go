package service

import (
	"encoding/json"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/model"
	"io"
	"net/http"
)

func (s *Service) requestHandler(w http.ResponseWriter, req *http.Request) {
	var err error

	w.Header().Set("Content-Type", "application/json")
	response := []model.Data{}
	data := model.Data{}
	if err = json.NewDecoder(req.Body).Decode(&data); err != nil {
		if req.Method != http.MethodGet && err != io.EOF {
			s.logger.Errorf("handlers error: %v", err)
		}
	}

	switch req.Method {
	case "POST":
		response, err = s.repository.CreateRecord(s.ctx, data)
	case "GET":
		response, err = s.repository.GetRecords(s.ctx)
	case "PUT":
		response, err = s.repository.UpdateRecord(s.ctx, data)
	case "DELETE":
		response, err = s.repository.DeleteRecord(s.ctx, data)
	}

	if err != nil {
		response = []model.Data{
			{
				Error: err.Error(),
			},
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err = enc.Encode(response); err != nil {
		s.logger.Errorf("handlers error: %v", err)
	}
}
