package main

import (
	"os"

	"github.com/kouyou13/vue-go-todo/backend/internal/router"
)

func main() {
	r := router.SetupRouter()

	// システムが割り当てられるポートに設定
	port := os.Getenv("PORT") // 環境変数PORTを取得
	if port == "" {
		port = "8080" // ローカル用のデフォルト
	}
	r.Run(":" + port)
}
