package main

import (
	"net/textproto"
	"reflect"
	"regexp"
	"text/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transform", func() {
	Describe("#openTemplate", func() {
		It("returns a template", func() {
			result := openTemplate()

			expectedType := "*template.Template"

			Expect(reflect.TypeOf(result).String()).To(Equal(expectedType))
		})
	})

	Describe("#injectIconsIntoSvgTemplate ", func() {
		It("injects icons into a template", func() {
			const svgTemplateContainer = `
		  <svg>
		    <defs>
		      {{.Icons}}
		    <defs>
		  </svg>`

			stringToBeInjected := `<symbol id="hello"></symbol>`

			tmpl, _ := template.New("svgSpriteTemplate").Parse(textproto.TrimString(svgTemplateContainer))

			resultByte := injectIconsIntoSvgTemplate(stringToBeInjected, tmpl)
			result := resultByte.String()

			expectedResult := `<svg>
		    <defs>
		      <symbol id="hello"></symbol>
		    <defs>
		  </svg>`

			Expect(result).To(Equal(expectedResult))
		})
	})

	Describe("#exists", func() {
		It("returns false if a folder doesn't exist", func() {
			Expect(exists("/not/a/folder")).To(Equal(false))
		})

		It("returns true if a folder exists", func() {
			Expect(exists(".")).To(Equal(true))
		})
	})

	Describe("#fetchIcons", func() {
		It("throws an exception if src folder doesn't exist", func() {
			expect := func() {
				fetchIcons("/not/a/folder")
			}

			Expect(expect).To(Panic())
		})

		It("returns a list of svg files", func() {
			files := fetchIcons("./test-assets")

			Expect(len(files)).To(Equal(2))
		})
	})

	Describe("#extractSvgContent ", func() {
		It("returns the expected content", func() {
			var expectedResult = `<symbol id="icon-1" viewBox="0 0 1024 1024">
				<title>icon 1</title>
				<path class="path1" d="M512 32l-512 512 96 96 96-96v416h256v-192h128v192h256v-416l96 96 96-96-512-512zM512 448c-35.346 0-64-28.654-64-64s28.654-64 64-64c35.346 0 64 28.654 64 64s-28.654 64-64 64z"></path>
			</symbol>
			<symbol id="icon-2" viewBox="0 0 1024 1024">
				<title>icon 2</title>
				<path class="path1" d="M512 32l-512 512 96 96 96-96v416h256v-192h128v192h256v-416l96 96 96-96-512-512zM512 448c-35.346 0-64-28.654-64-64s28.654-64 64-64c35.346 0 64 28.654 64 64s-28.654 64-64 64z"></path>
			</symbol>`

			files := fetchIcons("./test-assets")

			reg, _ := regexp.Compile("[\\s\t\n]")

			result := extractSvgContent(files)

			Expect(reg.ReplaceAllString(result, "")).To(Equal(reg.ReplaceAllString(expectedResult, "")))
		})
	})
})
