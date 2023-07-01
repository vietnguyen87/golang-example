package apiwrapper

import "github.com/vietnguyen87/pkg-golang/xerrors"

func SuccessResponse() *Response {
	return &Response{
		Error: xerrors.Success.New(),
	}
}

func SuccessWithDataResponse(data interface{}) *Response {
	return &Response{
		Error: xerrors.Success.New(),
		Data:  data,
	}
}

func FailureResponse(err error) *Response {
	return &Response{
		Error: err,
	}
}
func FailureWithDataResponse(err error, data interface{}) *Response {
	return &Response{
		Error: err,
		Data:  data,
	}
}
