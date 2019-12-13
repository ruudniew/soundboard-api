package main

import (
	"sbapi/customdb"
	"sbapi/gin"
)

func main() {
	customDB := customdb.CustomDB{}

	api := gin.API{}
	api.EventService = &customDB

	api.Start("", "3300")
}
