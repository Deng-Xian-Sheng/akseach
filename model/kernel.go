package model

import (
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

func httpClient() *http.Client {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		ErrorLog.Println(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar:     jar,
		Timeout: 10 * time.Second,
	}
	return &client
}
func Kernel(Dir, Url, Proxy []string) {
	limiter := NewConcurrencyLimiter(10)
	limiter.cond.L.Lock()
	num := 1
	sliceLowHigh := len(Url) / num
	for i := 1; i <= num; i++ {
		limiter.get()
		go func(Dir, Url, Proxy []string) {
			for _, v := range Url {
				num := 1
				sliceLowHigh := len(Dir) / num
				for i := 1; i <= num; i++ {
					limiter.get()
					go func(Dir, Proxy []string, Url string) {
						client := httpClient()
						for _, j := range Dir {
							limiter.get()
							go func(Proxy []string, Dir, Url string) {
								defer func() {
									if err := recover(); err != nil {
										ErrorLog.Println(err)
									}
								}()
								if len(Proxy) > 0 {
									parse, err := url.Parse(Proxy[rand.Intn(len(Proxy))])
									if err != nil {
										ErrorLog.Panic(err)
									}
									client.Transport.(*http.Transport).Proxy = http.ProxyURL(parse)
								}
								Url = FormatURL(Url)
								if !IsURLTail(Url) && !IsPathTail(Dir) {
									Url = FormatURLTail(Url)
								}
								result, err := client.Get(Url + Dir)
								if err != nil {
									ErrorLog.Panic(err)
								}
								if result.StatusCode == 200 || result.StatusCode == 301 || result.StatusCode == 302 {
									InfoLog.Println("Yes｜" + Url + "｜" + Dir + "｜" + Url + Dir)
								}
							}(Proxy, j, Url)
						}
					}(Dir[sliceLowHigh*(i-1):sliceLowHigh*i], Proxy, v)
				}
			}
		}(Dir, Url[sliceLowHigh*(i-1):sliceLowHigh*i], Proxy)
	}
	limiter.cond.Wait()
	limiter.cond.L.Unlock()
}
