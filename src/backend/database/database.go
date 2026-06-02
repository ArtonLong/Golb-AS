package database

import (
	"context"
	"fmt"
	"gold-as/src/backend/structs"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect() (*pgx.Conn, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db := os.Getenv("DB_URL")

	conn, err := pgx.Connect(context.Background(), db)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func QueryPostsData(conn *pgx.Conn) []structs.Post {
	rows, err := conn.Query(context.Background(), "SELECT id, author_id, title, slug, summary, body, meta_json, created_at, updated_at FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []structs.Post

	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.Id, &post.Author_id, &post.Title, &post.Slug, &post.Summary, &post.Body, &post.Meta_json, &post.Created_at, &post.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	return posts
}

func QueryUsersData(conn *pgx.Conn) []structs.User {
	rows, err := conn.Query(context.Background(), "SELECT id, username, email, password_hash, created_at FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []structs.User

	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password_hash, &user.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func InsertUserData(conn *pgx.Conn) {
	username := "bob"
	email := "bob@gmail.com"
	password := "123abc"

	tag, err := conn.Exec(context.Background(), "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", username, email, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tag)
}

func InsertPostData(conn *pgx.Conn) {
	author_id := 1
	title := "first post"
	slug := "first-post"
	summary := "this is the first post"
	body := "lorem ipsum bisbvdisbviasbdvosbvoabvdaiosdvbaivbaidvhb"

	tag, err := conn.Exec(context.Background(), "INSERT INTO posts (author_id, title, slug, summary, body) VALUES ($1, $2, $3, $4, $5)", author_id, title, slug, summary, body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tag)
}
