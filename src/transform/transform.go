package transform

import (
	"bytes"
	"text/template"
)

type templateContext struct {
	StringToInject string
}

func injectStringIntoTemplate(stringToInject string, tmpl *template.Template) bytes.Buffer {
	tc := templateContext{stringToInject}

	var output bytes.Buffer
	tmpl.Execute(&output, tc)

	return output
}
