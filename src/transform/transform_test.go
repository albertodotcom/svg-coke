package transform

import (
	"log"
	"net/textproto"
	"reflect"
	"testing"
	"text/template"
)

func TestOpenTemplate(t *testing.T) {
	result := openTemplate()

	expectedType := "*template.Template"

	if reflect.TypeOf(result).String() != expectedType {
		t.Logf("expected result to be of type \"%s\", but found type \"%s\"", expectedType, reflect.TypeOf(result))
		t.Fail()
	}
}

func TestInjectIconIntoTemaple(t *testing.T) {
	log.Println("hello")

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
