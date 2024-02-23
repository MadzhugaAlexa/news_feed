package repo

import (
	"context"
	"errors"
	"fmt"
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

	t, err := tx.Exec(
		context.Background(),
		sql,
		item.GUID, item.Title, item.Link, item.PdaLink, item.Description, item.PubDate,
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
func (r *Repo) ReadItems() ([]entities.Item, error) {
	items := make([]entities.Item, 0)

	rows, err := r.db.Query(context.Background(), "select * from news;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := entities.Item{}

		err = rows.Scan(
			&item.GUID, &item.Title, &item.Link, &item.PdaLink, &item.Description, &item.PubDate,
			&item.Category, &item.Author, &item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}
