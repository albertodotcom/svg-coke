package main

import (
	"bytes"
	"text/template"
)

const svgSpriteTemplate = "./svgSpriteTemplate.svg"

func openTemplate() *template.Template {
	tmpl, _ := template.New("svgSpriteTemplate").ParseFiles(svgSpriteTemplate)

	return tmpl
}

func injectIconsIntoSvgTemplate(icons string, tmpl *template.Template) bytes.Buffer {
	tc := struct {
		Icons string
	}{
		icons,
	}

	var output bytes.Buffer
	tmpl.Execute(&output, tc)

	return output
}
