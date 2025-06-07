package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(){
	dsn := os.Getenv("MYSQL_DSN")
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败： ", err)
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(1 * time.Hour)

	if err = DB.Ping(); err != nil {
		log.Fatal("无法连接到数据库：", err)
	}
}

func retrieve(id int) (post Post, err error){
	post = Post{}
	err = DB.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error){
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

