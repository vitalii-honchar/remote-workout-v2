package coach

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type CoachRepository struct {
	db *sql.DB
}

func CreateCoachRepository(db *sql.DB) *CoachRepository {
	return &CoachRepository{db: db}
}

func (cr *CoachRepository) FindById(id int) {
	log.Infof("Find by id: %v", id)
}
