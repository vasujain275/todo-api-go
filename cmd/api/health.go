package main

import (
	"net/http"

	"github.com/vasujain275/todo-api-go/internal/utils"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, http.StatusOK, map[string]string{
		"status":  "up",
		"version": "v1",
	})
}
