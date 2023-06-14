package data

import (
	"database/sql"
)

type Models struct {
	URL UrlModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		URL: UrlModel{DB: db},
	}
}
