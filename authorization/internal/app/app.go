package app

import (
	"authorization/config"
	http2 "authorization/internal/controllers/http"
	"authorization/internal/controllers/http/middleware"
	"authorization/internal/controllers/http/v1"
	pg2 "authorization/internal/infrastructure/datasources/pg"
	"authorization/internal/infrastructure/services/hash"
	mailers2 "authorization/internal/infrastructure/services/mailers"
	tokens2 "authorization/internal/infrastructure/services/tokens"
	"authorization/internal/repositories"
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
	"time"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run 1 : %w", err))
	}

	// Packages
	logger := logger.New("error")

	jwt := jwt.New(cfg.SigningString)

	size, err := strconv.Atoi(cfg.PG.PoolMax)

	if err != nil {
		logger.Fatal(fmt.Errorf("app - Run 2: %w", err))
	}
	pg, err := postgres.New(cfg.PG, postgres.MaxPoolSize(size))
	if err != nil {
		logger.Fatal(fmt.Errorf("app - Run 3: %w", err))
	}
	defer pg.Close()

	// Infrastructure
	accountDS := pg2.NewAccountDatasource(pg)
	sessionDS := pg2.NewSessionDatasource(pg)
	verificationDS := pg2.NewPgVerificationDatasource(pg)

	bcryptHashProvider := hash.NewBcryptHashProvider()
	accessTokenProvider := tokens2.NewJwtProvider(jwt)
	refreshTokenGenerator := tokens2.NewHashRefreshTokenGenerator(bcryptHashProvider)

	smtpClient := smtp.New(cfg.SMTP.Username, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)
	smtpMailer := mailers2.NewSmtpMailer(smtpClient)
	verificationMailer := mailers2.NewVerificationMailer(smtpMailer)
	sessionManager := tokens2.NewTokenManager(
		tokens2.TokenConfiguration{
			AccessTokenDuration:  time.Minute * 30,
			RefreshTokenDuration: time.Hour * 24 * 30,
			Issuer:               "TODO",
		},
		accessTokenProvider,
		refreshTokenGenerator)

	// Repository

	userRepository := repositories.NewUserRepo(accountDS)
	verificationRepo := repositories.NewVerificationRepo(verificationDS)
	sessionRepository := repositories.NewSessionRepository(sessionDS)

	// UseCases

	createUserUseCase := usecases.NewCreateUserUseCase(
		userRepository,
		sessionRepository,
		sessionManager,
		bcryptHashProvider,
		verificationMailer,
	)

	verificationUseCase := usecases.NewVerificationUseCase(
		userRepository,
		verificationRepo,
		verificationMailer,
	)

	signInUseCase := usecases.NewSignInUseCase(
		userRepository,
		sessionRepository,
		bcryptHashProvider,
		sessionManager,
	)

	refreshSessionUseCase := usecases.NewRefreshSessionUseCase(
		userRepository,
		sessionRepository,
		sessionManager,
	)

	requestVerificationUseCase := usecases.NewRequestVerificationRequest(
		userRepository,
		verificationRepo,
		verificationMailer,
	)
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
	v1.NewRequestVerificationRouter(router, createUserUseCase, requestVerificationUseCase, logger)

	http.Start(router, cfg.HTTP)
}
