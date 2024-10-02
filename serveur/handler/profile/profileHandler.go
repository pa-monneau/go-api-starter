package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func HelloProfileHandler(w http.ResponseWriter, _r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Hello from profile",
		Status:  "success",
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		fmt.Printf(err.Error())
	}
}
