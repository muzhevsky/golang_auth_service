package controllers

import (
	"encoding/json"
	"golang-app/src/services/core/entities"
	"golang-app/src/services/core/repositories/abstraction"
	"golang-app/utils/errorsAndPanics"
	"io"
	"net/http"
)

type translationController struct {
	repository *abstraction.TranslationRepository
}

func NewTranslationController(repository *abstraction.TranslationRepository) *translationController {
	return &translationController{repository}
}

func (controller *translationController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		controller.get(writer, request)
		break
	case http.MethodPost:
		controller.post(writer, request)
		break
	}
}

func (controller *translationController) get(writer http.ResponseWriter, request *http.Request) {
	repository := *controller.repository
	urlParams := request.URL.Query()
	result := repository.SelectByWord(urlParams.Get("notation"))
	jsonResult, err := json.Marshal(result)
	errorsAndPanics.HandleError(err)
	writer.Write(jsonResult)
}

func (controller *translationController) post(writer http.ResponseWriter, request *http.Request) {
	repository := *controller.repository
	body, err := io.ReadAll(request.Body)
	errorsAndPanics.HandleError(err)
	word := &(entities.Word{
		Notation:     "",
		Translations: nil,
		UseCases:     nil,
	})
	err = json.Unmarshal(body, word)
	errorsAndPanics.HandleError(err)

	repository.Insert(word)
}
