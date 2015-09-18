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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openTemplate() *template.Template {
	dat, err := ioutil.ReadFile(svgSpriteTemplate)
	check(err)

	tmpl, err := template.New("svgSpriteTemplate").Parse(string(dat))
	check(err)

	return tmpl
}

func injectIconsIntoSvgTemplate(icons string, tmpl *template.Template) []byte {
	tc := struct {
		Icons string
	}{
		icons,
	}

	var output bytes.Buffer

	tmpl.Execute(&output, tc)

	return output.Bytes()
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

	files, err := ioutil.ReadDir(srcFolder)
	check(err)

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
	check(err)

	icon := doc.Find("symbol").OuterHtml()

	if icon == "" {
		panic("I haven't found any <symbol> in the svg")
	}

	return icon
}

func extractSvgContent(files []string) string {
	var svgIcons string

	for _, file := range files {

		dat, err := ioutil.ReadFile(file)
		check(err)

		svgIcons = svgIcons + "\n" + getSvg(dat)
	}

	return svgIcons
}

func getFolderPath() (string, string) {
	if len(os.Args) != 3 {
		panic("you must pass a src folder path and a dest one")
	}

	folders := os.Args[1:len(os.Args)]

	if !exists(folders[0]) {
		panic(fmt.Sprintf("the source folder '%s' doesn't exist", folders[0]))
	}

	if !exists(folders[1]) {
		panic(fmt.Sprintf("the destination folder '%s' doesn't exist", folders[1]))
	}

	return os.Args[1], os.Args[2]
}

func removeFile(fileName string) {
	os.Remove(fileName)
}

func main() {
	// get srcFolder and destFolder from cli
	srcFolder, destFolder := getFolderPath()

	log.Printf("The source folder is : '%s'", srcFolder)
	log.Printf("The destination folder is: '%s'", destFolder)

	// create the svg output file and delete the previous one
	outFileName := path.Join(destFolder, "result.svg")
	removeFile(outFileName)

	// retrieve the icons from the src folder
	fileNames := fetchIcons(srcFolder)

	log.Printf("fileNames: %#+v\n", fileNames)

	// extract the svgContent
	svgIcons := extractSvgContent(fileNames)

	// open Template
	tmpl := openTemplate()

	// inject icons into svg Template
	svgFileOutput := injectIconsIntoSvgTemplate(svgIcons, tmpl)

	log.Printf("svgFileOutput: '%s'", string(svgFileOutput))

	ioutil.WriteFile(outFileName, svgFileOutput, 0644)
}
