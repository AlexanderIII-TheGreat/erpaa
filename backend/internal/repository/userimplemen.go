package repository

import (
	"context"
	"database/sql"
	"erpaa-backend/internal/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type UserImpl struct {
	DB *sql.DB
}

func NewUserImplemen(db *sql.DB) UserPatern{
	return &UserImpl{DB: db}
}

func (impl *UserImpl) FindUserPassword(ctx context.Context,username string) (*model.UserModel, error) {

	// tetap pakai env
	if err := godotenv.Load("/var/www/erpaa/app/backend/.env"); err != nil {
		return nil, err
	}

	table := os.Getenv("DB_USER")

	query := fmt.Sprintf(`
		SELECT username, password
		FROM %s
		WHERE username = ?
		LIMIT 1
	`, table)

	row := impl.DB.QueryRowContext(ctx, query, username)

	var user model.UserModel
	if err := row.Scan( &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (rpu *UserImpl) Insert(ctx context.Context, user model.UserModel)(error){

	if err := godotenv.Load("/var/www/erpaa/app/backend/.env"); err != nil {
		return err
	}

	Table := os.Getenv("DB_USER")

	script := fmt.Sprintf(`
	INSERT INTO %v (username, email, password) VALUES (?,?,?)`, Table)

	if _, err := rpu.DB.ExecContext(ctx, script, &user.Username, &user.Email, &user.Password ); err != nil {
		return err
	}

	return nil
	
}

func (rpu *UserImpl) Update(ctx context.Context, id int,  user model.UserModel) (error){
	if err := godotenv.Load("/var/www/erpaa/app/backend/.env"); err != nil {
		return err
	}

	Table := os.Getenv("DB_USER")
	Script := fmt.Sprintf("UPDATE FROM %v SET (username, email,password) VALUES (?,?,?) WHERE id = ?", Table)
	

	if _, err := rpu.DB.ExecContext(ctx,Script,&user.Username, &user.Email, &user.Password, id); err != nil {
		return err
	}

	return nil
}

