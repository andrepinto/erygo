// Code generated by erygo
package messages

import (
	"bytes"
	"text/template"
)

func renderTemplate(templText string) string {
	buf := &bytes.Buffer{}
	templ, err := template.New("").Parse(templText)
	if err != nil {
		return err.Error()
	}
	err = templ.Execute(buf, map[string]string{
		"env":     "dev",
		"release": "v1.0.0",
	})
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func checkStatus(status int) bool {
	return status < 400
}
