package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Task データ構造の定義
type Task struct {
	ID         string    `json:"id"`
	Created_at time.Time `json:"createdAt"`
	Title      string    `json:"title"`
	Completed  bool      `json:"completed"`
	Limited_at *time.Time `json:"limitedAt"`
}

// メモリ上のデータベース（簡易的なストレージ）
var timezone = time.FixedZone("Asia/Tokyo", 9*60*60)
var tasks = [] Task {
	{ID: "e9f9c03e-d36b-470a-8735-c79e3cb848a3", Created_at: time.Date(1900, 1, 1, 0, 0, 0, 0, timezone), Title: "Goの勉強", Completed: false, Limited_at: nil},
	{ID: "380153cc-3677-45c2-b122-8744075c9011", Created_at: time.Date(1900, 1, 1, 0, 0, 0, 0, timezone), Title: "Vue.jsのセットアップ", Completed: true, Limited_at: nil},
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
	}

	// サーバーをポート8080で起動
	r.Run(":8080")
}

// --- ハンドラー関数 ---

// 全てのタスクを取得
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
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
			tasks[i].Title = updatedTask.Title
			tasks[i].Completed = updatedTask.Completed
			tasks[i].Limited_at = updatedTask.Limited_at
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