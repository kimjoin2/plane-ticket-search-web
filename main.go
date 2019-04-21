package main

import (
	"log"
	"net/http"
	"simpleWebServer/controller"
	"simpleWebServer/webServerConfig"
)

func main() {
	http.HandleFunc("/test", controller.BaseController)
	http.HandleFunc("/", controller.TopPageController)
	log.Println("server start!")
	err := http.ListenAndServe(webServerConfig.GetServerPortData(), nil)
	if err != nil {
		log.Println(err)
	}
}

