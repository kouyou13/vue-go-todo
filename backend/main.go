package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Task データ構造の定義
type Task struct {
	ID         string     `json:"id"`
	Created_at time.Time  `json:"createdAt"`
	Title      string     `json:"title"`
	Completed  bool       `json:"completed"`
	Limited_at *time.Time `json:"limitedAt"`
	Note       string     `json:"note"`
	CategoryId *string     `json:"categoryId"`
}

// Category データ構造の定義
type Category struct {
	ID         string     `json:"id"`
	Created_at time.Time  `json:"createdAt"`
	Name       string     `json:"name"`
}

// メモリ上のデータベース（簡易的なストレージ）
var timezone = time.FixedZone("Asia/Tokyo", 9*60*60)
var tasks = [] Task {
	{ID: "e9f9c03e-d36b-470a-8735-c79e3cb848a3", Created_at: time.Date(1900, 1, 1, 0, 0, 0, 0, timezone), Title: "Goの勉強", Completed: false, Limited_at: nil, Note: "テスト", CategoryId: nil},
	{ID: "380153cc-3677-45c2-b122-8744075c9011", Created_at: time.Date(1900, 1, 1, 0, 0, 0, 0, timezone), Title: "Vue.jsのセットアップ", Completed: true, Limited_at: nil, Note: "", CategoryId: nil},
}
var categories = [] Category {
	{ID: "672bb739-506b-4263-8a87-23ccc1645fc7", Created_at: time.Date(1900, 1, 1, 0, 0, 0, 0, timezone), Name: "仕事"},
	{ID: "ec38e055-95a6-4180-ab8e-5ed5142cf533", Created_at: time.Date(1900, 1, 1, 0, 0, 0, 0, timezone), Name: "プライベート"},
}

func main() {
	r := gin.Default()

	// CORS設定 (フロントエンドからのアクセスを許可)
	// 開発中は全てのオリジンを許可する設定にしています。
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	// APIエンドポイントの定義
	api := r.Group("/api")
	{
		api.GET("/tasks", getTasks)
		api.POST("/tasks", createTask)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)

		api.GET("/categories", getCategories)
		api.POST("/categories", createCategory)
		api.PUT("/categories/:id", updateCategory)
		api.DELETE("/categories/:id", deleteCategory)
	}

	// システムが割り当てられるポートに設定
	port := os.Getenv("PORT") // 環境変数PORTを取得
	if port == "" {
			port = "8080" // ローカル用のデフォルト
	}
	r.Run(":" + port)
}

// --- ハンドラー関数 ---

// 全てのタスクを取得
func getTasks(c *gin.Context) {
	searchTitle := c.Query("title")

	if searchTitle == "" {
		c.JSON(http.StatusOK, tasks)
		return
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), strings.ToLower(searchTitle)){
			filteredTasks = append(filteredTasks, task)
		}
	}
	c.JSON(http.StatusOK, filteredTasks)
}

// 新しいタスクを作成
func createTask(c *gin.Context) {
	var newTask Task
	// リクエストボディのJSONを構造体にバインド
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = uuid.NewString()
	newTask.Created_at = time.Now()
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

// タスクの更新 (完了状態のトグルなど)
func updateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			// タイトルと完了状態を更新
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// タスクの削除
func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			// スライスから要素を削除
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// カテゴリー取得
func getCategory(c *gin.Context) {
	id := c.Param("id")

	for _, category := range categories {
		if category.ID == id {
			c.JSON(http.StatusOK, category)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

// カテゴリー全取得
func getCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

// カテゴリー追加
func createCategory(c *gin.Context) {
	var newCategory Category
	// リクエストボディのJSONを構造体にバインド
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCategory.ID = uuid.NewString()
	newCategory.Created_at = time.Now()
	categories = append(categories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}

// カテゴリーの更新
func updateCategory(c *gin.Context) {
	id := c.Param("id")

	var updatedCategory Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, category := range categories {
		if category.ID == id {
			// タイトルと完了状態を更新
			categories[i] = updatedCategory
			c.JSON(http.StatusOK, categories[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

// カテゴリーの削除
func deleteCategory(c *gin.Context) {
	id := c.Param("id")

	for i, category := range categories {
		if category.ID == id {
			// スライスから要素を削除
			categories = append(categories[:i], categories[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}