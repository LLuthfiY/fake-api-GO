package main

import (
	"github.com/LLuthfiY/fake-api-GO/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.User(r)
	routes.Article(r)
	r.Run(":4245")
}
