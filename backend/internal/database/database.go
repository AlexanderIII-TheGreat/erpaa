package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


// func DataCon() *sql.DB{
// 	err := godotenv.Load("C:/Users/user/Documents/ERPAA/backend/.env")
// 	if err != nil {
// 		panic(err)
// 	}

// 	HOST := os.Getenv("DB_HOST")
// 	USER := os.Getenv("DB_USERNAME")
// 	PASS := os.Getenv("DB_PASSWORD")
// 	PATH := os.Getenv("DB_PATH")
// 	CONNEC := os.Getenv("DB_CONNEC")
// 	DB := os.Getenv("DB_NAME")

// 	script := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", USER, PASS, HOST, PATH, CONNEC)

// 	result, err := sql.Open(DB,script)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result


// }


func DataCon() *sql.DB {
	err := godotenv.Load("/var/www/erpaa/app/backend/.env")
	if err != nil {
		panic("Gagal load file .env")
	}

	host := os.Getenv("DB_H")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	PATH := os.Getenv("DB_PATH")
	dbname := os.Getenv("DB_CONNEC")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, PATH, dbname,
	)

	fmt.Println("DB HOST:", host)
	fmt.Println("DB USER:", user)
	fmt.Println("DB NAME:", dbname)
	fmt.Println("DB PATH:", PATH)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 🔥 INI KUNCI DEBUGGING
	err = db.Ping()
	if err != nil {
		panic("DATABASE TIDAK TERKONEKSI: " + err.Error())
	}

	fmt.Println("✅ DATABASE TERHUBUNG DENGAN SUKSES")

	return db
}
