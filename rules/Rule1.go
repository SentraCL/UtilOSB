package rules

import (
	"strings"
)

//Execute1 :Todos los ComplexType\Sequence\Element del documento deben tener nombres con el primer caracter en minuscula
func (r *Rule) Execute1(xsd Schema) {
	r.result = "OK"
	if r.isXSD(xsd.XMLFile) {
		for _, complexType := range xsd.ComplexType {
			for _, element := range complexType.Sequence.Element {
				//fmt.Println(element.Name)
				if element.Name[0:1] == strings.ToLower(element.Name[0:1]) {
					r.result = "OK"
				} else {
					r.result = "NOK"
					r.detail = r.detail + "Nombre no valido " + element.Name + "\n"
				}
			}
		}
	}
}
