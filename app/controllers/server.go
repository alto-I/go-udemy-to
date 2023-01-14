package controllers

import (
	"fmt"
	"net/http"
	"udemy_todo/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	fmt.Println(config.Config.Static)
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
