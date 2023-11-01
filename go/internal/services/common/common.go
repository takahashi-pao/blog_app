package common

import (
	"path/filepath"

	"github.com/google/uuid"
)

/*
時刻に基づいた一意のファイル名を生成
*/
func GenerateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	uniqueName := uuid.New().String() + ext
	return uniqueName
}
