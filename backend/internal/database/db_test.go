package database

import (
	"context"
	"erpaa-backend/internal/model"
	"erpaa-backend/internal/repository"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestDb(t *testing.T){
	db := DataCon()
	err := godotenv.Load("C:/Users/user/Documents/ERPAA/backend/.env")
	if err != nil {
		panic(err)
	}
	TABLE := os.Getenv("DB_USER")
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	script := fmt.Sprintf(`CREATE TABLE %v (			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			qty INT,
			harga INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)`, TABLE)

	result, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("selamat anda berhasil membuat table")
}


func TestCreateUserTable(t *testing.T) {
	err := godotenv.Load("C:/Users/user/Documents/ERPAA/backend/.env")
	if err != nil {
		panic(err)
	}

	USERTABLE := os.Getenv("DB_USER")
	db := DataCon()

	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	script := fmt.Sprintf(`CREATE TABLE %v (
				id INT PRIMARY KEY AUTO_INCREMENT,
				username VARCHAR(255),
				email VARCHAR(255), 
				password VARCHAR(255),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
				)`, USERTABLE)

	result, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Printf("selamat anda berhasil membuat table user %v", result)
}

func TestInsert(t *testing.T){
	patern := repository.NewUserImplemen(DataCon())
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)

	defer cancel()

	data := model.UserModel{
		Username: "bimo satrio pandapotan",
		Email: "bimo@gmail.com",
		Password: "bimo1234",
	}

	if err := patern.Insert(ctx, data); err != nil {
		panic(err)
	}

	fmt.Println(data)

}