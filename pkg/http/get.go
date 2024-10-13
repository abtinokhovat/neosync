package http

import (
	"encoding/json"
	"fmt"
	"neosync/pkg/richerror"
	"net/http"
)

type Header struct {
	Key   string
	Value string
}

type Response[D any] struct {
	Data    D
	Code    int
	Message string
}

func Get[D any](op richerror.Op, url string, headers ...Header) (Response[D], error) {
	const method = "GET"

	responseBody, err := request(method, op, url, nil, headers...)
	// check for internal errors
	if err != nil {
		return Response[D]{}, err
	}

	response := Response[D]{
		Code:    responseBody.Code,
		Message: responseBody.Message,
	}

	// check for external http service errors
	if responseBody.Code != http.StatusOK {
		return response, richerror.New(op).WithMessage(fmt.Sprintf("external api failed with code %d, message: %s", responseBody.Code, responseBody.Message))
	}

	// marshaling request as json
	if uErr := json.Unmarshal(responseBody.Data, &response.Data); uErr != nil {
		return response, richerror.New(op).WithErr(uErr).WithKind(richerror.KindUnexpected)
	}

	return response, nil
}
