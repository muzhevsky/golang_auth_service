package app

import (
	"authorization/cmd/app/factories"
	"authorization/config"
	http2 "authorization/controllers/http"
	"authorization/controllers/http/middleware"
	v1 "authorization/controllers/http/v1"
	"authorization/internal/infrastructure/services/hash"
	mailers2 "authorization/internal/infrastructure/services/mailers"
	tokens2 "authorization/internal/infrastructure/services/tokens"
	"authorization/internal/usecases"
	"authorization/pkg/http"
	"authorization/pkg/jwt"
	"authorization/pkg/logger"
	"authorization/pkg/postgres"
	"authorization/pkg/redis"
	"authorization/pkg/smtp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run 1 : %w", err))
	}

	// Packages
	logger := logger.New("error")
	redisClient := redis.NewRedisClient(cfg.Redis)
	jwt := jwt.New(cfg.SigningString)

	size, err := strconv.Atoi(cfg.PG.PoolMax)

	if err != nil {
		logger.Fatal(fmt.Errorf("app - Run 2: %w", err))
	}
	pgClient, err := postgres.New(cfg.PG, postgres.MaxPoolSize(size))
	if err != nil {
		logger.Fatal(fmt.Errorf("app - Run 3: %w", err))
	}
	defer pgClient.Close()

	// Infrastructure

	bcryptHashProvider := hash.NewBcryptHashProvider()

	accessTokenProvider := tokens2.NewJwtProvider(jwt)
	refreshTokenGenerator := tokens2.NewHashRefreshTokenGenerator(bcryptHashProvider)
	sessionManager := tokens2.NewTokenManager(cfg.TokenConfiguration, accessTokenProvider, refreshTokenGenerator)

	smtpClient := smtp.New(cfg.SMTP.Username, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)
	verificationMailer := mailers2.NewSMTPVerificationMailer(smtpClient)

	// Repository

	accountRepository := factories.CreatePGAccountRepo(pgClient)
	sessionRepository := factories.CreateSessionRepo(pgClient, redisClient)
	verificationRepo := factories.CreateRedisVerificationRepo(redisClient)
	deviceRepository := factories.CreateDeviceRepo(pgClient, redisClient)

	// UseCases

	createUserUseCase := usecases.NewCreateUserUseCase(
		accountRepository,
		sessionRepository,
		sessionManager,
		bcryptHashProvider,
		verificationMailer,
	)

	verificationUseCase := usecases.NewVerificationUseCase(
		accountRepository,
		verificationRepo,
		verificationMailer,
	)

	signInUseCase := usecases.NewSignInUseCase(
		accountRepository,
		sessionRepository,
		bcryptHashProvider,
		sessionManager,
	)

	refreshSessionUseCase := usecases.NewRefreshSessionUseCase(
		accountRepository,
		sessionRepository,
		sessionManager,
	)

	requestVerificationUseCase := usecases.NewRequestVerificationUseCase(
		accountRepository,
		verificationRepo,
		verificationMailer,
	)

	checkVerificationUseCase := usecases.NewCheckVerificationUsecase(accountRepository)

	getAccountDevicesUseCase := usecases.NewGetAccountDevicesUseCase(deviceRepository)
	// Controllers

	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(middleware.NewAuthenticationHandler(sessionRepository, sessionManager).HandleAuth)

	http2.InitServiceMiddleware(router, logger)
	v1.NewAuthenticationController(router)
	v1.NewSignUpController(router, createUserUseCase, verificationUseCase, logger)
	v1.NewVerificationController(router, verificationUseCase, logger)
	v1.NewSignInController(router, signInUseCase, logger)
	v1.NewRefreshSessionController(router, refreshSessionUseCase, logger)
	v1.NewRequestVerificationController(router, createUserUseCase, requestVerificationUseCase, logger)
	v1.NewCheckVerificationController(router, checkVerificationUseCase)
	v1.NewGetAccountDevices(router, getAccountDevicesUseCase)

	http.Start(router, cfg.HTTP)
}
