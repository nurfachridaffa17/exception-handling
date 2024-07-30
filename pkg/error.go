package pkg

import (
	"fmt"
	"strings"
)

type Meta struct {
	Success bool   `json:"success" default:"true"`
	Message string `json:"message" default:"true"`
}

type errorResponse struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_DUPLICATE_ENTITY     = "duplicate_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_VALIDATION           = "validation"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	Duplicate                Error
	NotFound                 Error
	RouteNotFound            Error
	UnprocessableEntity      Error
	DuplicateEntity          Error
	Unauthorized             Error
	BadRequest               Error
	Validation               Error
	InternalServerError      Error
	EmailOrPasswordIncorrect Error
}

var ErrorConstant errorConstant = errorConstant{
	Duplicate: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "duplicate",
			},
			Error: "duplicate",
		},
		Code: 409,
	},

	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "not found",
			},
			Error: "not found",
		},
		Code: 404,
	},

	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "unprocessable entity",
			},
			Error: "unprocessable entity",
		},
		Code: 422,
	},

	DuplicateEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "duplicate entity",
			},
			Error: "duplicate entity",
		},
		Code: 409,
	},

	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "unauthorized",
			},
			Error: "unauthorized",
		},
		Code: 401,
	},

	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "bad request",
			},
			Error: "bad request",
		},
		Code: 400,
	},

	Validation: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "validation error",
			},
			Error: "validation error",
		},
		Code: 422,
	},

	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "internal server error",
			},
			Error: "internal server error",
		},
		Code: 500,
	},

	EmailOrPasswordIncorrect: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "email or password incorrect",
			},
			Error: "email or password incorrect",
		},
		Code: 401,
	},
}

func ErrorBuilder(res *Error, message error, customMessage ...string) *Error {
	res.ErrorMessage = message

	if strings.Contains(strings.Join([]string{E_VALIDATION, E_BAD_REQUEST, E_DUPLICATE}, ","), res.Response.Error) {
		res.Response.Meta.Message = message.Error()
	}
	if len(customMessage) > 0 {
		res.Response.Meta.Message = ""
		for i := range customMessage {
			res.Response.Meta.Message += customMessage[i]
		}
	}
	return res
}

func CustomErrorBuilder(code int, err string, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: err,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}

func (e *Error) ParseToError() error {
	return e
}
