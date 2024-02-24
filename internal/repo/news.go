package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"news_feed/internal/entities"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

var errFailedToSave = errors.New("failed to save")

func (r *Repo) AddItem(item entities.Item) error {
	item.CreatedAt = time.Now()
	tx, _ := r.db.Begin(context.Background())
	defer tx.Commit(context.Background())

	sql := `INSERT INTO news(guid, title, link, pdalink,description,pubDate,category,author, created_at) 
	values($1, $2, $3, $4, $5, $6, $7, $8, $9) `

	putDate, err := time.Parse(time.RFC1123, item.PubDate)
	if err != nil {
		putDate, err = time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			putDate, err = time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.PubDate)
			if err != nil {
				log.Fatalf("failed to parse data %#v\n", item)
			}

		}
	}

	t, err := tx.Exec(
		context.Background(),
		sql,
		item.GUID, item.Title, item.Link, item.PdaLink, item.Description, putDate.Unix(),
		item.Category, item.Author, item.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if t.RowsAffected() != 1 {
		return errFailedToSave
	}

	return nil
}
func (r *Repo) ReadItems(limit int) ([]entities.Post, error) {
	items := make([]entities.Post, 0)

	rows, err := r.db.Query(context.Background(), "select id, title, link, description, pubdate from news order by created_at desc limit $1", limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := entities.Post{}

		err = rows.Scan(
			&item.ID, &item.Title, &item.Link, &item.Content, &item.PubDate,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}
