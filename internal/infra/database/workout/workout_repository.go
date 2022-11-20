package workout

import "database/sql"

type WorkoutRepository struct {
	db *sql.DB
}
