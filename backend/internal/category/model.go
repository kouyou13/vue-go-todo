package category

import "time"

// Category データ構造の定義
type Category struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
