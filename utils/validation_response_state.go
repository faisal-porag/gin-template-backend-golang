package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type commonValidationResponse struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Lang    string      `json:"lang"`
	Data    interface{} `json:"data"`
}

func validateResponseMessageCommon(strMsgEn, lang string) commonValidationResponse {
	var commonMsg string
	if lang == "en" {
		commonMsg = strMsgEn
	} else {
		commonMsg = strMsgEn
	}
	errorResponse := commonValidationResponse{
		Code:    "SERVICE_ERROR",
		Success: false,
		Message: commonMsg,
		Lang:    lang,
	}

	return errorResponse
}

func ShowCommonValidationResponse(c *gin.Context, strMsgEn, lang string) {
	errorResponse := validateResponseMessageCommon(
		strMsgEn,
		lang,
	)
	c.JSON(http.StatusBadRequest, errorResponse)
	return
}

func ShowCommonValidationResponseWithStatusCode(c *gin.Context, strMsgEn, lang string, status int) {
	errorResponse := validateResponseMessageCommon(
		strMsgEn,
		lang,
	)
	c.JSON(status, errorResponse)
	return
}

func ShowAuthValidationResponse(c *gin.Context, strMsgEn, lang string) {
	errorResponse := commonValidationResponse{
		Code:    "AUTH_ERROR_401",
		Message: strMsgEn,
		Lang:    lang,
	}

	c.JSON(http.StatusUnauthorized, errorResponse)
}

type commonValidationResponseWithKeyValuePair struct {
	Code             string              `json:"code"`
	Success          bool                `json:"success"`
	Message          string              `json:"message"`
	Lang             string              `json:"lang"`
	ValidationErrors map[string][]string `json:"validationErrors"`
}

func validateResponseMessageCommonWithKeyValue(strMsgEn, lang string, vErrors map[string][]string) commonValidationResponseWithKeyValuePair {
	var commonMsg string
	if lang == "en" {
		commonMsg = strMsgEn
	} else {
		commonMsg = strMsgEn
	}
	errorResponse := commonValidationResponseWithKeyValuePair{
		Code:             "DATA_VALIDATION_FAILED",
		Success:          false,
		Message:          commonMsg,
		Lang:             lang,
		ValidationErrors: vErrors,
	}

	return errorResponse
}

func ShowCommonValidationResponseWithKeyValue(c *gin.Context, strMsgEn, lang string, vErrors map[string][]string) {
	errorResponse := validateResponseMessageCommonWithKeyValue(
		strMsgEn,
		lang,
		vErrors,
	)
	c.JSON(http.StatusBadRequest, errorResponse)
	return
}
