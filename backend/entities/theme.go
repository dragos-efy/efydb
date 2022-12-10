package entities

import "gorm.io/gorm"

type Theme struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Username    string `json:"username"`
	Config      string `json:"config"`
	ImageConfig string `json:"imageConfig"`
	Screenshot  string `json:"screenshot"`
	Uploaded    int64  `json:"uploaded"`
	EfyVersion  uint   `json:"efy_version"`
	Likes       uint   `json:"likes"`
	Approved    bool   `json:"approved"`
}
