package article

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"example.com/blog_app/go/internal/handlers/auth"
	article_model "example.com/blog_app/go/internal/models/article"
	timeformat "example.com/blog_app/go/internal/models/timeFormat"
	common "example.com/blog_app/go/internal/services/common"
	dbAccess "example.com/blog_app/go/internal/services/db/db_Access"
	"github.com/gin-gonic/gin"
)

/*
記事の取得
*/
func GetArticleData() []gin.H {
	var data []gin.H

	sqlCommand := "select id, title,update_date ,tag ,thumbnail  from article where delete_flag = false"

	db := dbAccess.AccessDB()
	rows, err := db.Query(sqlCommand)

	if err != nil {
		log.Printf("GetArticleData db.Query error err:%v", err)
		return data
	}

	defer rows.Close()

	for rows.Next() {
		m := &article_model.Article_list_model{}
		var tag_str string

		if err := rows.Scan(&m.ID, &m.Title, &m.DateTime, &tag_str, &m.Thumbnail); err != nil {
			log.Printf("GetArticleData rows.Scan error err:%v", err)
			return data
		}

		// time.RFC3339のフォーマットにパース
		parsedTime, err := time.Parse(time.RFC3339, m.DateTime)
		if err != nil {
			fmt.Println("エラー:", err)
			return data
		}

		m.DateTime = parsedTime.Format("2006/01/02")

		m.Tag = strings.Split(tag_str, ",")

		articleData := gin.H{
			"id":        m.ID,
			"title":     m.Title,
			"date":      m.DateTime,
			"tag":       m.Tag,
			"thumbnail": m.Thumbnail,
		}

		data = append(data, articleData)
	}
	return data
}

func DeleteArticle(c *gin.Context) {
	userId := auth.GetLoginUserId()

	ids_str := c.PostForm("ids")
	if ids_str == "" {
		log.Printf("ERROR:削除対象のIDが取得できませんでした")
		c.JSON(http.StatusOK, gin.H{"error": "削除対象のIDが取得できませんでした"})
		return
	}

	ids, err := common.ParseAndConvert(ids_str)
	if err != nil {
		log.Printf("ERROR:削除対象のIDを配列に変換できませんでした")
		c.JSON(http.StatusOK, gin.H{"error": "削除対象のIDを配列に変換できませんでした"})
		return
	}

	// ids_str := "("
	where_str := ""

	for i := 0; i < len(ids); i++ {
		if i == 0 {
			where_str += "("
		} else {
			where_str += ","
		}

		where_str += "'" + common.ConvertToString(ids[i]) + "'"

		if i == len(ids)-1 {
			where_str += ")"
		}
	}

	update_sql := "update article set delete_flag = $1, delete_user =$2, delete_date =$3 where id in " + where_str
	db := dbAccess.AccessDB()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(update_sql, true, userId, timeformat.GetTimeNowMoldFormat())

	if err != nil {
		log.Fatalf("記事の削除失敗 db.Exec error err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "記事の削除に失敗しました"})
		tx.Rollback()
		return
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "削除が完了しました"})
}
