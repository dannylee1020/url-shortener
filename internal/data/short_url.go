package data

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"github.com/dannylee1020/url-shortener/internal/helper"
)

type UrlData struct {
	ID       int64  `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

type UrlModel struct {
	DB *sql.DB
}

// const baseUrl = "tinyurl.com/"

func (m UrlModel) QueryWithLong(url string) (*UrlData, error) {
	query := `
		SELECT
			id,
			short_url,
			long_url
		FROM url
		WHERE
			long_url = $1
	`
	var data UrlData

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, url).Scan(
		&data.ID,
		&data.ShortURL,
		&data.LongURL,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m UrlModel) InsertURL(data *UrlData, shortUrl string) error {
	query := `
		INSERT INTO url (short_url, long_url)
		VALUES ($1, $2)
		RETURNING (id)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, shortUrl, data.LongURL).Scan(&data.ID)
	if err != nil {
		return err
	}

	return nil
}

func GenerateShortURL() string {
	num := rand.Int()
	hash := helper.EncodeBase62(num)

	return hash

}
