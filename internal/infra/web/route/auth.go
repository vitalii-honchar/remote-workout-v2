package route

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle get login")
	response := map[string]bool{
		"success": true,
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(responseBytes)
	}
}
