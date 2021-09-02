package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goder/models"
	"context"
)

var db *sql.DB
var err error

func New() *sql.DB {
	db,err = sql.Open("mysql", "myuser:password@/mydb")
	return db
}

func main() {
	db := New()
	ctx := context.Background()
	m, _ := models.Users().One(ctx,db)
	fmt.Println("jets:", m)
}