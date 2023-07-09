package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var members = []member{
	{
		ID:   1,
		Name: "test1",
		Icon: "🐶",
	},
	{
		ID:   2,
		Name: "test2",
		Icon: "🐷",
	},
	{
		ID:   3,
		Name: "test3",
		Icon: "🐙",
	},
}

func main() {
	// Ginのルーターを作成
	router := gin.Default()

	router.Use(corsMiddleware())

	// GETリクエストに対するハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, getMembers())
	})

	// POSTリクエストに対するハンドラの定義
	router.POST("/api/posts", func(c *gin.Context) {
		// リクエストからデータを受け取る
		var post struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// データベースに投稿を保存する処理などを実行

		// レスポンスを返す
		c.JSON(http.StatusOK, gin.H{
			"message": "Post created successfully",
			"post":    post,
		})
	})

	// サーバーを起動
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func getMembers() []gin.H {
	var data []gin.H

	for _, m := range members {
		memberData := gin.H{
			"id":   m.ID,
			"name": m.Name,
			"icon": m.Icon,
		}
		data = append(data, memberData)
	}

	return data
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエスト元のオリジンを取得
		origin := c.Request.Header.Get("Origin")

		// ポート番号が3000であるオリジンのみ許可
		if origin == "http://localhost:3000" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// 他の必要なCORSヘッダーを設定
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// OPTIONSリクエストに対応するために、OPTIONSメソッドの場合は早期に終了する
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// 次のミドルウェアやハンドラを実行
		c.Next()
	}
}
