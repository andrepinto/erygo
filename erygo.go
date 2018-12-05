package erygo

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//Info ...
type Info struct {
	Service string `json:"service"`
	Kind    int    `json:"kind"`
}

func (info *Info) String() string {
	return fmt.Sprintf("%v-%v", info.Service, info.Kind)
}

type Labels map[string]string

//Err ...
type Err struct {
	Message    string   `json:"message"`
	StatusHTTP int      `json:"status_http"`
	Info       Info     `json:"info"`
	Details    []string `json:"details,omitempty"`
	Labels     Labels   `json:"labels,omitempty"`
}

//AddDetails ...
func (err *Err) AddDetails(details ...string) *Err {
	err.Details = append(err.Details, details...)
	return err
}

//AddDetailF ...
func (err *Err) AddDetailF(formatS string, args ...interface{}) *Err {
	return err.AddDetails(fmt.Sprintf(formatS, args...))
}

//AddDetailsErr ...
func (err *Err) AddDetailsErr(details ...error) *Err {
	for _, detail := range details {
		err.AddDetails(detail.Error())
	}
	return err
}

//WithLabel ...
func (err *Err) WithLabel(code, value string) *Err {
	if err.Labels == nil {
		err.Labels = make(Labels)
	}

	err.Labels[code] = value
	return err
}

//WithLabels ...
func (err *Err) WithLabels(labels Labels) *Err {
	if err.Labels == nil {
		err.Labels = make(Labels)
	}

	for code, value := range labels {
		err.Labels[code] = value
	}

	return err
}

//Log ...
func (err *Err) Log(extErr error, logger Logger) *Err {
	logger.LogErr(extErr, err)
	return err
}

func (err *Err) Error() string {
	buf := bytes.NewBufferString("[" + err.Info.String() + "] " +
		http.StatusText(err.StatusHTTP) + " " +
		err.Message)
	if len(err.Details) > 0 {
		buf.WriteString(": ")
		buf.WriteString(strings.Join(err.Details, "; "))
	}
	if len(err.Labels) > 0 {
		buf.WriteString(": ")
		var fields []string
		for name, value := range err.Labels {
			fields = append(fields, name+"="+strconv.QuoteToASCII(value))
		}
		buf.WriteString(strings.Join(fields, ", "))
	}
	return buf.String()
}

//Response ...
type Response struct {
	ID            string      `json:"id" `
	CorrelationID string      `json:"correlation_id" `
	Success       bool        `json:"success" `
	Message       string      `json:"message" `
	Result        interface{} `json:"result" `
	StatusHTTP    int         `json:"status_http"`
	Info          Info        `json:"info"`
	Details       []string    `json:"details,omitempty"`
	Labels        Labels      `json:"labels,omitempty"`
}

//AddResult ...
func (resp *Response) AddResult(result interface{}) *Response {
	resp.Result = result
	return resp
}

//AddDetails ...
func (resp *Response) AddDetails(details ...string) *Response {
	resp.Details = append(resp.Details, details...)
	return resp
}

//AddDetailF ...
func (resp *Response) AddDetailF(formatS string, args ...interface{}) *Response {
	return resp.AddDetails(fmt.Sprintf(formatS, args...))
}

//AddDetailsErr ...
func (resp *Response) AddDetailsErr(details ...error) *Response {
	for _, detail := range details {
		resp.AddDetails(detail.Error())
	}
	return resp
}

//WithLabel ...
func (resp *Response) WithLabel(code, value string) *Response {
	if resp.Labels == nil {
		resp.Labels = make(Labels)
	}

	resp.Labels[code] = value
	return resp
}

//WithLabels ...
func (resp *Response) WithLabels(labels Labels) *Response {
	if resp.Labels == nil {
		resp.Labels = make(Labels)
	}

	for code, value := range labels {
		resp.Labels[code] = value
	}

	return resp
}

//Log ...
func (resp *Response) Log(msg string, logger Logger) *Response {
	logger.LogResp(msg, resp)
	return resp
}

/*
* Constructor
 */

//ErrConstruct ...
type ErrConstruct func(...func(*Err)) *Err

//Error ...
func (constr ErrConstruct) Error() string {
	return constr().Error()
}

//AddDetails ...
func (constr ErrConstruct) AddDetails(details ...string) ErrConstruct {
	return func(options ...func(*Err)) *Err {
		err := constr().AddDetails(details...)
		for _, option := range options {
			option(err)
		}
		return err
	}
}

//AddDetailsErr ...
func (constr ErrConstruct) AddDetailsErr(details ...error) ErrConstruct {
	return func(options ...func(*Err)) *Err {
		err := constr().AddDetailsErr(details...)
		for _, option := range options {
			option(err)
		}
		return err
	}
}

//AddDetailF ...
func (constr ErrConstruct) AddDetailF(f string, vals ...interface{}) ErrConstruct {
	return func(options ...func(*Err)) *Err {
		err := constr().AddDetailF(f, vals...)
		for _, option := range options {
			option(err)
		}
		return err
	}
}

//WithLabel ...
func (constr ErrConstruct) WithLabel(key, value string) ErrConstruct {
	return func(options ...func(*Err)) *Err {
		err := constr().WithLabel(key, value)
		for _, option := range options {
			option(err)
		}
		return err
	}
}

//WithLabels ...
func (constr ErrConstruct) WithLabels(fields Labels) ErrConstruct {
	return func(options ...func(*Err)) *Err {
		err := constr().WithLabels(fields)
		for _, option := range options {
			option(err)
		}
		return err
	}
}
