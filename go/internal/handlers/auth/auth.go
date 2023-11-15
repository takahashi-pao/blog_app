package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	auth_response_model "example.com/blog_app/go/internal/models/auth"
	timeformat "example.com/blog_app/go/internal/models/timeFormat"
	"example.com/blog_app/go/internal/services/common"
	dbAccess "example.com/blog_app/go/internal/services/db/db_Access"
)

// セッション名
var Session_name string = "auth_info"

// Cookie型のstore情報
var Store *sessions.CookieStore

// セッションオブジェクト
var Session *sessions.Session

// セッションキー：ユーザーID
const session_userId = "userId"

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
		// セッションにユーザーIDを保存
		SetSession(c.PostForm("id"), session_userId, c)

		c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
	} else { // ログイン失敗の場合
		// エラーメッセージを返却
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
	}
}

// サインアップ
func SignUp(c *gin.Context) {
	/* INSERT文 */
	insertQuery := "INSERT INTO user_info (id, password, create_date, update_date, delete_date, admin_flag) VALUES ($1, $2, $3, $4, $5, false)"

	db := dbAccess.AccessDB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insertQuery, c.PostForm("id"), c.PostForm("password"), timeformat.GetTimeNowMoldFormat(), timeformat.GetTimeNowMoldFormat(), nil)

	if err != nil {
		log.Fatalf("登録失敗 db.Exec error err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録に失敗しました"})
		tx.Rollback()
		return
	}
	tx.Commit()

	// セッションにユーザーIDを保存
	SetSession(c.PostForm("id"), session_userId, c)

	c.JSON(http.StatusOK, gin.H{"message": "登録されました"})

}

/*
サインアウト
*/
func SignOut(c *gin.Context) {
	delete(Session.Values, session_userId)

	err := Session.Save(c.Request, c.Writer)
	if err != nil {
		signOutResponse := auth_response_model.Auth_Response_Props{
			Message:  "",
			Error:    "サインアウトに失敗",
			UserId:   "",
			IsSignIn: true,
		}
		log.Printf("SignOut error err:%v", err)
		c.JSON(http.StatusConflict, signOutResponse)
		return
	}

	signOutResponse := auth_response_model.Auth_Response_Props{
		Message:  "サインアウトしました",
		Error:    "",
		UserId:   "",
		IsSignIn: false,
	}

	c.JSON(http.StatusOK, signOutResponse)
}

/*
IDの重複チェック
*/
func CheckExistId(c *gin.Context) {
	uniqueIdFlag := true

	selectQuery := `select 
										case 
											when count(*) > 0 then false
											else true
										end as uniqueIdFlag	
									from 
										"user_info" ui 
									where id = "$1" and delete_date is null`

	db := dbAccess.AccessDB()
	err := db.QueryRow(selectQuery, c.Param("id")).Scan(&uniqueIdFlag)

	// SELECT文エラー
	if err != nil {
		log.Printf("SignIn db.Query error err:%v", err)
		c.JSON(http.StatusConflict, gin.H{"error": "データベースとの接続に失敗しました"})
		return
	}

	if uniqueIdFlag { // 重複するIDが存在しない場合
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	} else { // 重複するIDが存在する場合
		// エラーメッセージを返却
		c.JSON(http.StatusOK, gin.H{"error": "このIDは既に利用されています"})
	}

}

/*
ログイン状態のチェック
*/
func IsLogin(c *gin.Context) {
	// セッションを取得
	userId := common.ConvertToString(Session.Values["userId"])

	if userId != "nil" {
		signOutResponse := auth_response_model.Auth_Response_Props{
			Message:  "サインインしています",
			Error:    "",
			UserId:   userId,
			IsSignIn: true,
		}
		c.JSON(http.StatusOK, signOutResponse)
		return
	}

	signOutResponse := auth_response_model.Auth_Response_Props{
		Message:  "サインアウトしています",
		Error:    "",
		UserId:   "",
		IsSignIn: false,
	}

	// セッション情報をAPIレスポンスとして返す
	c.JSON(http.StatusOK, signOutResponse)
}

/*
セッションへ値とキーをセット
*/
func SetSession(val interface{}, key string, c *gin.Context) {
	Session.Values[key] = val
	log.Printf(("set val in session(%v):%v"), key, key)
	sessions.Save(c.Request, c.Writer)
}

// ログインユーザーIDを取得
func GetLoginUserId() string {
	return common.ConvertToString(Session.Values["userId"])
}
