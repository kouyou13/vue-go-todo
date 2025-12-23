package task

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// tasks はメモリ上の擬似DB
var tasks = []Task{}

// GetAll は全タスクを返す
func GetAll(searchTitle string) []Task {
	if searchTitle == "" {
		return tasks
	}

	var filtered []Task
	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), strings.ToLower(searchTitle)) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

// Create はタスクを追加する処理
func Create(t Task) Task {
	t.ID = uuid.NewString()
	t.CreatedAt = time.Now()
	tasks = append(tasks, t)
	return t
}

// Update はID一致するタスクを更新する
func Update(id string, updated Task) (Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			updated.CreatedAt = t.CreatedAt
			tasks[i] = updated
			return updated, nil
		}
	}
	return Task{}, errors.New("task not found")
}

// Delete はタスクを削除する
func Delete(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
