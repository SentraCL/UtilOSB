package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

//StringifyJSON , convierte cualquier objeto en una representacion STRING en formato JSON
func StringifyJSON(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		return "Error"
	}
	return string(b)
}

//CreateFolder , Crea carpeta
func CreateFolder(folder string) bool {
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return false
	}
	return true
}

//CopyFile , Copiar archivo
func CopyFile(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}
