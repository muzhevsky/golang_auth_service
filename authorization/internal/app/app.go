package app

import (
	"authorization/config"
	"authorization/internal/controller/http/v1"
	"authorization/internal/infrastructure"
	"authorization/internal/usecase"
	"authorization/pkg/http"
	"authorization/pkg/logger"
	"authorization/pkg/postgres"
	"authorization/pkg/smtp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run - NewConfig: %w", err))
	}

	log := logger.New(cfg.Level)

	jwtGenerator := &infrastructure.JWTGenerator{}

	// Repository
	pg, err := postgres.New(cfg.PG.Url, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewUser: %w", err))
	}
	defer pg.Close()

	userRepo := infrastructure.NewUserRepo(pg)
	userDataRepo := infrastructure.NewBcryptHashProvider(10) // TODO вынести в конфиг (посоветоваться)
	verificationRepo := infrastructure.NewVerificationRepo(pg)
	sessionRepo := infrastructure.NewSessionRepo(pg)

	// Other infrastructure

	smtpClient := smtp.New(cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host)
	smtpMailer := infrastructure.NewSmtpMailer(smtpClient)

	userUseCase := usecase.NewUser(
		userRepo,
		verificationRepo,
		userDataRepo,
	)

	verificationUseCase := usecase.NewVerificationUseCase(
		userRepo,
		verificationRepo,
		smtpMailer,
	)
	sessionUseCase := usecase.NewSessionUseCase(jwtGenerator, sessionRepo)

	router := gin.New()
	v1.NewRouter(router, log, userUseCase, verificationUseCase, sessionUseCase)
	httpServer := http.New(
		router,
		http.FullAddress(cfg.Addr, cfg.HTTP.Port))

	select {
	case err := <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
