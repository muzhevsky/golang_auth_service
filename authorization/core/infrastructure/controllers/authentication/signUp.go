package authentication

import (
	"authorization/core/infrastructure/repositories/abstraction"
	"net/http"
)

type signUpController struct {
	repository *abstraction.UserRepository
}

func NewSignUpController(repository *abstraction.UserRepository) *signUpController {
	return &signUpController{repository}
}

func (controller *signUpController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
