package register

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	timeformat "example.com/blog_app/go/internal/models/timeFormat"
	dbAccess "example.com/blog_app/go/internal/services/db_Access"
)

func Register(c *gin.Context) {
	// フォームからファイルを取得
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ファイルが送信されていません"})
		return
	}

	insertQuery := "INSERT INTO article (title, create_user, create_date, update_user, update_date, delete_user, delete_date, tag, thumbnail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	db := dbAccess.AccessDB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insertQuery, c.PostForm("title"), "USER", timeformat.GetTimeNowMoldFormat(), "USER", timeformat.GetTimeNowMoldFormat(), nil, nil, c.PostForm("tag"), file.Filename)

	if err != nil {
		log.Fatalf("登録失敗 db.Exec error err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録に失敗しました"})
		return
	}

	// ファイルを指定のディレクトリに保存
	destinationPath := "img/" + file.Filename
	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの保存中にエラーが発生しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登録が完了しました", "path": destinationPath})
}
