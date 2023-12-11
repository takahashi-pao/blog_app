package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogInit() {
	// ログを書き込むファイルを開く（なければ作成）
	file, err := os.OpenFile("testlogfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// エラーハンドリング
		log.Fatal(err)
	}

	// 関数が終了する際にファイルを閉じる
	defer file.Close()

	// ログの出力先をファイルに設定
	log.SetOutput(file)

	// ログのフォーマットを設定（時間を含める）
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
