package main

import (
	"log"
	"net/http"
	"sync"

	"example.com/blog_app/go/internal/handlers/article"
	"example.com/blog_app/go/internal/handlers/auth"
	"example.com/blog_app/go/internal/handlers/register"
	"example.com/blog_app/go/internal/handlers/thumbnail"
	"example.com/blog_app/go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのルーターを作成
	router := gin.Default()
	// リクエスト元ポートの設定
	router.Use(middleware.CorsMiddleware())

	// セッションの設定
	middleware.SessionInit()

	// 認証済みのみリクエスト可能
	authGroup := router.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware())
	{
		// 記事登録
		authGroup.POST("/Register", register.Register)
	}

	/* API　*****************************************************/

	// 記事情報の取得
	router.GET("/GetArticle", func(c *gin.Context) {
		var wg sync.WaitGroup

		// 非同期に記事情報を取得
		wg.Add(1)
		go func() {
			defer wg.Done()
			articleData := article.GetArticleData()
			// 取得したデータをJSONで返す
			c.JSON(http.StatusOK, articleData)
		}()

		// ゴルーチンが完了するまで待つ
		wg.Wait()
	})

	// ログイン状態の取得
	router.GET("/IsLogin", auth.IsLogin)

	// 記事サムネイル取得
	router.GET("/GetThumbnail/:path", thumbnail.GetThumbnail)

	// 重複IDのチェック
	router.GET("/CheckExistId/:id", auth.CheckExistId)

	// サインアウト
	router.GET("/SignOut", auth.SignOut)
	/**********************************************************/

	/* POST *****************************************************/
	// サインイン
	router.POST("/SignIn", auth.SignIn)
	// サインアップ
	router.POST("/SignUp", auth.SignUp)
	// 記事削除
	router.POST("/deleteArticle", article.DeleteArticle)

	/**********************************************************/

	// サーバーを起動
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
