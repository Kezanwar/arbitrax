package strategy_repo

type Model struct {
	ID          int    `json:"-" db:"id"`
	Key         string `json:"key" db:"key"`
	Label       string `json:"label" db:"label"`
	Description string `json:"description" db:"description"`
}
