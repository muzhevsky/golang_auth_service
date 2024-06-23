package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"smartri_app/config"
	"smartri_app/internal/controllers/http"
	v1 "smartri_app/internal/controllers/http/v1"
	pg2 "smartri_app/internal/infrastructure/datasources/pg"
	"smartri_app/internal/repositories"
	"smartri_app/internal/usecases"
	http2 "smartri_app/pkg/http"
	"smartri_app/pkg/logger"
	"smartri_app/pkg/postgres"
	"strconv"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run 1 : %w", err))
	}

	// Packages
	logger := logger.New("error")

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

	questionsDS := pg2.NewQuestionsDataSource(pgClient)
	answersDS := pg2.NewAnswersDataSource(pgClient)
	userAnswersDS := pg2.NewUserTestResultsSource(pgClient)

	userDataDS := pg2.NewUsersDataSource(pgClient)

	// Repository

	testRepo := repositories.NewTestRepository(questionsDS, answersDS, nil, userAnswersDS)
	userDataRepo := repositories.NewUserRepository(userDataDS)

	// UseCases

	addUserDataUseCase := usecases.NewAddUserData(userDataRepo)
	addUserAnswersUseCase := usecases.NewAddUserAnswers(testRepo)

	// Controllers

	getTestController := v1.NewGetTestController(testRepo)
	addUserDataController := v1.NewAddUserDataController(addUserDataUseCase)
	addUserAnswersController := v1.NewAddUserAnswersController(addUserAnswersUseCase)

	router := gin.New()
	router.HandleMethodNotAllowed = true

	http.InitServiceMiddleware(router, logger)
	router.GET("/test", getTestController.GetQuestions)
	router.POST("/user/data", addUserDataController.AddUserData)
	router.POST("/user/test", addUserAnswersController.AddUserAnswers)
	http2.Start(router, cfg.HTTP)
}
