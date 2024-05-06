package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseState struct {
	StatusCode int
	Code       string
	Success    bool
	MessageEn  string
}

type successResponseForApplication struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Lang    string      `json:"lang"`
	Data    interface{} `json:"data"`
}

// ShowCommonSuccessResponse generates a success JSON response.
func ShowCommonSuccessResponse(c *gin.Context, message string, data interface{}, lang string) {
	var commonMsg string
	if lang == "en" {
		commonMsg = message
	} else {
		commonMsg = message
	}

	response := &successResponseForApplication{
		Code:    "SUCCESS",
		Success: true,
		Message: commonMsg,
		Lang:    lang,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
	return
}

// ===================================================================================================================

// Error Data Response

type errorResponseForApplication struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Lang    string      `json:"lang"`
	Data    interface{} `json:"data"`
}

func ShowCommonErrorResponse(c *gin.Context, statusCode int, errorMessage string, lang string) {
	var commonMsg string
	if lang == "en" {
		commonMsg = errorMessage
	} else {
		commonMsg = errorMessage
	}

	response := errorResponseForApplication{
		Code:    "SERVICE_ERROR",
		Success: false,
		Message: commonMsg,
		Lang:    lang,
		Data:    nil,
	}

	c.JSON(statusCode, response)
	return
}

// ========================================= Common Responses ==================================

type CommonResponseForApplication struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Lang    string      `json:"lang"`
	Data    interface{} `json:"data"`
}

func (rs ResponseState) CommonResponse(lang string, data interface{}, isSuccess bool) CommonResponseForApplication {
	var message string
	if lang == "en" {
		message = rs.MessageEn
	} else {
		message = rs.MessageEn
	}

	return CommonResponseForApplication{
		Code:    rs.Code,
		Success: isSuccess,
		Message: message,
		Lang:    lang,
		Data:    data,
	}
}

func (rs ResponseState) WriteToResponse(c *gin.Context, lang string, data interface{}, isSuccess bool) {
	c.JSON(rs.StatusCode, rs.CommonResponse(lang, data, isSuccess))
}
