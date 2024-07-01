package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"smartri_app/config"
	"smartri_app/internal/controllers/http"
	v1 "smartri_app/internal/controllers/http/v1"
	"smartri_app/internal/infrastructure/datasources/pg/commands/answers"
	"smartri_app/internal/infrastructure/datasources/pg/commands/questions"
	"smartri_app/internal/infrastructure/datasources/pg/commands/skills"
	"smartri_app/internal/infrastructure/datasources/pg/commands/test"
	"smartri_app/internal/infrastructure/datasources/pg/commands/user_data"
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

	selectQuestionsCommand := questions.NewSelectAllQuestionsCommand(pgClient)
	selectAnswerByIdCommand := answers.NewSelectAnswerByIdCommand(pgClient)
	selectAnswersByQuestionIdCommand := answers.NewSelectAnswersByQuestionIdCommand(pgClient)
	selectAnswerWithValuesCommand := answers.NewSelectAnswerValuesByAnswerIdCommand(pgClient)
	insertUserTestResultsCommand := test.NewInsertUserTestResultsCommand(pgClient)
	checkIfUserHasAnswersByAccountIdCommand := user_data.NewSelectUserHasAnswersByAccountIdCommand(pgClient)
	applyUserXpChangeByAccoundIdCommand := user_data.NewApplySkillChangesByAccountId(pgClient)

	selectAllSkillsCommand := skills.NewSelectAllSkillsCommand(pgClient)
	selectAllSkillsByAccountIdCommand := skills.NewSelectSkillsByAccountIdCommand(pgClient)
	selectNormalizationsBySkillIdCommand := skills.NewSelectSkillNormalizationBySkillIdCommand(pgClient)

	insertUserDataCommand := user_data.NewInsertUserDataCommand(pgClient, selectAllSkillsCommand)
	selectUserDataByAccountIdCommand := user_data.NewSelectUserDataByAccountId(pgClient)
	updateUserDataByAccountIdCommand := user_data.NewUpdateUserDataByAccountIdCommand(pgClient, selectUserDataByAccountIdCommand)

	// Repository

	testRepo := repositories.NewTestRepository(
		selectQuestionsCommand,
		selectAnswerByIdCommand,
		selectAnswersByQuestionIdCommand,
		selectAnswerWithValuesCommand,
		insertUserTestResultsCommand)
	skillRepo := repositories.NewSkillRepository(selectAllSkillsCommand, selectAllSkillsByAccountIdCommand, selectNormalizationsBySkillIdCommand)
	userDataRepo := repositories.NewUserRepository(selectUserDataByAccountIdCommand, updateUserDataByAccountIdCommand,
		insertUserDataCommand, checkIfUserHasAnswersByAccountIdCommand, applyUserXpChangeByAccoundIdCommand)

	// UseCases

	addUserDataUseCase := usecases.NewAddOrUpdateUserDataUseCase(userDataRepo)
	addUserAnswersUseCase := usecases.NewAddUserAnswers(testRepo, skillRepo, userDataRepo)
	addUserXpChangeUseCase := usecases.NewAddUserXpChange(skillRepo, userDataRepo)
	checkIfUserHasPassedTestYet := usecases.NewCheckUserHasPassedTestYetUseCase(userDataRepo)

	// Controllers

	getTestController := v1.NewGetTestController(testRepo)
	getUserSkillsController := v1.NewGetSkillDataController(skillRepo)

	addUserXpChangeController := v1.NewAddUserXpController(addUserXpChangeUseCase)
	addUserDataController := v1.NewAddUserDataController(addUserDataUseCase)
	addUserAnswersController := v1.NewAddUserAnswersController(addUserAnswersUseCase)
	checkIfUserHasPassedTestYetController := v1.NewCheckIfUserHasPassedTestYetController(checkIfUserHasPassedTestYet)

	router := gin.New()
	router.HandleMethodNotAllowed = true

	http.InitServiceMiddleware(router, logger)
	router.GET("/test", getTestController.GetQuestions)
	router.GET("/test/passed", checkIfUserHasPassedTestYetController.Check)

	router.GET("/user/skills", getUserSkillsController.GetUserSkills)
	router.POST("/user/xp", addUserXpChangeController.Add)
	router.POST("/user/data", addUserDataController.AddUserData)
	router.POST("/user/test", addUserAnswersController.AddUserAnswers)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	http2.Start(router, cfg.HTTP)
}
