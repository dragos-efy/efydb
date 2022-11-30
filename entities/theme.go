package entities

import "gorm.io/gorm"

type Theme struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Config      string   `json:"config"`
	Screenshot  string   `json:"screenshot"`
	Uploaded    uint     `json:"uploaded"`
	EfyVersion  uint     `json:"efy_version"`
	Likes       uint     `json:"likes"`
	Approved    bool     `json:"approved"`
}
