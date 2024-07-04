package app

import (
	"api_gateway/config"
	v1 "api_gateway/internal/controllers/http/v1"
	"api_gateway/pkg/http"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("invalid config, %v", err.Error())
	}

	authUrl := "http://" + cfg.AuthHost + ":" + cfg.AuthPort
	applicationUrl := "http://" + cfg.ApplicationHost + ":" + cfg.ApplicationPort

	serviceMap := make(map[string]string)
	serviceMap[cfg.AuthHost] = authUrl
	serviceMap[cfg.ApplicationHost] = applicationUrl

	router := gin.New()
	router.Use(v1.NewAuthProxy(authUrl).Handle)
	router.Use(v1.NewProxy(serviceMap).Handle)
	router.Static("/static", "./static")
	http.Start(router, cfg.HTTP)
}
