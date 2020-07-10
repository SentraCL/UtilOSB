package rules

import (
	"bytes"
	"fmt"
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

//RuleJSON , Reglas en JSON
type RuleJSON struct {
	ID   string `json:"id"`
	Info string `json:"info"`
	Ext  string `json:"ext"`
}

//CheckRuleManager , Manager Administrador de Ejecucion de Reglas.
type CheckRuleManager struct {
	sourcespath string
	resultpath  string
	rulesDef    []RuleJSON
	report      map[string][]CheckRule
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
	m.rulesDef = m.readRulesFromConfig()
	m.sourcespath = sourcespath
	m.resultpath = resultpath
	folderXSD, _ := m.getDefJobFromXML()
	m.report = map[string][]CheckRule{}
	for _, xsdContent := range folderXSD {
		rDefJob := Schema{}
		rDefJob.XMLFile = xsdContent.XMLFile
		m.report[xsdContent.XMLFile] = []CheckRule{}
		m.executeRules(xsdContent)
	}
	m.createJStoReport()
}

func (m *CheckRuleManager) createJStoReport() {

	if util.CreateFolder(m.resultpath) {
		util.CopyFile("./static/report.html", m.resultpath+"/report.html")

		f, err := os.Create(m.resultpath + "\\result.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		code := "var data =" + util.StringifyJSON(m.report) + "\n" + "var rulesInfo=" + util.StringifyJSON(m.rulesDef)

		_, err = f.WriteString(code)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("No fue posible crear la carpeta de reporte :(")
	}
}

func (m *CheckRuleManager) readRulesFromConfig() []RuleJSON {
	rulesJSON, err := ioutil.ReadFile("./rules/rules.json")
	if err != nil {
		log.Print(err)
	}
	var config struct {
		Rules []RuleJSON `json:"rules"`
	}
	err = json.Unmarshal(rulesJSON, &config)
	if err != nil {
		log.Println("error:", err)
	}
	return config.Rules
}

func (m *CheckRuleManager) executeRules(xsd Schema) {

	for _, ruleDef := range m.rulesDef {
		check := CheckRule{}
		status := m.executeRule(ruleDef, xsd)
		check.Rule = ruleDef.ID
		check.Status = status

		if status != "OK" {
			log.Println("Regla " + ruleDef.ID + " = " + status)
		}
	}
}

func (m *CheckRuleManager) toRules(any interface{}, name string, args ...interface{}) {
	params := make([]reflect.Value, len(args))
	for i := range args {
		params[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(params)
}

func (m *CheckRuleManager) executeRule(ruleDef RuleJSON, xsd Schema) string {
	r := Rule{}
	r.MyRule = ruleDef
	m.toRules(&r, "Execute"+ruleDef.ID, xsd)
	result := CheckRule{}
	result.Rule = ruleDef.ID
	result.Status = r.GetResult()
	m.report[xsd.XMLFile] = append(m.report[xsd.XMLFile], result)

	return r.GetResult()
}
