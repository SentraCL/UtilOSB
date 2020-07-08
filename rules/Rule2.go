package rules

import (
	"strings"
)

//Execute2: Cuando un element tiene nombre "codigoRetorno" o "glosaRetorno" deben tener los valores
// minOccurs="1" maxOccurs="1" y El tipo de dato del codigo de retorno debe ser String
func (r *Rule) Execute2(xsd Schema) {
	r.result = "OK"
	if r.isResponseXSD(xsd.XMLFile) {
		validacionNombre := false
		validacionString := false
		validacionOccurs := false
		for _, complexType := range xsd.ComplexType {
			for _, element := range complexType.Sequence.Element {
				if element.Name == "codigoRetorno" || element.Name == "glosaRetorno" {
					validacionNombre = true
					if strings.Contains(element.Type, "String") {
						validacionString = true
					}
					if element.MinOccurs == "1" && element.MaxOccurs == "1" {
						validacionOccurs = true
					}
				}
			}
		}
		if validacionNombre && validacionString && validacionOccurs {
			r.result = "OK"
		} else {
			r.result = "NOK"
			r.detail = "El elemento no cumple con alguna de las siguientes reglas\n1. Contiene codigoRetorno o glosaRetorno\n2. Valor en 1 en minOccurs y maxOccurs\n3. El tipo de dato debe ser String"
		}
	}
}
