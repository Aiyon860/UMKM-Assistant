package service

import (
	"strings"
)


type ChatbotService struct{}

func NewChatbotService() *ChatbotService {
	return &ChatbotService{}
}

func (s *ChatbotService) Chat(message string) string {
	msg := strings.ToLower(message)

	switch {
	case strings.Contains(msg, "produk"):
		return "Kami punya Sofa, Meja, dan Lemari. Anda ingin lihat yang mana?"
	case strings.Contains(msg, "sofa"):
		return "Sofa kami tersedia dalam berbagai warna dan ukuran."
	case strings.Contains(msg, "harga"):
		return "Harga produk mulai dari Rp1.000.000."
	default:
		return "Maaf, saya belum paham maksud Anda. Bisa dijelaskan lagi?"
	}
}
