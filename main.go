package main

import (
  "github.com/gin-gonic/gin"
	"github.com/tranquockhang1810/Golang-Gin-Lab04/routes"
)

func main() {
  r := gin.Default()
	routes.StudentRoute(r)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}