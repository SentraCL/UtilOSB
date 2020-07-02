package rules

import (
	"strings"
)

//Execute1 :Todos los ComplexType\Sequence\Element del documento deben tener nombres con el primer caracter en minuscula
func (r *Rule) Execute1(xsd Schema) {
	for _, complexType := range xsd.ComplexType {
		for _, element := range complexType.Sequence.Element {
			if element.Name[0:1] == strings.ToLower(element.Name[0:1]) {
				r.result = "OK"
			} else {
				r.result = "NOK"
			}
		}
	}
}
