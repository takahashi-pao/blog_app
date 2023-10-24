package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	timeformat "example.com/blog_app/go/internal/models/timeFormat"
	dbAccess "example.com/blog_app/go/internal/services/db_Access"
)

// セッションストアの初期化
var store = sessions.NewCookieStore([]byte("your-secret-key"))

// サインイン
func SignIn(c *gin.Context) {
	/* ログイン成功フラグ */
	isLoginSuccessFlg := false

	/* 認証確認のSELECT文 */
	selectQuery := `select 
										case 
											when count(*) > 0 then true
											else false
										end as isLoginSuccessFlg	
									from 
										"user_info" ui 
									where id = $1 and password = $2 and delete_date is null`

	db := dbAccess.AccessDB()
	err := db.QueryRow(selectQuery, c.PostForm("id"), c.PostForm("password")).Scan(&isLoginSuccessFlg)

	// SELECT文エラー
	if err != nil {
		log.Fatalf("SignIn db.Query error err:%v", err)
	}

	if isLoginSuccessFlg { // ログイン成功の場合
		// セッションを開始
		session, _ := store.Get(c.Request, "session-name")

		// セッションにセッションIDを追加
		session.Values["userID"] = c.PostForm("id")

		// セッションを保存
		session.Save(c.Request, c.Writer)

		c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
	} else { // ログイン失敗の場合
		// エラーメッセージを返却
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
	}
}

// サインアップ
func SignUp(c *gin.Context) {
	/* INSERT文 */
	insertQuery := "INSERT INTO user_info (id, password, create_date, update_date, delete_date) VALUES ($1, $2, $3, $4, $5)"

	db := dbAccess.AccessDB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insertQuery, c.PostForm("id"), c.PostForm("password"), timeformat.GetTimeNowMoldFormat(), timeformat.GetTimeNowMoldFormat(), nil)

	if err != nil {
		log.Fatalf("登録失敗 db.Exec error err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登録されました"})
	// home画面へリダイレクト
	c.Redirect(http.StatusFound, "/")
}
