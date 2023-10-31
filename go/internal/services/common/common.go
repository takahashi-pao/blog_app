package common

import (
	"path/filepath"

	"github.com/google/uuid"
)

func GenerateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	uniqueName := uuid.New().String() + ext
	return uniqueName
}
