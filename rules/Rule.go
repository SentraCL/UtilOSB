package rules

import "strings"

//Rule : Reglas de Ejecucion
type Rule struct {
	result string
	detail string
	MyRule RuleJSON
}

//GetResult : Retorna el resultado de la regla
func (r *Rule) GetResult() string {
	if r.result != "OK" {
		return r.result + ":" + r.detail
	}
	return r.result
}

func (r *Rule) isXSD(nameFile string) bool {
	return strings.HasSuffix(strings.ToLower(nameFile), "xsd")
}

func (r *Rule) isRequestXSD(nameFile string) bool {
	return strings.HasSuffix(strings.ToLower(nameFile), "rq.xsd")
}

func (r *Rule) isResponseXSD(nameFile string) bool {
	return strings.HasSuffix(strings.ToLower(nameFile), "rs.xsd")
}

func (r *Rule) isCorrectEXT(nameFile string) bool {
	return strings.HasSuffix(strings.ToLower(nameFile), r.MyRule.Ext)
}

func (r *Rule) isWSDL(nameFile string) bool {
	return strings.HasSuffix(strings.ToLower(nameFile), "wsdl")
}
