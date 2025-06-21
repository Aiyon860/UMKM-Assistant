package repository

import "todo-go/internal/model"

type CaptionRepository interface {
	LogMockRequest(data model.CaptionRequest) error
}

type mockCaptionRepo struct{}

func NewMockCaptionRepository() CaptionRepository {
	return &mockCaptionRepo{}
}

func (r *mockCaptionRepo) LogMockRequest(data model.CaptionRequest) error {
	// Tidak melakukan apapun
	return nil
}