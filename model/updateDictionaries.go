package model

import "encoding/base64"

var (
	DefaultDictionaries = "./defDic.txt"
	GithubDictionaries1 = "https://github.com/TheKingOfDuck/fuzzDicts/raw/master/directoryDicts/top7000.txt"
	GithubDictionaries2 = "https://github.com/TheKingOfDuck/fuzzDicts/raw/master/directoryDicts/vuls/all.txt"
	GithubAgent         = "https://service-hdekv5a8-1301929310.hk.apigw.tencentcs.com/release/GithubRaw/"
)

func updateDictionaries() string {
	InfoLog.Println("Get Cloud Dictionaries ....")
	var Dictionaries string
	bytes, err := Get(GithubDictionaries1)
	if err == nil {
		decodeString, err := base64.StdEncoding.DecodeString(string(bytes))
		if err != nil {
			PanicLog.Panic(err)
		}
		Dictionaries = string(decodeString)
	} else {
		InfoLog.Println("Get Cloud Dictionaries fail ....")
		InfoLog.Println("Test use vpn ....")
		bytes, err = Get(GithubAgent + GithubDictionaries1)
		if err != nil {
			PanicLog.Panic("Not Get Cloud Dictionaries!")
		}
		decodeString, err := base64.StdEncoding.DecodeString(string(bytes))
		if err != nil {
			PanicLog.Panic(err)
		}
		Dictionaries = string(decodeString)

		bytes, err = Get(GithubAgent + GithubDictionaries2)
		if err != nil {
			PanicLog.Panic("Not Get Cloud Dictionaries!")
		}
		decodeString, err = base64.StdEncoding.DecodeString(string(bytes))
		if err != nil {
			PanicLog.Panic(err)
		}
		return Dictionaries + "\n" + string(decodeString)
	}
	bytes, err = Get(GithubDictionaries2)
	if err != nil {
		PanicLog.Panic(err)
	}
	decodeString, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		PanicLog.Panic(err)
	}
	return Dictionaries + "\n" + string(decodeString)
}

func UpdateDictionaries() {
	if IfFileDir(DefaultDictionaries) {
		updateDictionaries()
		//此处写入字典
	} else {
		//此处检测字典更新
		//计划通过字典文件的文件名记录版本。
	}
}
