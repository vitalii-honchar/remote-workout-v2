package route

import (
	"remoteworkout/internal/domain"
	"remoteworkout/internal/infra/database/workout"
	"remoteworkout/internal/infra/web/request"

	log "github.com/sirupsen/logrus"
)

var workouts = []domain.Workout{
	{
		Coach:       "Coach 1",
		Id:          1,
		Name:        "Workout 1",
		Description: "Test 1",
		Videos: []domain.WorkoutVideo{
			{Link: "1"},
		},
	},
	{
		Coach:       "Coach 2",
		Id:          2,
		Name:        "Workout 2",
		Description: "Test 2",
		Videos: []domain.WorkoutVideo{
			{Link: "1"},
		},
	},
}

func GetWorkouts(wr *workout.WorkoutRepository, r *request.Request, c chan request.Response) {
	log.Println("Handle get workouts")
	// responseBytes, err := json.Marshal(workouts)
	// if err != nil {
	// 	fmt.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// } else {
	// 	w.Write(responseBytes)
	// }
}
