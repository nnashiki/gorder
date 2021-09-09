package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goder/models"
	"log"
)

var DB *sql.DB
var err error

func New() *sql.DB {
	DB,err = sql.Open("mysql", "myuser:password@/mydb?parseTime=true")
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	return DB
}

func main() {
	db := New()
	ctx := context.Background()
	m, _ := models.Users().One(ctx,db)
	fmt.Println("jets:", m)
}