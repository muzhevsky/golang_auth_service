package http

import (
	"fmt"
	"log"
	"net/http"
	"smartri_app/config"
)

func Start(handler http.Handler, cfg config.HTTP) {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port), handler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
