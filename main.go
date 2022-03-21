package main

import (
	"akseach/model"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func bufioScannerToString(file *bufio.Scanner) []string {
	file.Split(bufio.ScanLines)
	var slice []string
	for file.Scan() {
		slice = append(slice, file.Text())
	}
	return slice
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
	var Dir, Url []string
	switch clix.Type {
	case "auto":
		if clix.Dir == "Stillness Speaks" {
			model.UpdateDictionaries()
			dictionariesFile, err := model.SearchDictionariesFile()
			if err != nil {
				model.PanicLog.Panic(err)
			}
			if len(dictionariesFile) == 0 {
				model.ErrorLog.Fatal("not found defaultDictionaries file")
			}
			if len(dictionariesFile) >= 2 {
				model.ErrorLog.Fatal("There are multiple dictionary files: " + fmt.Sprint(dictionariesFile))
			}
			file, err := model.ReadFile(dictionariesFile[0])
			if err != nil {
				model.PanicLog.Panic(err)
			}
			Dir = bufioScannerToString(file)
		} else {
			file, err := model.ReadFile(clix.Dir)
			if err != nil {
				model.PanicLog.Panic(err)
			}
			Dir = bufioScannerToString(file)
		}
		file, err := model.ReadFile(clix.Url)
		if err != nil {
			model.PanicLog.Panic(err)
		}
		Url = bufioScannerToString(file)
		break
	default:
		if clix.Dir == "Stillness Speaks" {
			model.UpdateDictionaries()
			dictionariesFile, err := model.SearchDictionariesFile()
			if err != nil {
				model.PanicLog.Panic(err)
			}
			if len(dictionariesFile) == 0 {
				model.ErrorLog.Fatal("not found defaultDictionaries file")
			}
			if len(dictionariesFile) >= 2 {
				model.ErrorLog.Fatal("There are multiple dictionary files: " + fmt.Sprint(dictionariesFile))
			}
			file, err := model.ReadFile(dictionariesFile[0])
			if err != nil {
				model.PanicLog.Panic(err)
			}
			Dir = bufioScannerToString(file)
		} else {
			file, err := model.ReadFile(clix.Dir)
			if err != nil {
				model.PanicLog.Panic(err)
			}
			Dir = bufioScannerToString(file)
		}
		Url = strings.Split(clix.Url, ",")
	}
	fmt.Println(Dir, Url)
}
