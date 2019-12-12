package main

import (
  "sbapi/gin"
  "sbapi/customdb"
)

func main() {
  customDB := customdb.CustomDB{}

	api := gin.API{}
  api.EventService = &customDB

	api.Start("", "3310")
}