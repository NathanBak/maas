package meme

import (
	"errors"

	"github.com/google/uuid"
)

type Meme struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text,omitempty"`
	URL  string    `json:"url,omitempty"`
}

func NewTextMeme(text string) (Meme, error) {
	if text == "" {
		return Meme{}, errors.New("text cannot be empty for text meme")
	}
	return Meme{
		ID:   uuid.New(),
		Text: text,
	}, nil
}

func NewURLMeme(url string) (Meme, error) {
	if url == "" {
		return Meme{}, errors.New("url cannot be empty for url meme")
	}
	return Meme{
		ID:  uuid.New(),
		URL: url,
	}, nil
}
