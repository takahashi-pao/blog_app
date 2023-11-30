package article

import (
	"database/sql"
	"log"
	"net/http"
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
	db := dbAccess.AccessDB()
	articleData, err := getArticleData(db)
	if err != nil {
		// エラーハンドリング
		log.Print(err)
		return data
	}

	return articleData
}

func getTagsForArticle(db *sql.DB, articleID int) ([]string, error) {
	var tags []string

	sqlCommand := "SELECT tags.name FROM tags JOIN article_tags ON tags.tag_id = article_tags.tag_id WHERE article_tags.article_id = $1"
	rows, err := db.Query(sqlCommand, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func parseDateTime(dateTime string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return "", err
	}
	return parsedTime.Format("2006/01/02"), nil
}

func getArticleData(db *sql.DB) ([]gin.H, error) {
	var data []gin.H

	sqlCommand := "SELECT id, title, update_date, thumbnail FROM article WHERE delete_flag = false"
	rows, err := db.Query(sqlCommand)
	if err != nil {
		log.Printf("GetArticleData db.Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		m := &article_model.Article_list_model{}

		if err := rows.Scan(&m.ID, &m.Title, &m.DateTime, &m.Thumbnail); err != nil {
			log.Printf("GetArticleData rows.Scan error: %v", err)
			return nil, err
		}

		tags, err := getTagsForArticle(db, m.ID)
		if err != nil {
			log.Printf("getTagsForArticle error: %v", err)
			return nil, err
		}
		m.Tag = tags

		dateTime, err := parseDateTime(m.DateTime)
		if err != nil {
			log.Printf("parseDateTime error: %v", err)
			return nil, err
		}
		m.DateTime = dateTime

		articleData := gin.H{
			"id":        m.ID,
			"title":     m.Title,
			"date":      m.DateTime,
			"tag":       m.Tag,
			"thumbnail": m.Thumbnail,
		}

		data = append(data, articleData)
	}

	if err := rows.Err(); err != nil {
		log.Printf("GetArticleData rows.Err error: %v", err)
		return nil, err
	}

	return data, nil
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
