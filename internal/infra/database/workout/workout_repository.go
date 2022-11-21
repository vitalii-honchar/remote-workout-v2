package workout

import (
	"database/sql"
	"remoteworkout/internal/domain"
)

const queryFindByCoachId = `
	SELECT 
		w.id, 
		w.coach_id, 
		w.name, 
		w.description
	FROM workout w 
	WHERE w.coach_id = $1`

type WorkoutRepository struct {
	db *sql.DB
	vr *VideoRepository
}

func CreateWorkoutRepository(db *sql.DB, vr *VideoRepository) *WorkoutRepository {
	return &WorkoutRepository{db: db, vr: vr}
}

func (wr *WorkoutRepository) FindByCoachId(coachId int) ([]domain.Workout, error) {
	rows, err := wr.db.Query(queryFindByCoachId, coachId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var workouts []domain.Workout

	for rows.Next() {
		var workout domain.Workout
		rows.Scan(&workout.Id, &workout.CoachId, &workout.Name, &workout.Description)
		workouts = append(workouts, workout)
	}

	return workouts, nil
}
