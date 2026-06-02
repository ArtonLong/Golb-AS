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
	rows, err := conn.Query(context.Background(), "SELECT id, username, password_hash, created_at FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []structs.User

	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password_hash, &user.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func InsertUserData(conn *pgx.Conn, username string, password string) bool {
	_, err := conn.Exec(context.Background(), "INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, password)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func FindUser(conn *pgx.Conn, username string) structs.User {
	var user structs.User
	err := conn.QueryRow(context.Background(), "SELECT id, username, password_hash FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Password_hash)
	if err != nil {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAA")
		log.Fatal(err)
	}
	return user
}

func FindUserById(conn *pgx.Conn, id int) structs.User {
	var user structs.User
	err := conn.QueryRow(context.Background(), "SELECT id, username FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username)
	if err != nil {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAA")
		log.Fatal(err)

	}
	return user
}

func InsertPostData(conn *pgx.Conn, author_id int, title string, slug string, summary string, body string) {
	tag, err := conn.Exec(context.Background(), "INSERT INTO posts (author_id, title, slug, summary, body) VALUES ($1, $2, $3, $4, $5)", author_id, title, slug, summary, body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tag)
}
