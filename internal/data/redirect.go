package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (m UrlModel) QueryWithShort(shortUrl string) (*UrlData, error) {
	query := `
		SELECT
			*
		FROM url
		WHERE
			short_url = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var data UrlData

	err := m.DB.QueryRowContext(ctx, query, shortUrl).Scan(
		&data.ID,
		&data.ShortURL,
		&data.LongURL,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			// * handle special error
			return nil, err

		default:
			return nil, err
		}

	}

	return &data, nil
}
