package src

import (
	"authorization/core/external/databases"
	"authorization/core/external/databases/postgresql"
	"authorization/core/infrastructure/controllers/authentication"
	"authorization/core/infrastructure/repositories"
	"authorization/core/infrastructure/repositories/abstraction"
	"authorization/utils/errorsAndPanics"
	"fmt"
	"net/http"
)

type Server struct {
	dbClient databases.Database
}

func (s Server) Start() {
	dbConfig := postgresql.NewConfig()
	dbClient := postgresql.NewPostgresClient(dbConfig)
	err := dbClient.Connect()
	errorsAndPanics.HandleError(err)
	s.dbClient = dbClient

	repositoryCore := dbClient.RepositoryCore()
	var userRespository abstraction.UserRepository = repositories.NewUserRepository(repositoryCore)
	http.Handle("/word", authentication.NewSignUpController(&userRespository))
	fmt.Println("all good")
	http.ListenAndServe(":8080", nil)
}

func (s Server) Stop() {
	s.dbClient.Disconnect()
}
