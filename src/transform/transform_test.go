package transform

import (
	"log"
	"net/textproto"
	"testing"
	"text/template"
)

func Test_InjectStringIntoTemaple(t *testing.T) {
	const svgTemplateContainer = `
  <svg>
    <defs>
      {{.StringToInject}}
    <defs>
  </svg>`

	stringToBeInjected := `<symbol id="hello"></symbol>`

	tmpl, _ := template.New("svgSpriteTemplate").Parse(textproto.TrimString(svgTemplateContainer))

	resultByte := injectStringIntoTemplate(stringToBeInjected, tmpl)
	result := resultByte.String()

	log.Printf("result: %#+v\n", result)

	expectedResult := `<svg>
    <defs>
      <symbol id="hello"></symbol>
    <defs>
  </svg>`

	if result != expectedResult {
		t.Logf("expected \"%s\" to be equal to \"%s\"", result, expectedResult)
		t.Fail()
	}
}
