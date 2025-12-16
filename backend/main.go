package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Task データ構造の定義
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// メモリ上のデータベース（簡易的なストレージ）
var tasks = [] Task {
	{ID: 1, Title: "Goの勉強", Completed: false},
	{ID: 2, Title: "Vue.jsのセットアップ", Completed: true},
}
var nextID = 3

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

	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

// タスクの更新 (完了状態のトグルなど)
func updateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

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
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// タスクの削除
func deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

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