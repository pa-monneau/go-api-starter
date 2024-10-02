package handler

import (
	"go-api-starter/serveur/handler/profile"
	"go-api-starter/serveur/handler/tasks"
)

func Matcher() []Matchers {
	matcher := []Matchers{
		{"/profile", profile.HelloProfileHandler},
		{"/task", tasks.HelloTasksHandler},
	}

	return matcher
}
