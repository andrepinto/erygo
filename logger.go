package erygo

//Logger ...
type Logger interface {
	LogErr(err error, erygoErr *Err)
	LogResp(msg string, erygoErr *Response)
}
