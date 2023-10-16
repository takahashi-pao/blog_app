package article

import (
	"log"
	"strings"
	"time"

	article_model "github.com/blog_app/internal/models/article"
	dbAccess "github.com/blog_app/internal/services/db_Access"
	"github.com/gin-gonic/gin"
)

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

var members = []article_model.Article_list_model{
	{
		ID:       1,
		Title:    "test1",
		DateTime: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local).Format("2006/01/02"),
		Tag:      []string{"💐", "☘️", "🌛"},
	},
	{
		ID:       2,
		Title:    "test2",
		DateTime: time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local).Format("2006/01/02"),
		Tag:      []string{"🐀", "🌳", "💆🏻"},
	},
	{
		ID:       3,
		Title:    "test3",
		DateTime: time.Date(2023, 7, 1, 0, 0, 0, 0, time.Local).Format("2006/01/02"),
		Tag:      []string{"🤠", "🥴", "🐙"},
	},
	{
		ID:       4,
		Title:    "test4",
		DateTime: time.Date(2021, 7, 1, 0, 0, 0, 0, time.Local).Format("2006/01/02"),
		Tag:      []string{"🐷", "ごちポ", "レモン"},
	},
}
