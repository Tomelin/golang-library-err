package golang-library-err

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func ErrorBadRequest(message string) *RestErr {
	return RestErr {
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func ErrorNotFound(message string) *RestErr {
	return RestErr {
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}
}

func ErrorInternalServerError(message string) *RestErr {
	return RestErr {
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}

