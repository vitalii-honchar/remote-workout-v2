package domain

type Workout struct {
	Coach       string
	Id          int
	Name        string
	Description string
	Videos      []WorkoutVideo
}

type WorkoutVideo struct {
	Link string
}
