package common

import (
	"encoding/json"
	"log"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/harakeishi/gats"
)

/*
時刻に基づいた一意のファイル名を生成
*/
func GenerateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	uniqueName := uuid.New().String() + ext
	return uniqueName
}

// JSON形式の文字列を[]interface{}にパース
func ParseAndConvert(input string) ([]interface{}, error) {
	var values []interface{}
	err := json.Unmarshal([]byte(input), &values)
	if err != nil {
		return nil, err
	}

	return values, nil
}

// interface{}型の値を文字列に変換
func ConvertToString(value interface{}) string {
	strValue, _ := gats.ToString(value)
	return strValue
}

// メソッド開始・終了ログ
func StartEndLog(methodName string) func() {
	startTime := time.Now()
	log.Printf("【START】%v-------------------------", methodName)

	return func() {
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Printf("【END: %v】%v-------------------------", elapsed, methodName)
	}
}
