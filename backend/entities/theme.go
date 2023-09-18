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
	Approved    bool   `json:"approved"`
	Score       int    `json:"score" gorm:"-"`
	UserScore   int    `json:"user_score" gorm:"-"`
}
