package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

var exit = _exit

func _exit() {
	os.Exit(1)
}

func logAndExit(message string) {
	fmt.Println("######")
	fmt.Println(message)
	fmt.Println("######")
	exit()
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
		logAndExit("I haven't found any <symbol> in the svg")
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
		logAndExit("You must pass a source folder path and a destination one")
	}

	folders := os.Args[1:len(os.Args)]

	if !exists(folders[0]) {
		logAndExit(fmt.Sprintf("The source folder '%s' doesn't exist", folders[0]))
	}

	if !exists(folders[1]) {
		logAndExit(fmt.Sprintf("The destination folder '%s' doesn't exist", folders[1]))
	}

	return os.Args[1], os.Args[2]
}

func removeFile(fileName string) {
	os.Remove(fileName)
}

func main() {
	fmt.Print(`
Welcome to Svg-Coke - A small terminal app for creating svg sprites
            ___
          .!...!.
           !   !
           ;   :
          ;     :
         ;_______:
         !svgcoke!
         !_______!
         :       :
         :       :
          ;     ;
         :       :
          '''''''
`)

	// get srcFolder and destFolder from cli
	srcFolder, destFolder := getFolderPath()

	fmt.Printf("The source folder is : '%s'\n", srcFolder)
	fmt.Printf("The destination folder is: '%s'\n", destFolder)

	// create the svg output file and delete the previous one
	outFileName := path.Join(destFolder, "result.svg")
	removeFile(outFileName)

	// retrieve the icons from the src folder
	fileNames := fetchIcons(srcFolder)

	// extract the svgContent
	svgIcons := extractSvgContent(fileNames)

	// open Template
	tmpl := openTemplate()

	// inject icons into svg Template
	svgFileOutput := injectIconsIntoSvgTemplate(svgIcons, tmpl)

	ioutil.WriteFile(outFileName, svgFileOutput, 0644)
	fmt.Println("Done!")
}
