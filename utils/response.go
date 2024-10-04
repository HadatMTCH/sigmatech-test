package utils

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	contentType              = "Content-Type"
	contentTypeValue         = "application/json; charset=utf-8"
	xContentTypeOptions      = "X-Content-Type-Options"
	xContentTypeOptionsValue = "nosniff"

	STATUS_INTERNAL_ERR = "STATUS_INTERNAL_ERROR"
	STATUS_BAD_REQUEST  = "STATUS_BAD_REQUEST"
	STATUS_UNAUTHORIZED = "STATUS_UNAUTHORIZED"
	STATUS_FORBIDDEN    = "STATUS_FORBIDDEN"
	STATUS_NOT_FOUND    = "STATUS_NOT_FOUND"
)

// Response is object for response http
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
	Msg  string      `json:"message"`
}

func SetResponseJSON(code int, data interface{}, err string, msg string) *Response {
	return &Response{
		Code: code,
		Data: data,
		Err:  err,
		Msg:  msg,
	}
}

// JSONResponseWithErr is method for return http json err
func (r *Response) JSONResponseWithErr(w http.ResponseWriter) {
	w.Header().Set(contentType, contentTypeValue)
	w.Header().Set(xContentTypeOptions, xContentTypeOptionsValue)
	w.WriteHeader(r.Code)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		logrus.Error(err)
	}
}

// JSONResponseWithErr is method for return http json
func (r *Response) JSONResponse(w http.ResponseWriter) {
	w.Header().Set(contentType, contentTypeValue)
	w.Header().Set(xContentTypeOptions, xContentTypeOptionsValue)
	w.WriteHeader(r.Code)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		logrus.Error(err)
	}
}
