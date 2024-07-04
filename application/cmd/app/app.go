package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"smartri_app/config"
	"smartri_app/controllers/http"
	v12 "smartri_app/controllers/http/v1"
	"smartri_app/internal/infrastructure/datasources/pg/commands/answers"
	"smartri_app/internal/infrastructure/datasources/pg/commands/avatars"
	"smartri_app/internal/infrastructure/datasources/pg/commands/questions"
	"smartri_app/internal/infrastructure/datasources/pg/commands/skill_changes"
	"smartri_app/internal/infrastructure/datasources/pg/commands/skills"
	"smartri_app/internal/infrastructure/datasources/pg/commands/user_answers"
	"smartri_app/internal/infrastructure/datasources/pg/commands/user_data"
	"smartri_app/internal/infrastructure/datasources/pg/commands/user_skills"
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

	selectQuestionsCommand := questions.NewSelectAllQuestionsPGCommand(pgClient)
	selectAnswerByIdCommand := answers.NewSelectAnswerByIdPGCommand(pgClient)
	selectAnswersByQuestionIdCommand := answers.NewSelectAnswersByQuestionIdPGCommand(pgClient)
	selectAnswerWithValuesCommand := answers.NewSelectAnswerValuesByAnswerIdPGCommand(pgClient)
	insertUserTestResultsCommand := user_answers.NewInsertUserTestResultsPGCommand(pgClient)
	checkIfUserHasAnswersByAccountIdCommand := user_answers.NewSelectUserHasAnswersByAccountIdPGCommand(pgClient)
	applyUserXpChangeByAccountIdCommand := skill_changes.NewApplySkillChangesByAccountIdPGCommand(pgClient)

	selectAllSkillsCommand := skills.NewSelectAllSkillsPGCommand(pgClient)
	selectAllSkillsByAccountIdCommand := user_skills.NewSelectSkillsByAccountIdPGCommand(pgClient)
	selectNormalizationsBySkillIdCommand := skills.NewSelectSkillNormalizationBySkillIdPGCommand(pgClient)

	insertUserDataCommand := user_data.NewInsertUserDataPGCommand(pgClient, selectAllSkillsCommand)
	selectUserDataByAccountIdCommand := user_data.NewSelectUserDataByAccountIdPGCommand(pgClient)
	updateUserDataByAccountIdCommand := user_data.NewUpdateUserDataByAccountIdPGCommand(pgClient, selectUserDataByAccountIdCommand)

	selectAvatarCommand := avatars.NewSelectAvatarByAccountIdPGCommand(pgClient)
	insertAvatarCommand := avatars.NewInsertAvatarPGCommand(pgClient)
	updateAvatarCommand := avatars.NewUpdateAvatarByAccountIdPGCommand(pgClient)
	// Repository

	testRepo := repositories.NewTestRepository(
		selectQuestionsCommand,
		selectAnswerByIdCommand,
		selectAnswersByQuestionIdCommand,
		selectAnswerWithValuesCommand,
		insertUserTestResultsCommand)
	skillRepo := repositories.NewSkillRepository(selectAllSkillsCommand, selectAllSkillsByAccountIdCommand, selectNormalizationsBySkillIdCommand)
	avatarRepo := repositories.NewAvatarRepository(selectAvatarCommand, insertAvatarCommand, updateAvatarCommand)
	userDataRepo := repositories.NewUserDataRepository(selectUserDataByAccountIdCommand, updateUserDataByAccountIdCommand, insertUserDataCommand)
	userSkillsRepo := repositories.NewUserSkillRepository(applyUserXpChangeByAccountIdCommand)
	userAnswersRepo := repositories.NewUserTestAnswersRepository(checkIfUserHasAnswersByAccountIdCommand)

	// UseCases

	addUserDataUseCase := usecases.NewAddOrUpdateUserDataUseCase(userDataRepo)
	addUserAnswersUseCase := usecases.NewAddUserAnswers(testRepo, skillRepo, userDataRepo, userAnswersRepo)
	addUserXpChangeUseCase := usecases.NewAddUserXpChange(skillRepo, userDataRepo, userSkillsRepo)

	getUserAvatarUseCase := usecases.NewGetUserAvatarUseCase(avatarRepo)
	initOrUpdateAvatarUseCase := usecases.NewInitOrUpdateAvatarUseCase(avatarRepo)

	checkIfUserHasPassedTestYetUseCase := usecases.NewCheckUserHasPassedTestYetUseCase(userAnswersRepo)

	// Controllers

	getTestController := v12.NewGetTestController(testRepo)
	getUserSkillsController := v12.NewGetSkillDataController(skillRepo)

	addUserXpChangeController := v12.NewAddUserXpController(addUserXpChangeUseCase)
	addUserDataController := v12.NewAddUserDataController(addUserDataUseCase)
	addUserAnswersController := v12.NewAddUserAnswersController(addUserAnswersUseCase)
	checkIfUserHasPassedTestYetController := v12.NewCheckIfUserHasPassedTestYetController(checkIfUserHasPassedTestYetUseCase)

	getUserAvatarController := v12.NewGetUserAvatarController(getUserAvatarUseCase)
	initOrUpdateAvatarController := v12.NewInitOrUpdateUserAvatarController(initOrUpdateAvatarUseCase)

	router := gin.New()
	router.HandleMethodNotAllowed = true

	http.InitServiceMiddleware(router, logger)
	router.GET("/test", getTestController.GetQuestions)
	router.GET("/test/passed", checkIfUserHasPassedTestYetController.CheckIfUserHasPassedTestYet)

	router.GET("/user/skills", getUserSkillsController.GetUserSkills)
	router.POST("/user/xp", addUserXpChangeController.AddUserXp)
	router.POST("/user/data", addUserDataController.AddUserData)
	router.POST("/user/test", addUserAnswersController.AddUserAnswers)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/user/avatar", getUserAvatarController.GetUserAvatar)
	router.POST("/user/avatar", initOrUpdateAvatarController.InitOrUpdateAvatar)
	http2.Start(router, cfg.HTTP)
}
