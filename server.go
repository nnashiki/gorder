package main

import (
	"goder/router"
)

func main()  {
	e := router.Init()
	e.Start(":1323")
}