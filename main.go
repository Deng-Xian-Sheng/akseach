package main

import (
	"akseach/model"
	"bufio"
	"log"
)

func bufioScannerToString(file *bufio.Scanner) string {
	var content string
	file.Split(bufio.ScanLines)
	for file.Scan() {
		if content == "" {
			content = content + file.Text()
		} else {
			content = content + "\n" + file.Text()
		}
	}
	return content
}

func main() {
	if model.LogOsFileErr != nil {
		log.Panic(model.LogOsFileErr)
	}
	defer model.LogOsFile.Close()
	clix, err := model.Clix()
	if err != nil {
		model.ErrorLog.Fatal(err)
	}
	if clix.Type == "auto" {
		if clix.Dir == "Stillness Speaks" {
			model.UpdateDictionaries()
		} else {
			file, err := model.ReadFile(clix.Dir)
			if err != nil {
				model.PanicLog.Panic(err)
			}
			clix.Dir = bufioScannerToString(file)
		}
		file, err := model.ReadFile(clix.Url)
		if err != nil {
			model.PanicLog.Panic(err)
		}
		clix.Url = bufioScannerToString(file)
	}
	//此处写交互模式内容
}
