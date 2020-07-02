package main

import (
	"flag"
	"log"
	"path/filepath"

	manager "./rules"
)

func main() {
	var sourcepath string
	var reportpath string

	flag.StringVar(&sourcepath, "p", "", "Path donde se encuentran los fuentes a evaluar")
	flag.StringVar(&reportpath, "r", "", "Path donde se generara el reporte HTML.")
	flag.Parse()

	sourcepath, _ = filepath.Abs(sourcepath)
	reportpath, _ = filepath.Abs(reportpath)

	manager := manager.CheckRuleManager{}
	manager.Run(sourcepath, reportpath)
	log.Println("Finish!")
}
