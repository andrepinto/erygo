package project

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/dave/jennifer/jen"
)

func sanitizeCommentLine(commentLine string) string {
	commentLine = strings.TrimSpace(commentLine)
	commentLine = strings.TrimPrefix(commentLine, "//")
	commentLine = strings.TrimPrefix(commentLine, "/*")
	commentLine = strings.TrimSuffix(commentLine, "*/")
	commentLine = strings.Replace(commentLine, "\n", "", -1)
	return commentLine
}
func buildComments(text string) *jen.Statement {
	var comments *jen.Statement
	i := 0
	scanner := bufio.NewScanner(bytes.NewReader([]byte(text)))
	for scanner.Scan() {
		commentLine := sanitizeCommentLine(scanner.Text())
		if commentLine == "" {
			continue
		}
		if i == 0 {
			i++
			comments = jen.Comment(commentLine)
			continue
		}
		comments.Add(jen.Line().Comment(commentLine))
	}
	return comments.Line()
}
