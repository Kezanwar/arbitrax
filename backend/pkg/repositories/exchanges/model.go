package exchanges_repo

type Model struct {
	UUID    string `json:"uuid" db:"uuid"`
	Label   string `json:"label" db:"label"`
	LogoURL string `json:"logo_url" db:"logo_url"`
}
