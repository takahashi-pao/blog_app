package logUtil

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// ログ初期化
// 【引数】
// fileName: ログファイル名
// 【返り値】
// func(logContnt string) → 受け取った文字列をログに吐き出す
// func() → ファイルを閉じる　※deferで呼ぶこと
func LogInit(fileName string) (func(logContnt string), func()) {
	// ログを書き込むファイルを開く（なければ作成）
	filePath := "../Log/" + fileName + ".txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// エラーハンドリング
		log.Fatal(err)
	}

	// ログの出力先をファイルに設定
	log.SetOutput(file)

	// ログのフォーマットを設定（時間を含める）
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	logStartStr := os.Getenv("LOG_START") + fileName + os.Getenv("LOG_CLOSE") + time.Now().String()
	writeLog(logStartStr, file)

	// 関数が終了する際にファイルを閉じる
	return func(logContnt string) { writeLog(logContnt, file) }, func() { endLog(fileName, file) }
}

// ログ排出
func writeLog(logContent string, file *os.File) {
	_, err := fmt.Fprintln(file, logContent)
	if err != nil {
		// ファイルエラーをどうする？
		log.Print("ファイルエラー")
	}
}

// ログ終了
func endLog(fileName string, file *os.File) {
	logStartStr := os.Getenv("LOG_END") + fileName + os.Getenv("LOG_CLOSE") + time.Now().String()
	writeLog(logStartStr, file)
	writeLog("", file)

	err := file.Close()
	if err != nil {
		// ファイルを閉じる際のエラーをどうする？
		writeLog("ファイルを閉じる時にエラー", file)
	}
}
