package category

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// categories はメモリ上の擬似DB
var categories = []Category{}

// GetAll は全タスクを返す
func GetAll() []Category {
	return categories
}

// Create はタスクを追加する処理
func Create(t Category) Category {
	t.ID = uuid.NewString()
	t.CreatedAt = time.Now()
	categories = append(categories, t)
	return t
}

// Update はID一致するタスクを更新する
func Update(id string, updated Category) (Category, error) {
	for i, t := range categories {
		if t.ID == id {
			updated.CreatedAt = t.CreatedAt
			categories[i] = updated
			return updated, nil
		}
	}
	return Category{}, errors.New("Category not found")
}

// Delete はタスクを削除する
func Delete(id string) error {
	for i, t := range categories {
		if t.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}
	return errors.New("Category not found")
}
