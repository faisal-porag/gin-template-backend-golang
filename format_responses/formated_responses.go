package format_responses

import (
	"github.com/faisal-porag/gin-template-backend-golang/utils"
	"net/http"
)

var UnauthorizedResponseData = utils.ResponseState{
	StatusCode: http.StatusUnauthorized,
	Code:       "UNAUTHORIZED_REQUEST",
	MessageEn:  "You are not authorized to perform this request. You have to login first.",
}

var ResponseSuccess = utils.ResponseState{
	StatusCode: http.StatusOK,
	Code:       "REQUEST_SUCCESS",
	MessageEn:  "Your request is successful.",
}
