package internal

import (
	"authorization/internal/controllers/http/authentication"
	"authorization/internal/domain/infrastructure/lib/encryption"
	"authorization/internal/domain/infrastructure/repositories/postgresql"
	"authorization/internal/domain/useCases"
	"authorization/package/databases"
	pgPackage "authorization/package/databases/postgres"
	"authorization/utils/errorHandling"
	"fmt"
	"net/http"
)

type Server struct {
	dbClient databases.ISqlClient
}

func (s Server) Start() {
	config := NewConfig()
	dbConfig := config.DatabaseConfig
	connectionConfig := dbConfig.Connection
	connectionString := pgPackage.ConnectionString{}
	connectionString.SetUser(connectionConfig.User).SetPassword(connectionConfig.Password).SetHostname(connectionConfig.Hostname).SetPort(connectionConfig.Port).SetDatabaseName(connectionConfig.DatabaseName)
	dbClient := pgPackage.NewDatabase(&connectionString)
	err := dbClient.Connect()
	if err != nil {
		errorHandling.LogError(err)
		return
	}

	s.dbClient = dbClient

	connection := dbClient.GetConnection()
	userRespository := postgresql.NewUserRepository(connection)
	authenticationUsecase := InitAuthenticationUsecase(userRespository)

	http.Handle("/signup", authentication.NewSignUpController(authenticationUsecase))
	fmt.Println("all good")
	http.ListenAndServe(":8080", nil)
}

func InitAuthenticationUsecase(repository useCases.UserRepository) useCases.AuthenticationUseCase {
	var encryptionService = encryption.New(10)
	return useCases.NewAuthenticationService(repository, encryptionService)
}

func (s Server) Stop() {
	err := s.dbClient.Disconnect()
	if err != nil {
		return
	}
}
