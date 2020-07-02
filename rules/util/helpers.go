package helpers

import (
	"encoding/json"
	"io/ioutil"
)

//LoadFile :Carga archivos!
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//IsEmpty es un string vacio?
func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
}

//ToStringifyJSON , convierte cualquier objeto en una representacion STRING en formato JSON
func StringifyJSON(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		return "Error"
	}
	return string(b)
}
