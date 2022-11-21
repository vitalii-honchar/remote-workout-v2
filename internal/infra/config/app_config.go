//go:build wireinject
// +build wireinject

package config

import (
	"database/sql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"remoteworkout/internal/infra/database/workout"
	"remoteworkout/internal/infra/web"
)

type AppContext struct {
	Server            *web.HttpServer
	WorkoutRepository *workout.WorkoutRepository
}

func ProvideDatabaseContext() *web.DatabaseContext {
	wire.Build(ContextSet)
	return &web.DatabaseContext{}
}

func ProvideHttpServer() *web.HttpServer {
	wire.Build(web.ProvideHttpServer, ProvideDatabaseContext)
	return &web.HttpServer{}
}

func ProvideDb() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://user:password@localhost:5432/remote_workout?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func ProvideAppContext() *AppContext {
	wire.Build(ContextSet, ProvideHttpServer)
	return &AppContext{}
}

var ContextSet = wire.NewSet(
	ProvideDb,
	workout.CreateVideoRepository,
	workout.CreateWorkoutRepository,
	wire.Struct(new(web.DatabaseContext), "WorkoutRepository"),
	wire.Struct(new(AppContext), "*"),
)
