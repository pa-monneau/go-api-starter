package serveur

import (
	"fmt"
	"go-api-starter/serveur/handler"
	"net/http"
)

func Init() {
	fmt.Printf("init serveur...\n")
	matchers := handler.Matcher()

	for _, matcher := range matchers {
		http.HandleFunc(matcher.Path, matcher.Handler)
	}

	fmt.Printf("server started at http://localhost:8080...\n")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf(err.Error())
	}
}
