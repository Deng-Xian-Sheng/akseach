package model

//URL格式化
func FormatURL(url string) string {
	if url == "" {
		return ""
	}
	if url[0] == '/' {
		return url
	}
	if url[0:4] == "http" {
		return url
	}
	return "http://" + url
}

//URL尾部格式化
func FormatURLTail(url string) string {
	if url == "" {
		return ""
	}
	if url[len(url)-1] == '/' {
		return url
	}
	return url + "/"
}

//判断URL尾部是否为/
func IsURLTail(url string) bool {
	if url == "" {
		return false
	}
	if url[len(url)-1] == '/' {
		return true
	}
	return false
}

//判断路径头部是否为/
func IsPathTail(path string) bool {
	if path == "" {
		return false
	}
	if path[0] == '/' {
		return true
	}
	return false
}
