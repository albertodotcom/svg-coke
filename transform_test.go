package main

import (
	"net/textproto"
	"text/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transform", func() {
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
})