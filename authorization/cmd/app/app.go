package app

import (
	"authorization/config"
	http2 "authorization/controllers/http"
	"authorization/controllers/http/middleware"
	v1 "authorization/controllers/http/v1"
	pg2 "authorization/internal/infrastructure/datasources/pg"
	"authorization/internal/infrastructure/datasources/pg/commands/accounts"
	"authorization/internal/infrastructure/datasources/pg/commands/sessions"
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
	selectAccountByIdCommand := accounts.NewSelectAccountByIdPGCommand(pg)
	selectAccountByEmailCommand := accounts.NewSelectAccountByEmailPGCommand(pg)
	selectAccountByLoginCommand := accounts.NewSelectAccountByLoginPGCommand(pg)
	updateAccountByIdCommand := accounts.NewUpdateAccountByIdPGCommand(pg)
	insertAccountCommand := accounts.NewInsertAccountPGCommand(pg)

	selectSessionByIdCommand := sessions.NewSelectSessionByIdPGCommand(pg)
	selectSessionByAccessTokenCommand := sessions.NewSelectSessionByAccessTokenPGCommand(pg)
	selectSessionsByAccountIdCommand := sessions.NewSelectSessionByAccountIdPGCommand(pg)
	insertSessionCommand := sessions.NewInsertSessionPGCommand(pg)
	updateSessionByIdCommand := sessions.NewUpdateSessionByIdPGCommand(pg)

	verificationDS := pg2.NewPgVerificationDatasource(pg)

	bcryptHashProvider := hash.NewBcryptHashProvider()

	accessTokenProvider := tokens2.NewJwtProvider(jwt)
	refreshTokenGenerator := tokens2.NewHashRefreshTokenGenerator(bcryptHashProvider)
	sessionManager := tokens2.NewTokenManager(
		tokens2.TokenConfiguration{
			AccessTokenDuration:  time.Second * 30, //TODO minutes
			RefreshTokenDuration: time.Hour * 24 * 30,
			Issuer:               "TODO",
		},
		accessTokenProvider,
		refreshTokenGenerator)

	smtpClient := smtp.New(cfg.SMTP.Username, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)
	smtpMailer := mailers2.NewSmtpMailer(smtpClient)
	verificationMailer := mailers2.NewVerificationMailer(smtpMailer)

	// Repository

	accountRepository := repositories.NewAccountRepository(
		selectAccountByIdCommand,
		selectAccountByEmailCommand,
		selectAccountByLoginCommand,
		updateAccountByIdCommand,
		insertAccountCommand)
	sessionRepository := repositories.NewSessionRepository(
		selectSessionByIdCommand,
		selectSessionByAccessTokenCommand,
		selectSessionsByAccountIdCommand,
		insertSessionCommand,
		updateSessionByIdCommand)
	verificationRepo := repositories.NewVerificationRepo(verificationDS)

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

	checkVerificationUseCase := usecases.NewCheckVerificationUsecase(
		accountRepository)
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

	http.Start(router, cfg.HTTP)
}
