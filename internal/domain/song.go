package domain

import "time"

type Song struct {
	Id            int       `json:"id"`
	Name          string    `json:"song_name"`
	PerformerName string    `json:"performer_name"`
	Link          string    `json:"link,omitempty"`
	Text          string    `json:"song_text,omitempty"`
	ReleaseDate   time.Time `json:"release_date,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}
type UpdateSong struct {
	Name          string `json:"song_name"`
	PerformerName string `json:"performer_name"`
	Link          string `json:"link,omitempty"`
	Text          string `json:"song_text,omitempty"`
}
type FiltersForSong struct {
	Name          *string `json:"song_name,omitempty"`
	PerformerName *string `json:"performer_name,omitempty"`
	Link          *string `json:"link,omitempty"`
	Text          *string `json:"song_text,omitempty"`
	ReleaseDate   *string `json:"release_date,omitempty"`
	Limit         *int32  `json:"limit,omitempty"`
	Offset        *int32  `json:"offset,omitempty"`
}
