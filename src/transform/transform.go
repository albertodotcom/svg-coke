package transform

import (
	"bytes"
	"text/template"
)

const svgSpriteTemplate = "../../svgSpriteTemplate.svg"

func openTemplate() *template.Template {
	tmpl, err := template.New("svgSpriteTemplate").ParseFiles(svgSpriteTemplate)

	if err != nil {
		panic(err)
	}

	return tmpl
}

type templateContext struct {
	StringToInject string
}

func injectIconsIntoSvgTemplate(icons string, tmpl *template.Template) bytes.Buffer {
	tc := templateContext{icons}

	var output bytes.Buffer
	tmpl.Execute(&output, tc)

	return output
}
