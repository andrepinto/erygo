package erygo

import "fmt"

//Info ...
type Info struct {
	Service string `json:"sid"`
	Kind    int    `json:"kind"`
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
