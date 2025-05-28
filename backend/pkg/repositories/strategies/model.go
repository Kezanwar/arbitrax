package strategy_repo

type Model struct {
	UUID        string `json:"uuid" db:"uuid"`
	Label       string `json:"label" db:"label"`
	AvatarURL   string `json:"avatar_url" db:"avatar_url"`
	Description string `json:"description" db:"description"`
}
