package app

import (
	"authorization/config"
	"authorization/internal/controllers/http/v1"
	"authorization/internal/infrastructure"
	"authorization/internal/usecases"
	"authorization/pkg/http"
	"authorization/pkg/jwt"
	"authorization/pkg/logger"
	"authorization/pkg/postgres"
	"authorization/pkg/smtp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run: %w", err))
	}

	log := logger.New("error")

	jwt := jwt.New(cfg.SigningString)

	// Repository
	size, err := strconv.Atoi(cfg.PG.PoolMax)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run: %w", err))
	}
	pg, err := postgres.New(cfg.PG, postgres.MaxPoolSize(size))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run: %w", err))
	}
	defer pg.Close()

	userRepo := infrastructure.NewUserRepo(pg)
	bcryptHashProvider := infrastructure.NewBcryptHashProvider()
	accessTokenProvider := infrastructure.NewJwtProvider(jwt)
	refreshTokenGenerator := infrastructure.NewHashRefreshTokenGenerator(bcryptHashProvider)
	verificationRepo := infrastructure.NewVerificationRepo(pg)
	sessionRepo := infrastructure.NewSessionRepo(pg)

	// Other infrastructure

	smtpClient := smtp.New(cfg.SMTP.Username, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)
	smtpMailer := infrastructure.NewSmtpMailer(smtpClient)

	// UseCases

	userUseCase := usecases.NewUser(
		userRepo,
		verificationRepo,
		bcryptHashProvider,
	)

	verificationUseCase := usecases.NewVerificationUseCase(
		userRepo,
		verificationRepo,
		smtpMailer,
	)
	sessionUseCase := usecases.NewSessionUseCase(accessTokenProvider, refreshTokenGenerator, sessionRepo, bcryptHashProvider)

	// Controllers

	router := gin.New()
	v1.InitServiceMiddleware(router)
	v1.NewAuthenticationRouter(router, log, userUseCase, sessionUseCase, verificationUseCase)
	v1.NewAuthorizationRouter(router, userUseCase, log, sessionUseCase)

	http.Start(router, cfg.HTTP)
}
