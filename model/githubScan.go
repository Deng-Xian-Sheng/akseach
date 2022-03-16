package model

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

//根据Github用户名加仓库名获得更新时间；如 Deng-Xian-Sheng/xxx
func UpdateScan(path string) (string, error) {
	resp, err := http.Get("https://api.github.com/repos/" + path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	date := gjson.Get(string(result), "updated_at").Str
	return date, nil
}
