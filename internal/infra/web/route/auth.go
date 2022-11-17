package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle get login")
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
