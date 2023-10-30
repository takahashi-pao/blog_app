package thumbnail

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetThumbnail(c *gin.Context) {
	// URLパラメータから引数を取得
	argument := c.Param("path")

	// 引数を使ってサムネイル画像データを読み取り
	imageData, err := os.ReadFile("img/" + argument)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "画像が見つかりません"})
		return
	}

	// 画像データをレスポンスとして返す
	c.Data(http.StatusOK, "image/png", imageData)
}
