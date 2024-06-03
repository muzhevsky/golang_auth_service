package app

import (
	"authorization/config"
	http2 "authorization/internal/controllers/http"
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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
	"time"
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

	userDS := pg2.NewPgUserDatasource(pg)
	sessionDS := pg2.NewSessionRepo(pg)
	verificationDS := pg2.NewPgVerificationDatasource(pg)

	bcryptHashProvider := hash.NewBcryptHashProvider()
	accessTokenProvider := tokens2.NewJwtProvider(jwt)
	refreshTokenGenerator := tokens2.NewHashRefreshTokenGenerator(bcryptHashProvider)

	userRepository := repositories.NewUserRepo(userDS)
	verificationRepo := repositories.NewVerificationRepo(verificationDS)
	sessionRepository := repositories.NewSessionRepository(sessionDS)

	// Other infrastructure

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

	// UseCases

	userUseCase := usecases.NewCreateUserUseCase(
		userRepository,
		verificationRepo,
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
	// Controllers

	router := gin.New()
	router.HandleMethodNotAllowed = true
	http2.InitServiceMiddleware(router)
	v1.NewSignUpRouter(router, userUseCase, verificationUseCase, log)
	v1.NewVerificationRouter(router, verificationUseCase, log)
	v1.NewSignInRouter(router, signInUseCase, log)
	v1.NewRefreshSessionRouter(router, refreshSessionUseCase, log)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//http2.NewAuthorizationRouter(router, userUseCase, log, sessionUseCase)

	http.Start(router, cfg.HTTP)
}
