package secretkey

import (
	"crypto/rand"
	"encoding/base64"
)

// 秘密鍵の生成
func GenerateSecretKey() (string, error) {
	// 32バイトのランダムなデータを生成
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// ランダムデータをBase64エンコードして秘密鍵として使用
	secretKey := base64.URLEncoding.EncodeToString(randomBytes)
	return secretKey, nil
}
