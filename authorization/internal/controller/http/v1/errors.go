package v1

import "github.com/gin-gonic/gin"

const DefaultErrorCode = "D-0001"
const ValidationErrorCode = "V-0001"
const RecordExistErrorCode = "R-0001"

type response struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func errorResponse(c *gin.Context, code int, msg string, developerCode string) { // не надо в msg передавать err.Error(), там всякие данные бывают
	c.AbortWithStatusJSON(code, response{msg, developerCode})
}
