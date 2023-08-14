package entities

type UserInfo struct {
	User   User    `json:"user"`
	Themes []Theme `json:"themes"`
}
