package response

import (
	"net/http"
)

type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool        `json:"success" example:"false"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
	}
}
