// Code generated by erygo
package messages

import erygo "github.com/andrepinto/erygo"

// UserNotFoud error
func UserNotFoud(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{"hello {{.env}}"},
		Info: erygo.Info{
			Kind:    1,
			Service: "users",
		},
		Message:    "user not found",
		StatusHTTP: 404,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// InternalServerError error
func InternalServerError(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{"hello {{.env}}"},
		Info: erygo.Info{
			Kind:    2,
			Service: "users",
		},
		Message:    "internal server error",
		StatusHTTP: 500,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}
