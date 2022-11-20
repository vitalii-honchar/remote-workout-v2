package price

import "database/sql"

type PriceRepository struct {
	db *sql.DB
}
