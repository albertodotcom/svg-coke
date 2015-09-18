package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/opesun/goquery"
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

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

func fetchIcons(srcFolder string) []string {
	if !exists(srcFolder) {
		panic(fmt.Sprintf("'%s' folder doesn't exist", srcFolder))
	}

	files, _ := ioutil.ReadDir(srcFolder)

	var svgFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".svg") {
			svgFiles = append(svgFiles, path.Join(srcFolder, file.Name()))
		}
	}

	return svgFiles
}

func getSvg(svgBytes []byte) string {
	svg := bytes.NewReader(svgBytes)
	doc, err := goquery.Parse(svg)
	if err != nil {
		log.Fatal(err)
	}

	icon := doc.Find("symbol").OuterHtml()

	if icon == "" {
		panic("I haven't found any <symbol> in the svg")
	}

	return icon
}

func extractSvgContent(files []string) string {
	var svgIcons string

	for _, file := range files {

		dat, _ := ioutil.ReadFile(file)

		svgIcons = svgIcons + "\n" + getSvg(dat)
	}

	return svgIcons
}
