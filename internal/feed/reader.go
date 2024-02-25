package feed

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"news_feed/internal/config"
	"news_feed/internal/entities"
	"news_feed/internal/repo"
	"time"
)

func Read(url string, out chan<- entities.Item, ech chan<- entities.Error) {
	response, err := http.Get(url)
	if err != nil {
		ech <- entities.Error{Error: err}
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ech <- entities.Error{Error: err}
		return
	}

	rss := entities.RSS{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		ech <- entities.Error{Error: err}
		return
	}

	for _, item := range rss.Channel.Items {
		out <- item
	}
}

func FeatchFeeds(repo *repo.Repo) {
	cfg := config.LoadConfig("./config.json")
	ch := make(chan entities.Item)
	ech := make(chan entities.Error)

	for _, rss := range cfg.RSS {
		ticker := time.NewTicker(time.Duration(cfg.RequestPeriod) * time.Second)
		go func(rss string) {
			for range ticker.C {
				Read(rss, ch, ech)
			}
		}(rss)
	}

	go func(ch chan entities.Item) {
		for {
			item := <-ch
			repo.AddItem(item)
		}
	}(ch)

	go func(ech chan entities.Error) {
		for {
			err := <-ech
			log.Printf("Ошибка: %v\n", err)
		}
	}(ech)
}
