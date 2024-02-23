package main

import (
	"context"
	"fmt"
	"log"
	"news_feed/internal/config"
	"news_feed/internal/entities"
	"news_feed/internal/feed"
	"news_feed/internal/handler"
	"news_feed/internal/repo"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	DB_URL := "postgres://alexa:alexa@localhost:5432/rss"

	db, err := pgxpool.New(context.Background(), DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v\n", err)
	}
	defer db.Close()

	repo := repo.NewRepo(db)

	cfg := config.LoadConfig()
	ch := make(chan entities.Item)

	for _, rss := range cfg.RSS {
		ticker := time.NewTicker(time.Duration(cfg.RequestPeriod) * time.Second)
		go func(rss string) {
			for range ticker.C {
				feed.Read(rss, ch)
			}
		}(rss)
	}

	go func(ch chan entities.Item) {
		for {
			item := <-ch
			fmt.Println("storing item")
			repo.AddItem(item)
		}
	}(ch)

	e := echo.New()
	h := handler.NewHandler(repo)
	e.GET("/items", h.GetItems)
	e.Logger.Fatal(e.Start(":8080"))
}
