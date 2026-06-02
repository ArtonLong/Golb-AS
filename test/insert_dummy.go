package main

import "gold-as/src/backend/database"

// #### sql queries
// CREATE TABLE users (
//     id SERIAL PRIMARY KEY,
//     username VARCHAR(50) UNIQUE NOT NULL,
//     password_hash TEXT NOT NULL,
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

// CREATE TABLE posts (
//     id SERIAL PRIMARY KEY,
//     author_id INT REFERENCES users(id) ON DELETE SET NULL,
//     title VARCHAR(255) NOT NULL,
//     slug VARCHAR(255) UNIQUE NOT NULL,
//     summary TEXT,
//     body TEXT NOT NULL,
//     meta_json JSONB DEFAULT '{}',
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

func main() {
	var conn, _ = database.Connect()
	database.InsertUserData(conn, "bob", "123abc")
	database.InsertPostData(conn, 1, "first post", "first-post", "this is the first post", "lorem ipsum djbafidvbaoibvdb9aosvbaovb")
}
