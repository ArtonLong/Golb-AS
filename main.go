package main

import (
	"gold-as/src/backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var conn, _ = database.Connect()
	//users := database.QueryUsersData(conn)
	posts := database.QueryPostsData(conn)

	router := gin.Default()
	router.Static("src/frontend/static", "./src/frontend/static")
	router.LoadHTMLGlob("src/frontend/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Blogs",
			"heading": "Posts",
			"posts":   posts,
		})
	})

	router.GET("/post/:slug", func(c *gin.Context) {
		//slug := c.Param("slug")

		c.HTML(http.StatusOK, "post.html", gin.H{
			"title": "Post",
		})
	})

	router.Run("localhost:8080")
}
