package model

import (
	"encoding/base64"
	"github.com/tidwall/gjson"
	"os"
	"path/filepath"
	"regexp"
)

var (
	DefaultDictionaries = "defDic_"
	GithubWarehouse     = "TheKingOfDuck/fuzzDicts"
	GithubDictionaries1 = "https://github.com/TheKingOfDuck/fuzzDicts/raw/master/directoryDicts/top7000.txt"
	GithubDictionaries2 = "https://github.com/TheKingOfDuck/fuzzDicts/raw/master/directoryDicts/vuls/all.txt"
	GithubAgent         = "https://service-hdekv5a8-1301929310.hk.apigw.tencentcs.com/release/GithubRaw/"
)

func updateDictionaries(update bool) string {
	if update {
		defer func() {
			if err := recover(); err != nil {
				LogWendy.Error(err)
				LogWendy.Error("Dictionaries Update fail, but can use old")
			}
		}()
	}
	LogWendy.Info("Get Cloud Dictionaries ....")
	var Dictionaries string
	bytes, err := Get(GithubDictionaries1)
	if err == nil {
		decodeString, err := base64.StdEncoding.DecodeString(gjson.Get(string(bytes), "body").Str)
		if err != nil {
			LogWendy.Fatal(err)
		}
		Dictionaries = string(decodeString)
	} else {
		LogWendy.Error("Get Cloud Dictionaries fail ....")
		LogWendy.Info("Test use vpn ....")
		bytes, err = Get(GithubAgent + GithubDictionaries1)
		if err != nil {
			LogWendy.Fatal("Not Get Cloud Dictionaries!")
		}
		decodeString, err := base64.StdEncoding.DecodeString(gjson.Get(string(bytes), "body").Str)
		if err != nil {
			LogWendy.Fatal(err)
		}
		Dictionaries = string(decodeString)

		bytes, err = Get(GithubAgent + GithubDictionaries2)
		if err != nil {
			LogWendy.Fatal("Not Get Cloud Dictionaries!")
		}
		decodeString, err = base64.StdEncoding.DecodeString(gjson.Get(string(bytes), "body").Str)
		if err != nil {
			LogWendy.Fatal(err)
		}
		return Dictionaries + "\n" + string(decodeString)
	}
	bytes, err = Get(GithubDictionaries2)
	if err != nil {
		LogWendy.Fatal(err)
	}
	decodeString, err := base64.StdEncoding.DecodeString(gjson.Get(string(bytes), "body").Str)
	if err != nil {
		LogWendy.Fatal(err)
	}
	return Dictionaries + "\n" + string(decodeString)
}
func SearchDictionariesFile() ([]string, error) {
	files, err := filepath.Glob("*")
	if err != nil {
		return nil, err
	}
	reg := DefaultDictionaries + `[0-9]{4}\-[0-9]{2}\-[0-9]{2}T[0-9]{2}\:[0-9]{2}\:[0-9]{2}Z\.txt`
	var resultSlice []string
	for _, v := range files {
		if match, err := regexp.Match(reg, []byte(v)); err == nil && match == true {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice, nil
}
func UpdateDictionaries() {
	resultSlice, err := SearchDictionariesFile()
	if err != nil {
		LogWendy.Fatal(err)
	}
	if len(resultSlice) == 0 {
		dictionaries := updateDictionaries(false)
		scan, err := UpdateScan(GithubWarehouse)
		if err != nil {
			LogWendy.Fatal(err)
		}
		err = WriteFile(DefaultDictionaries+scan+".txt", []byte(dictionaries))
		if err != nil {
			LogWendy.Fatal(err)
		}
		LogWendy.Info("Create dictionaries success!")
		return
	}
	if len(resultSlice) >= 2 {
		LogWendy.Fatal("There are multiple dictionary files: ", resultSlice)
	}
	compile := regexp.MustCompile(DefaultDictionaries + `([0-9]{4}\-[0-9]{2}\-[0-9]{2}T[0-9]{2}\:[0-9]{2}\:[0-9]{2}Z)\.txt`)
	if compile == nil {
		LogWendy.Fatal("MustCompile err")
	}
	submatch := compile.FindAllStringSubmatch(resultSlice[0], -1)
	scan, err := UpdateScan(GithubWarehouse)
	if err != nil {
		LogWendy.Error(err)
		LogWendy.Error("Dictionaries Update fail, but can use old")
		return
	}
	if submatch[0][1] != scan {
		dictionaries := updateDictionaries(true)
		err = WriteFile(DefaultDictionaries+scan+".txt", []byte(dictionaries))
		if err != nil {
			LogWendy.Fatal(err)
		}
		err := os.Remove(DefaultDictionaries + submatch[0][1] + ".txt")
		if err != nil {
			LogWendy.Fatal(err)
		}
		LogWendy.Info("Update dictionaries success!")
		return
	}
	LogWendy.Info("Dictionaries is very new!")
	return
}
