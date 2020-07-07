package rules

/*

 1) Todos los element del documento deben tener nombres con el primer caracter en minuscula
 2) Cuando un element tiene nombre "codigoRetorno" y "glosaRetorno" deben tener los valores
 	minOccurs="1" maxOccurs="1" y El tipo de dato del codigo de retorno debe ser String
 3) Los element no deben tener caracteres especiales, solo palabras y numeros
 4) Primer ComplexType debe tener 3 elementos : "codigoRetorno" , "glosaRetorno"  , "respuesta"

*/
//Rule : Reglas de Ejecucion
type Rule struct {
	result string
	detail string
}

//GetResult : Retorna el resultado de la regla
func (r *Rule) GetResult() string {
	if r.result != "OK" {
		return r.result + ":" + r.detail
	}
	return r.result
}
