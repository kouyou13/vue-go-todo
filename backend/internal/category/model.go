package category

import (
	"time"
)

// Category データ構造の定義
type Category struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Color     Color    `json:"color"`
}

type Color string

const (
	Blue = "blue"
	Red = "red"
	Orange = "orange"
	Yellow = "yellow"
	Green = "green"
	Purple = "purple"
)
