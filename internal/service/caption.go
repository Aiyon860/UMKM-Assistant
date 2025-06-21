package service

import "todo-go/internal/model"

type CaptionService interface {
	GenerateCaption(req model.CaptionRequest) (string, error)
}

type mockCaptionService struct{}

func NewMockCaptionService() CaptionService {
	return &mockCaptionService{}
}

func (s *mockCaptionService) GenerateCaption(req model.CaptionRequest) (string, error) {
	// Caption dummy/mock
	caption := "Promo Spesial! Dapatkan Kopi Gayo 100% Asli hanya Rp25.000 sekarang juga! ☕🔥"
	return caption, nil
}
