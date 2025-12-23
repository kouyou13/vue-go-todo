package task

import "time"

// Task はタスク情報の定義にゃん
type Task struct {
	ID         string     `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	Title      string     `json:"title"`
	Completed  bool       `json:"completed"`
	LimitedAt  *time.Time `json:"limitedAt,omitempty"`
	Note       string     `json:"note"`
	CategoryID *string    `json:"categoryId,omitempty"`
}
