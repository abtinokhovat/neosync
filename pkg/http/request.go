package http

import (
	"bytes"
	"encoding/json"
	"io"
	"neosync/pkg/richerror"
	"net/http"
)

func request(method string, op richerror.Op, url string, requestBody any, headers ...Header) (Response[[]byte], error) {
	res := Response[[]byte]{}

	var bodyReader io.Reader
	if method != http.MethodGet {
		// marshaling json and pass it as body
		jsonBody, mErr := json.Marshal(requestBody)
		if mErr != nil {
			return res, richerror.New(op).WithErr(mErr).WithKind(richerror.KindUnexpected)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	client := &http.Client{}
	req, rErr := http.NewRequest(method, url, bodyReader)
	if rErr != nil {
		//tracing.SpanError(span, rErr)
		return res, richerror.New(op).WithErr(rErr).WithKind(richerror.KindUnexpected)
	}

	// adding headers to the request
	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	// executing http request
	response, dErr := client.Do(req)
	if dErr != nil {
		//tracing.SpanError(span, dErr)
		return res, richerror.New(op).WithErr(dErr).WithKind(richerror.KindUnexpected)
	}
	defer response.Body.Close()

	// before here errors are for the whole request process
	res.Code = response.StatusCode

	// error handling logic
	if response.StatusCode != http.StatusOK {
		kind := MapHTTPStatusCodeToKind(response.StatusCode)
		errorMessage := "some thing went wrong"

		requestError := richerror.New(op).WithMessage(errorMessage).WithKind(kind).WithMeta(map[string]interface{}{
			"code": response.StatusCode,
			"url":  url,
		})

		res.Message = requestError.Error()
		return res, nil
	}

	body, rErr := io.ReadAll(response.Body)
	if rErr != nil {
		return res, richerror.New(op).WithErr(rErr).WithKind(richerror.KindUnexpected)
	}

	res.Data = body
	return res, nil
}
