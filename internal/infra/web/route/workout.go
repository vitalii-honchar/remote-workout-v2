package route

import (
	"remoteworkout/internal/infra/database/workout"
	"remoteworkout/internal/infra/web/request"

	log "github.com/sirupsen/logrus"
)

func GetWorkouts(wr *workout.WorkoutRepository, r *request.Request, c chan request.Response) {
	log.Println("Handle get workouts")
	workouts, err := wr.FindByCoachId(1)
	log.Infof("Workouts: %v", workouts)
	// responseBytes, err := json.Marshal(workouts)
	if err != nil {
		c <- *request.CreateErrorResponse(500)
	} else {
		c <- *request.CreateResponse(workouts)
	}
}
