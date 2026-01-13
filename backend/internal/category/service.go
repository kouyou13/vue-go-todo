package category

import (
	"errors"
	"time"
	"strings"
)

// categories はメモリ上の擬似DB
var categories = []Category{
	{
		ID: "308cb5ae-a1a8-41e7-9527-315a319ad399",
		CreatedAt: time.Date(2026, 1, 1, 10, 0, 0, 0, time.Local),
		Name: "カテゴリー1",
	},
	{
		ID: "4695698c-2dd1-457a-aa90-5e5878980997",
		CreatedAt: time.Date(2026, 1, 2, 10, 0, 0, 0, time.Local),
		Name: "カテゴリー2",
	},
}

// GetAll は全タスクを返す
func GetAll(searchWord string) []Category {
	if searchWord == "" {
		return categories
	}

	var filtered []Category
	for _, t := range categories {
		if strings.Contains(strings.ToLower(t.Name), strings.ToLower(searchWord)) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

// Create はタスクを追加する処理
func Create(t Category) Category {
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
