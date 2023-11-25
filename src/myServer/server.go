package myServer

import (
	"fmt"
	"golang-app/src/data/databases"
	"golang-app/src/data/databases/mongo"
	"golang-app/src/myServer/controllers"
	"golang-app/src/services/core/repositories"
	"golang-app/src/services/core/repositories/abstraction"
	"golang-app/utils/errorsAndPanics"
	"net/http"
)

type Server struct {
	dbClient databases.Database
}

func (s Server) Start() {
	dbConfig := databases.NewConfig()
	dbClient := mongo.NewMongoClient(dbConfig)
	err := dbClient.Connect()
	errorsAndPanics.HandleError(err)
	s.dbClient = dbClient

	repositoryCore := dbClient.RepositoryCore()
	var translationRepository abstraction.TranslationRepository = repositories.NewTranslationRepository(repositoryCore)
	http.Handle("/word", controllers.NewTranslationController(&translationRepository))
	fmt.Println("all good")
	http.ListenAndServe(":8080", nil)
}

func (s Server) Stop() {
	s.dbClient.Disconnect()
}
