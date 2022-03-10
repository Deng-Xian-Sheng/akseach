package main

import (
	"akseach/model"
	"bufio"
	"log"
	"os"
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
	defer func() {
		os.File.Close(model.PanicLog)
		os.File.Close(model.ErrorLog)
		os.File.Close(model.InfoLog)
	}()
	clix, err := model.Clix()
	if err != nil {
		model.PanicLog.Panic(err)
	}
	if clix.Type == "auto" {
		if clix.Dir == "Stillness Speaks" {
			//此处获取云端字典并赋值给clix.Dir
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
