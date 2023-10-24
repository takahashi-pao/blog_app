package timeformat

import "time"

/* 現在時刻(フォーマット整形済)を取得 */
func GetTimeNowMoldFormat() string {
	return time.Now().Format("2006-01-02 15:04:05.111")
}
