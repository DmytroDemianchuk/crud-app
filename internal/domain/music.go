package domain

import "errors"

var (
	ErrorMusicNotFound = errors.New("music not found")
)

type Music struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Artist       string `json:"artist"`
	Album        string `json:"album"`
	Genre        string `json:"genre"`
	ReleasedYear int    `json:"released_year"`
}

type UpdateMusicInput struct {
	Name         *string `json:"name"`
	Artist       *string `json:"artist"`
	Album        *string `json:"album"`
	Genre        *string `json:"genre"`
	ReleasedYear *int    `json:"released_year"`
}
