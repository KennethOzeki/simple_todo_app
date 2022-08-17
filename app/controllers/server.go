package controllers

import (
	"net/http"
	"simple_todo_app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
