package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseWithError struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Stack   string `json:"stack"`
	Code    string `json:"code"`
}

type jsonRes struct {
	w                http.ResponseWriter
	r                *http.Request
	err              error
	errorDescription string
	body             interface{}
	statusCode       int
}

func NewResp(w http.ResponseWriter, r *http.Request) *jsonRes {
	return &jsonRes{
		w: w,
		r: r,
	}
}

func (r *jsonRes) Error(err error) *jsonRes {
	r.err = err
	return r
}

func (r *jsonRes) Status(statusCode int) *jsonRes {
	r.statusCode = statusCode
	return r
}

func (r *jsonRes) Json(body interface{}) *jsonRes {
	r.body = body
	return r
}

func (r *jsonRes) Send() {
	r.w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.statusCode != 0 {
		r.w.WriteHeader(r.statusCode)
	}

	if r.err != nil {

		err := &ResponseWithError{Error: Error{
			Message: r.err.Error(),
			Stack:   "",
		}}

		res, sendErr := json.Marshal(err)
		if sendErr != nil {
			r.w.WriteHeader(http.StatusInternalServerError)
			_, _ = r.w.Write([]byte(sendErr.Error()))
			return
		}
		_, _ = r.w.Write(res)
		return
	}

	if r.body != nil {
		body, err := json.Marshal(r.body)
		if err != nil {
			r.w.WriteHeader(http.StatusInternalServerError)
			_, _ = r.w.Write([]byte(err.Error()))
			return
		}
		_, err = r.w.Write(body)
		if err != nil {
			r.w.WriteHeader(http.StatusInternalServerError)
			_, _ = r.w.Write([]byte(err.Error()))
			return
		}
	}
}
