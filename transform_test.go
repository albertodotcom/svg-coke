package main

import (
	"errors"
	"io/ioutil"
	"net/textproto"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func mockExit() {
	panic("exit")
}

var _ = Describe("Transform", func() {
	Describe("#check", func() {
		It("panics when e is not nil", func() {
			e := errors.New("emit macho dwarf: elf header corrupted")

			result := func() {
				check(e)
			}

			Expect(result).To(Panic())
		})

		It("doesn't panic when e isn't defined", func() {
			result := func() {
				check(nil)
			}

			Expect(result).NotTo(Panic())
		})
	})

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
			result := string(resultByte)

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

	Describe("#getSvg ", func() {
		It("returns the symbol out of an svg", func() {
			var expectedResult = `
			<symbol id="icon-1" viewBox="0 0 1024 1024">
				<title>icon 1</title>
				<path class="path1" d="M512 32l-512 512 96 96 96-96v416h256v-192h128v192h256v-416l96 96 96-96-512-512zM512 448c-35.346 0-64-28.654-64-64s28.654-64 64-64c35.346 0 64 28.654 64 64s-28.654 64-64 64z"></path>
			</symbol>`

			files := fetchIcons("./test-assets")

			reg, _ := regexp.Compile("[\\s\t\n]")

			fileBytes, _ := ioutil.ReadFile(files[0])

			result := getSvg(fileBytes)

			Expect(reg.ReplaceAllString(result, "")).To(Equal(reg.ReplaceAllString(expectedResult, "")))
		})

		It("returns an error if the svg doesn't have a symbol element", func() {
			files := fetchIcons("./test-assets")

			fileBytes, _ := ioutil.ReadFile(files[0])

			fileString := string(fileBytes)

			fileString = strings.Replace(fileString, "symbol", "g", -1)

			fileBytes = []byte(fileString)

			expect := func() {
				getSvg(fileBytes)
			}

			origExit := exit
			exit = mockExit
			defer func() { exit = origExit }()

			Expect(expect).To(Panic())
		})
	})

	Describe("#extractSvgContent ", func() {
		It("returns the expected content", func() {
			var expectedResult = `
			<symbol id="icon-1" viewBox="0 0 1024 1024">
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

	Describe("#getFolderPath ", func() {
		It("panics if args isn't 3", func() {
			os.Args[4] = "four"
			expect := func() {
				getFolderPath()
			}

			origExit := exit
			exit = mockExit
			defer func() { exit = origExit }()

			Expect(expect).To(Panic())

			os.Args = append(os.Args[:0], os.Args[1])

			Expect(expect).To(Panic())
		})

		It("returns 2 strings ./test-assets ./test-assets", func() {
			os.Args = append(os.Args[:0], "ignore", "./test-assets", "./test-assets")

			first, second := getFolderPath()

			Expect(first).To(Equal("./test-assets"))
			Expect(second).To(Equal("./test-assets"))
		})

		It("exits if src folder doesn't exit", func() {
			os.Args = append(os.Args[:0], "ignore", "./doesntExist", "./test-assets")

			origExit := exit
			exit = mockExit
			defer func() { exit = origExit }()

			expect := func() {
				getFolderPath()
			}

			Expect(expect).To(Panic())
		})

		It("exits if dest folder doesn't exit", func() {
			os.Args = append(os.Args[:0], "ignore", "./test-assets", "./doesntExist")

			origExit := exit
			exit = mockExit
			defer func() { exit = origExit }()

			expect := func() {
				getFolderPath()
			}

			Expect(expect).To(Panic())
		})
	})
})
