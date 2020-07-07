package rules

import (
	"regexp"
)

// 3) Los element no deben tener caracteres especiales, solo palabras y numeros

func (r *Rule) Execute3(xsd Schema) {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	invalidElementList := []string{}
	for _, complexType := range xsd.ComplexType {
		for _, element := range complexType.Sequence.Element {
			if re.MatchString(element.Name) {
				invalidElementList = append(invalidElementList, element.Name)
				r.detail = r.detail + element.Name
			}
		}
	}
	if len(invalidElementList) > 0 {
		r.result = "NOK"
	} else {
		r.result = "OK"
	}
}
