package model

type CaptionRequest struct {
	ProductName string
	Price       string
	Tone        string
	ImagePath   string
}

type CaptionResponse struct {
	Success bool   `json:"success"`
	Caption string `json:"caption,omitempty"`
	Message string `json:"message,omitempty"`
}