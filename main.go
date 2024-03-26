package main

import (
	"akseach/model"
	"fmt"
	"strings"
)

func main() {
	defer model.LogOsFile.Close()
	clix, err := model.Clix()
	if err != nil {
		model.LogWendy.Fatal(err)
	}
	var Dir, Url, Proxy []string
	switch clix.Type {
	case "auto":
		if clix.Dir == "Stillness Speaks" {
			model.UpdateDictionaries()
			dictionariesFile, err := model.SearchDictionariesFile()
			if err != nil {
				model.LogWendy.Fatal(err)
			}
			if len(dictionariesFile) == 0 {
				model.LogWendy.Fatal("not found defaultDictionaries file")
			}
			if len(dictionariesFile) >= 2 {
				model.LogWendy.Fatal("There are multiple dictionary files: " + fmt.Sprint(dictionariesFile))
			}
			Dir, err = model.ReadFile(dictionariesFile[0])
			if err != nil {
				model.LogWendy.Fatal(err)
			}
		} else {
			Dir, err = model.ReadFile(clix.Dir)
			if err != nil {
				model.LogWendy.Fatal(err)
			}
		}
		Url, err = model.ReadFile(clix.Url)
		if err != nil {
			model.LogWendy.Fatal(err)
		}
		Proxy, err = model.ReadFile(clix.Proxy)
		if err != nil {
			model.LogWendy.Fatal(err)
		}
	default:
		if clix.Dir == "Stillness Speaks" {
			model.UpdateDictionaries()
			dictionariesFile, err := model.SearchDictionariesFile()
			if err != nil {
				model.LogWendy.Fatal(err)
			}
			if len(dictionariesFile) == 0 {
				model.LogWendy.Fatal("not found defaultDictionaries file")
			}
			if len(dictionariesFile) >= 2 {
				model.LogWendy.Fatal("There are multiple dictionary files: " + fmt.Sprint(dictionariesFile))
			}
			Dir, err = model.ReadFile(dictionariesFile[0])
			if err != nil {
				model.LogWendy.Fatal(err)
			}
		} else {
			Dir, err = model.ReadFile(clix.Dir)
			if err != nil {
				model.LogWendy.Fatal(err)
			}
		}
		Url = strings.Split(clix.Url, ",")
		Proxy = strings.Split(clix.Proxy, ",")
		if Proxy[0] == "" {
			Proxy = []string{}
		}
	}
	//判断Dir、Url不等于空就执Kernel函数，否则退出
	if len(Dir) != 0 && len(Url) != 0 {
		model.Kernel(Dir, Url, Proxy)
	} else {
		model.LogWendy.Fatal("Dir or Url is empty")
	}
}
