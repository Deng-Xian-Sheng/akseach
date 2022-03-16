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

var (
	DefaultDictionaries = "./defDic.txt"
	GithubDictionaries1 = "https://github.com/TheKingOfDuck/fuzzDicts/raw/master/directoryDicts/top7000.txt"
	GithubDictionaries2 = "https://github.com/TheKingOfDuck/fuzzDicts/raw/master/directoryDicts/vuls/all.txt"
	GithubAgent         = "https://service-hdekv5a8-1301929310.hk.apigw.tencentcs.com/release/GithubRaw/"
)

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
			if model.IfFileDir(DefaultDictionaries) {
				model.InfoLog.Println("Get Cloud Dictionaries ....")
				//此处获取云端字典

				//resp, err := http.Get(GithubDictionaries1)
				//if err == nil {
				//	defer resp.Body.Close()
				//	result, err := ioutil.ReadAll(resp.Body)
				//}
				//resp, err := http.Get(GithubDictionaries2)
				//if err == nil {
				//	defer resp.Body.Close()
				//	result, err := ioutil.ReadAll(resp.Body)
				//}
				model.InfoLog.Println("Get Cloud Dictionaries fail ....")
				model.InfoLog.Println("Test use vpn ....")
				//此处通过代理获取云端字典
				//通过代理获取的信息是base64的
				//resp, err = http.Get(GithubAgent + GithubDictionaries1)
				//if err == nil {
				//	defer resp.Body.Close()
				//	result, err := ioutil.ReadAll(resp.Body)
				//}
				//resp, err = http.Get(GithubAgent + GithubDictionaries2)
				//if err == nil {
				//	defer resp.Body.Close()
				//	result, err := ioutil.ReadAll(resp.Body)
				//}
			} else {
				//此处检测字典更新
				//计划通过字典文件的文件名记录版本。
			}
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
