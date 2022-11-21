package workout

import "database/sql"

type VideoRepository struct {
	db *sql.DB
}

func CreateVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

const queryWorkoutVideosByWorkoutId = `
	SELECT 
		v.id, 
		v.link
	FROM video_to_workout vw 
	INNER JOIN video v ON v.id = vw.video_id
	WHERE vw.id = $1`
