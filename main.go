package main

import (
	"github.com/gin-gonic/gin"
	"github.com/livanjimenez/blogposts-timeline/cmd"
	"github.com/livanjimenez/blogposts-timeline/internal/db"
)

func main() {
	r := gin.Default()
	db.Connect()
	r.POST("/posts", cmd.CreatePostHandler)
	r.Run()
}
