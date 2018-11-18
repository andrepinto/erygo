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
