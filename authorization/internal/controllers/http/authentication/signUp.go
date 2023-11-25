package authentication

import (
	"authorization/internal/domain/entities"
	"authorization/internal/domain/useCases"
	"authorization/utils/errorHandling"
	"encoding/json"
	"net/http"
)

type signUpController struct {
	useCase useCases.AuthenticationUseCase
}

func NewSignUpController(useCase useCases.AuthenticationUseCase) *signUpController {
	return &signUpController{useCase}
}

func (controller *signUpController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dto := entities.User{}
	err := json.NewDecoder(request.Body).Decode(&dto)
	errorHandling.LogError(err)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.useCase.SignUp(&dto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handleBadRequest() {

}
