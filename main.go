package main

import (
	"im/database"
	"im/router"
)

func main() {
	database.Mysql()
	router.Router()
}
