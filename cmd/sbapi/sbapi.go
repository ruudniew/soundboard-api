package main

import (
	"sbapi/customdb"
	"sbapi/gin"
)

func main() {
	customDB := customdb.CustomDB{}

	api := gin.API{}
	api.EventService = &customDB

	api.Start("", "3300", []string{"http://127.0.0.1:3400", "http://localhost:3200", "https://soundboard.ruudniew.com"})
}
