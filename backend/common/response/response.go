package response

import (
	"encoding/json"
	"log"
)

type ErrorCode int64

const (
	UNDEFINED_ERROR ErrorCode = iota
	VALIDATION_ERROR
	AUTHENTICATION_ERROR
	AUTHORIZATION_ERROR
	SERVER_ERROR
	DB_RESOURCE_ERROR
	TRANSACTION_ERROR
)

func (e ErrorCode) String() string {
	switch e {
	case VALIDATION_ERROR:
		return "VALIDATION_ERROR"

	case AUTHENTICATION_ERROR:
		return "AUTHENTICATION_ERROR"

	case SERVER_ERROR:
		return "SERVER_ERROR"

	case DB_RESOURCE_ERROR:
		return "RESOURCE_ERROR"

	case TRANSACTION_ERROR:
		return "TRANSACTION_ERROR"
	case AUTHORIZATION_ERROR:
		return "AUTHORIZATION_ERROR"
	}

	return "UNDEFINED_ERROR"
}

type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type Error struct {
	Code    string `json:"code"`
	Details string `json:"details"`
}

func New() Response {
	return Response{
		Error: true,
		Data: Error{
			Code:    UNDEFINED_ERROR.String(),
			Details: "Error is not defined.",
		},
	}
}

func (r *Response) SetError(code ErrorCode, err error) {
	new_err := Error{
		Code:    code.String(),
		Details: err.Error(),
	}

	r.Error = true
	r.Data = new_err
}

func (r *Response) SetData(data interface{}) {
	r.Error = false
	r.Data = data
}

func (r *Response) Msg() []byte {
	data, err := json.Marshal(r)

	if err != nil {
		log.Fatalln(err)
	}

	return data
}
