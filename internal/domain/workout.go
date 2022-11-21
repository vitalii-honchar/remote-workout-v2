package domain

type Workout struct {
	Id          int    `json:"id"`
	CoachId     int    `json:"coach_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Videos      []WorkoutVideo
}

type WorkoutVideo struct {
	Id   int
	Link string
}
