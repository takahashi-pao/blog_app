package article

import (
	"fmt"
	"log"
	"strings"
	"time"

	article_model "example.com/blog_app/go/internal/models/article"
	dbAccess "example.com/blog_app/go/internal/services/db_Access"
	"github.com/gin-gonic/gin"
)

/*
記事の取得
*/
func GetArticleData() []gin.H {
	var data []gin.H

	sqlCommand := "select id, title,update_date ,tag ,thumbnail  from article"

	db := dbAccess.AccessDB()
	rows, err := db.Query(sqlCommand)

	if err != nil {
		log.Fatalf("GetArticleData db.Query error err:%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		m := &article_model.Article_list_model{}
		var tag_str string

		if err := rows.Scan(&m.ID, &m.Title, &m.DateTime, &tag_str, &m.Thumbnail); err != nil {
			log.Fatalf("GetArticleData rows.Scan error err:%v", err)
		}

		// time.RFC3339のフォーマットにパース
		parsedTime, err := time.Parse(time.RFC3339, m.DateTime)
		if err != nil {
			fmt.Println("エラー:", err)
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
