package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"example.com/blog_app/go/internal/handlers/article"
	"example.com/blog_app/go/internal/handlers/auth"
	"example.com/blog_app/go/internal/handlers/register"
	"example.com/blog_app/go/internal/handlers/thumbnail"
	"example.com/blog_app/go/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func main() {
	// Ginのルーターを作成
	router := gin.Default()
	// リクエスト元ポートの設定
	router.Use(corsMiddleware())

	// セッションの設定
	sessionInit()
	authGroup := router.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/Register", register.Register)
	}

	/* API　*****************************************************/

	// 記事情報の取得
	router.GET("/GetArticle", func(c *gin.Context) {
		c.JSON(http.StatusOK, article.GetArticleData())
	})

	// ログイン状態の取得
	router.GET("/IsLogin", auth.IsLogin)

	// 記事サムネイル取得
	router.GET("/GetThumbnail/:path", thumbnail.GetThumbnail)

	// 重複IDのチェック
	router.GET("/CheckExistId/:id", auth.CheckExistId)

	router.GET("/SignOut", auth.SignOut)
	/**********************************************************/

	/* POST *****************************************************/

	router.POST("/SignIn", auth.SignIn)

	router.POST("/SignUp", auth.SignUp)

	/**********************************************************/

	// サーバーを起動
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
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

// セッション用の初期処理
func sessionInit() {

	// 乱数生成
	b := make([]byte, 48)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	str := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	// 新しいstoreとセッションを準備
	auth.Store = sessions.NewCookieStore([]byte(str))
	auth.Session = sessions.NewSession(auth.Store, auth.Session_name)

	// セッションの有効範囲を指定
	auth.Store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 24 * 7,
		Secure:   false,
		HttpOnly: true,
	}

	// log
	fmt.Println("key     data --")
	fmt.Println(str)
	fmt.Println("")
	fmt.Println("store   data --")
	fmt.Println(auth.Store)
	fmt.Println("")
	fmt.Println("session data --")
	fmt.Println(auth.Session)
	fmt.Println("")

}
