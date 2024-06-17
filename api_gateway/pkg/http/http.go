package http

import (
	"api_gateway/config"
	"fmt"
	"log"
	"net/http"
)

func Start(handler http.Handler, cfg config.HTTP) {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port), handler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
