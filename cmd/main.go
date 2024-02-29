package main

import (
	"context"
	"log"
	"news_feed/internal/feed"
	"news_feed/internal/handler"
	"news_feed/internal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	DB_URL := "postgres://alexa:alexa@localhost:5432/rss"

	db, err := pgxpool.New(context.Background(), DB_URL)
	if err != nil {
		log.Fatalf("Не смогли подключиться к БД: %v\n", err)
	}
	defer db.Close()

	repo := repo.NewRepo(db)

	feed.FeatchFeeds(repo)

	e := echo.New()
	h := handler.NewHandler(repo)
	e.GET("/news/", h.GetItems)
	e.GET("/news/:limit", h.GetItems)
	e.File("/", "./webapp/index.html")
	e.Static("/", "./webapp/")
	e.Logger.Fatal(e.Start(":80"))
}
