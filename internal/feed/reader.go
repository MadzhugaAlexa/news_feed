package feed

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"news_feed/internal/entities"
)

func Read(url string, out chan<- entities.Item) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	rss := entities.RSS{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range rss.Channel.Items {
		out <- item
	}
}
