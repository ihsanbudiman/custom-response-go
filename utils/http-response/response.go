package httpresponse

import (
	"encoding/json"
	error_utils "kocag/utils/error"
	"net/http"
)

type IResponseWriter interface {
	HttpError(w http.ResponseWriter, err error)
	JSON(w http.ResponseWriter, status int, v interface{}) error
}

type ResponseWriter struct {
}

// HttpError implements IResponseWriter.
func (r *ResponseWriter) HttpError(w http.ResponseWriter, err error) {
	customErr, ok := err.(*error_utils.CustomError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customErr.GetStatus())
		w.Write([]byte(customErr.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// JSON implements IResponseWriter.
func (r *ResponseWriter) JSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func NewResponseWriter() IResponseWriter {
	return &ResponseWriter{}
}
