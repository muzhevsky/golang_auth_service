package errorHandling

import (
	"fmt"
	"net/http"
)

func LogError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func SetHttpHeader(err error, w http.ResponseWriter, responseStatus int) {
	LogError(err)
	if err != nil {
		w.WriteHeader(responseStatus)
	}
}
