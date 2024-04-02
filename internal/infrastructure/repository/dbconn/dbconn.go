package dbconn

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	DB       *sql.DB
	ctxTxKey = struct{}{}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	boil.DebugMode = true
	boil.SetLocation(time.Local)
	boil.SetDB(DB)
}

func GetContextExecutor(ctx context.Context) boil.ContextExecutor {
	if tx, ok := ctx.Value(ctxTxKey).(*sql.Tx); ok {
		return tx
	}

	return boil.GetContextDB()
}

var RunTx = func(ctx context.Context, fn func(context.Context) error) (err error) {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			fmt.Println("Rollback transaction due to panic or error")
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Rollback error: %v\n", rbErr)
			}
			if r != nil {
				err = fmt.Errorf("panic occurred: %v", r)
			}
		}
	}()

	ctx = context.WithValue(ctx, ctxTxKey, tx)

	if err = fn(ctx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
