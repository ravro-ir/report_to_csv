package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
)

type CVS struct {
	report Report
}

func main() {

	var (
		keyFixPath  string
		outFixpath  string
		curretnPath string
		status      bool
	)
	if runtime.GOOS == "windows" {
		keyFixPath = "key/key.private"
		outFixpath = "decrypt"
	} else {
		keyFixPath = "key/key.private"
		outFixpath = "decrypt"
	}
	publicMessage := "شرح داده نشد است"
	inputDir := flag.String("in", "in", "input directory of report encrypt file")
	outputDir := flag.String("out", "out", "output directory for decrypt report file ")
	key := flag.String("key", "key", "private key")
	init := flag.String("init", "", "input directory of report encrypt file")
	flag.Parse()
	if *init == "init" {
		AddDir("decrypt")
		AddDir("encrypt")
		AddDir("key")
		fmt.Println("[++] Created directory decrypt, encrypt, key")
		return
	}
	status = false
	if *inputDir != "in" {
		status = true
		curretnPath = *inputDir
	}
	if *outputDir != "out" {
		status = true
		outFixpath = *outputDir
	}
	if *key != "key" {
		status = true
		keyFixPath = *key
	}
	fmt.Println("[++++] Starting for decrypting Report . . . ")
	report, err := DcrptReport(curretnPath, keyFixPath, outFixpath, status)
	if err != nil {
		log.Fatalln(err)
	}
	cvs := CVS{report: report}
	if cvs.report.Reproduce == "" {
		cvs.report.Reproduce = publicMessage
	}
	fmt.Println("[++++] Decrypted successfully ")
	fmt.Println("[++++] Start to import to CSV ")
	col := []string{cvs.report.Slug, cvs.report.Description, cvs.report.Reproduce}
	addCol("reports.csv", col)
	fmt.Println("[++++] Import to CSV successfully")
}
