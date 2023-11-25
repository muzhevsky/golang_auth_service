package authentication

import (
	"authorization/core/domain/dtos"
	"authorization/core/domain/useCases"
	"authorization/utils/errorHandling"
	"encoding/json"
	"net/http"
)

type signUpController struct {
	useCases.AuthenticationService
}

func NewSignUpController(useCase useCases.AuthenticationService) *signUpController {
	return &signUpController{useCase}
}

func (controller *signUpController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var dto dtos.SignUpDto
	err := json.NewDecoder(request.Body).Decode(&dto)
	errorHandling.LogError(err)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.AuthenticationService.SignUp(&dto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handleBadRequest() {

}
