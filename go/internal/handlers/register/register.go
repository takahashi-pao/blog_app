package register

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	timeformat "example.com/blog_app/go/internal/models/timeFormat"
	"example.com/blog_app/go/internal/services/common"
	dbAccess "example.com/blog_app/go/internal/services/db/db_Access"
)

func Register(c *gin.Context) {
	// フォームからファイルを取得
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ファイルが送信されていません"})
		return
	}

	// ファイルの拡張子を取得
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// 許可された拡張子のみを受け入れ
	allowedExts := []string{".jpg", ".jpeg", ".png"}
	allowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			allowed = true
			break
		}
	}

	if !allowed {
		log.Println("不正な拡張子を含むファイルです")
		return
	}

	// ファイル名を重複しない文字列に変換
	file.Filename = common.GenerateUniqueFileName(file.Filename)

	// ファイル名の重複は許さない
	isExistsFileName := true

	selectQuery := `select 
										case 
											when count(*) > 0 then true
											else false
										end as isExistsFileName	
									from 
										"article" ui 
									where thumbnail = $1 and delete_date is null`

	db := dbAccess.AccessDB()
	err = db.QueryRow(selectQuery, file.Filename).Scan(&isExistsFileName)
	if err != nil {
		log.Println(err)
		return
	}
	if isExistsFileName {
		log.Println("ファイルの登録に失敗")
		return
	}

	// ファイルの登録処理
	insertQuery := "INSERT INTO article (title, create_user, create_date, update_user, update_date, delete_user, delete_date, tag, thumbnail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	_, err = tx.Exec(insertQuery, c.PostForm("title"), "USER", timeformat.GetTimeNowMoldFormat(), "USER", timeformat.GetTimeNowMoldFormat(), nil, nil, c.PostForm("tag"), file.Filename)

	if err != nil {
		log.Fatalf("登録失敗 db.Exec error err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録に失敗しました"})
		tx.Rollback()
		return
	}

	// ファイルを指定のディレクトリに保存
	destinationPath := "img/" + file.Filename
	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの保存中にエラーが発生しました"})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "登録が完了しました", "path": destinationPath})
}
