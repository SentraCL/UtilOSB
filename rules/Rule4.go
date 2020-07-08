package rules

//Execute4 :Primer ComplexType debe tener 3 elementos : "codigoRetorno" , "glosaRetorno"  , "respuesta"
func (r *Rule) Execute4(xsd Schema) {
	r.result = "OK"
	if r.isResponseXSD(xsd.XMLFile) {
		codigoRetorno := false
		glosaRetorno := false
		respuesta := false
		for _, complexType := range xsd.ComplexType {
			for _, element := range complexType.Sequence.Element {
				if element.Name == "codigoRetorno" {
					codigoRetorno = true
				}
				if element.Name == "glosaRetorno" {
					glosaRetorno = true
				}
				if element.Name == "respuesta" {
					respuesta = true
				}
			}
			break
		}
		if codigoRetorno && glosaRetorno && respuesta {
			r.result = "OK"
		} else {
			r.result = "NOK"
			r.detail = "ComplexType no contiene algunos de los siguintes elementos codigoRetorno, glosaRetorno, respuesta"
		}
	}
}
