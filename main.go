package main

import (
	"gold-as/src/backend/auth"
	"gold-as/src/backend/database"
	"gold-as/src/backend/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	var conn, _ = database.Connect()
	posts := database.QueryPostsData(conn)

	r := gin.Default()
	r.Static("src/frontend/static", "./src/frontend/static")
	r.LoadHTMLGlob("src/frontend/templates/*")

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Blogs",
			"heading": "Posts",
			"posts":   posts,
		})
	})

	public.GET("/post/:slug", func(c *gin.Context) {
		slug := c.Param("slug")

		var post structs.Post

		for _, val := range posts {
			if val.Slug == slug {
				post = val
				break
			}
		}

		c.HTML(http.StatusOK, "post.html", gin.H{
			"title": "Post",
			"post":  post,
		})
	})

	public.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Reqister",
		})
	})

	public.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})

	public.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}

		ok := database.InsertUserData(conn, username, string(hashedPassword))
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or email not uniqe"})
			return
		}

	})

	public.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user := database.FindUser(conn, username)

		err := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		token, err := auth.CreateToken(user.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ""})
			return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	protected := r.Group("/api/admin")
	protected.Use(auth.JwtAuthMiddleware())
	protected.GET("/new-blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new-blog.html", gin.H{
			"title": "new blog",
		})
	})
	protected.POST("/new-blog", func(c *gin.Context) {
		user_id, err := auth.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		title := c.PostForm("title")
		summary := c.PostForm("summary")
		body := c.PostForm("body")

		slug := strings.ToLower(title)
		slug = strings.ReplaceAll(slug, " ", "-")

		database.InsertPostData(conn, int(user_id), title, slug, summary, body)

	})

	r.Run("localhost:8080")
}
