package route

import (
	"remoteworkout/internal/infra/web"

	log "github.com/sirupsen/logrus"
)

func GetLogin(r *web.Request, c chan web.Response) {
	log.Info("Handle get login")
	response := map[string]bool{
		"success": true,
	}
	c <- web.Response{Body: response}
}
