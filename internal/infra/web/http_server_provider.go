package web

import (
	"net/http"
	"remoteworkout/internal/infra/database/coach"
	"remoteworkout/internal/infra/database/price"
	"remoteworkout/internal/infra/database/student"
	"remoteworkout/internal/infra/database/workout"
	"remoteworkout/internal/infra/web/request"
	"remoteworkout/internal/infra/web/route"
)

type DatabaseContext struct {
	CoachRepository   *coach.CoachRepository
	PriceRepository   *price.PriceRepository
	StudentRepository *student.StudentRepository
	WorkoutRepository *workout.WorkoutRepository
}

func ProvideHttpServer(dc *DatabaseContext) *HttpServer {
	server := HttpServer{Origin: http.NewServeMux()}

	server.Get("/auth/login", route.GetLogin)
	server.Get("/workout", func(r *request.Request, c chan request.Response) {
		route.GetWorkouts(dc.WorkoutRepository, r, c)
	})

	return &server
}
