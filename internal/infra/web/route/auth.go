package route

import (
	"remoteworkout/internal/infra/web/request"

	log "github.com/sirupsen/logrus"
)

func GetLogin(r *request.Request, c chan request.Response) {
	log.Info("Handle get login")
	response := map[string]bool{
		"success": true,
	}
	c <- request.Response{Body: response}
}
