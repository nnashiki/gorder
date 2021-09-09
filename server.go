package main

import (
	"goder/router"
	"goder/mysql"
)

func main()  {
	mysql.New()
	e := router.Init()
	e.Start(":1323")
}