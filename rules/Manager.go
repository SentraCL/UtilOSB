package rules

import (
	"bytes"
	"reflect"

	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"

	util "./util"
)

//CheckRule : Resultado de la ejecucion de reglas
type CheckRule struct {
	Rule   string `json:"rule"`
	Status string `json:"status"`
}

//CheckRuleManager , Manager Administrador de Ejecucion de Reglas.
type CheckRuleManager struct {
	sourcespath string
	resultpath  string
	rulesID     []string
}

func (m *CheckRuleManager) getDefJobFromXML() ([]Schema, error) {

	var files []string
	defjobs := []Schema{}

	err := filepath.Walk(m.sourcespath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".xsd" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return defjobs, err
	}

	for _, file := range files {
		xsdFile, _ := os.Open(file)
		xmlBytes, _ := ioutil.ReadAll(xsdFile)

		reader := bytes.NewReader(xmlBytes)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReader

		var xsdSchm Schema
		err = decoder.Decode(&xsdSchm)
		xsdSchm.XMLFile = filepath.Base(file)
		if err != nil {
			return defjobs, err
		}
		defjobs = append(defjobs, xsdSchm)
	}
	return defjobs, nil
}

//Run , Ejecuta las reglas.
func (m *CheckRuleManager) Run(sourcespath, resultpath string) {
	m.rulesID = m.readRulesFromConfig()
	m.sourcespath = sourcespath
	m.resultpath = resultpath
	folderXSD, _ := m.getDefJobFromXML()
	//result := []Schema{}

	for _, xsdContent := range folderXSD {
		rDefJob := Schema{}
		rDefJob.XMLFile = xsdContent.XMLFile
		log.Println("CHECK: \t:" + util.StringifyJSON(xsdContent.XMLFile))
		m.executeRules(xsdContent)
	}
	//file, _ := json.MarshalIndent(result, "", " ")

	//_ = ioutil.WriteFile(m.resultpath+"\\result.json", file, 0644)
}

func (m *CheckRuleManager) readRulesFromConfig() []string {
	rulesJSON, err := ioutil.ReadFile("./rules/rules.json")
	if err != nil {
		log.Print(err)
	}
	var config struct {
		Rules []string `json:"rules"`
	}
	err = json.Unmarshal(rulesJSON, &config)
	if err != nil {
		log.Println("error:", err)
	}
	return config.Rules
}

func (m *CheckRuleManager) executeRules(xsd Schema) {

	for _, id := range m.rulesID {
		check := CheckRule{}
		status := m.executeRuleByID(id, xsd)
		check.Rule = id
		check.Status = status

		//if status != "OK" {
		log.Println("Regla " + id + " = " + status)
		//}
	}
}

func (m *CheckRuleManager) toRules(any interface{}, name string, args ...interface{}) {
	params := make([]reflect.Value, len(args))
	for i := range args {
		params[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(params)
}

func (m *CheckRuleManager) executeRuleByID(id string, xsd Schema) string {
	r := Rule{}
	m.toRules(&r, "Execute"+id, xsd)
	return r.GetResult()
}
