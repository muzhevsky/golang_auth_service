package app

import (
	"api_gateway/config"
	v1 "api_gateway/internal/controllers/http/v1"
	"api_gateway/pkg/http"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("invalid config, %v", err.Error())
	}

	authUrl, _ := url.Parse("http://" + cfg.AuthHost + ":" + cfg.AuthPort)

	router := gin.New()
	router.Use(v1.NewProxy(authUrl, cfg.AuthHost).Handle)
	http.Start(router, cfg.HTTP)
}
