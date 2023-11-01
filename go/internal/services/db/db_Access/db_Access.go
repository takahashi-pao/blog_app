package dbAccess

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DBのセットアップ（Open処理)
func setupDB(dbDriver string, dsn string) *sql.DB {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalf("GetArticleData rows.Scan error err:%v", err)
	}

	return db
}

/*
DBインスタンスの生成
*/
func AccessDB() *sql.DB {
	// .env ファイルをロード
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から値を取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dbDriver := "postgres"
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	db := setupDB(dbDriver, dsn)

	return db
}
