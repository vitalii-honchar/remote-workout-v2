package student

import "database/sql"

type StudentRepository struct {
	db *sql.DB
}
