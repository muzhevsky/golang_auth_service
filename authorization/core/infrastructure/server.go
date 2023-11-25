package infrastructure

import (
	"authorization/core/domain/repositories/abstraction"
	"authorization/core/domain/useCases"
	"authorization/core/external/databases"
	"authorization/core/external/databases/postgresql"
	"authorization/core/external/databases/postgresql/repositories"
	"authorization/core/infrastructure/controllers/authentication"
	"authorization/utils/errorHandling"
	"fmt"
	"net/http"
)

type Server struct {
	dbClient databases.Database
}

func (s Server) Start() {
	dbConfig := postgresql.NewConfig()
	dbClient := postgresql.NewDatabase(dbConfig)
	err := dbClient.Connect()
	if err != nil {
		errorHandling.LogError(err)
		return
	}

	s.dbClient = dbClient

	var userRespository = repositories.NewUserRepository(dbClient.GetConnection())
	var authenticationUsecase = InitAuthenticationUsecase(userRespository)

	http.Handle("/signup", authentication.NewSignUpController(authenticationUsecase))
	fmt.Println("all good")
	http.ListenAndServe(":8080", nil)
}

func InitAuthenticationUsecase(repository abstraction.UserRepository) useCases.AuthenticationService {
	var saltService = useCases.NewSaltService()
	var encryptionService = useCases.NewEncryptionService()
	return useCases.NewAuthenticationService(repository, saltService, encryptionService)
}

func (s Server) Stop() {
	s.dbClient.Disconnect()
}
