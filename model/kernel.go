package model

import (
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

func httpClient() *http.Client {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		LogWendy.Warn(err)
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
	//计算程序耗时
	startTime := time.Now().Unix()
	var wg sync.WaitGroup
	limiter := NewConcurrencyLimiter(2)
	num := 1
	sliceLowHigh := len(Url) / num
	for i := 1; i <= num; i++ {
		wg.Add(1)
		limiter.Get()
		go func(Dir, Url, Proxy []string) {
			for _, v := range Url {
				num := 1
				sliceLowHigh := len(Dir) / num
				for i := 1; i <= num; i++ {
					wg.Add(1)
					limiter.Get()
					go func(Dir, Proxy []string, Url string) {
						client := httpClient()
						for _, j := range Dir {
							wg.Add(1)
							limiter.Get()
							go func(Proxy []string, Dir, Url string) {
								defer func() {
									if err := recover(); err != nil {
										LogWendy.Error(err)
									}
								}()
								if len(Proxy) > 0 {
									parse, err := url.Parse(Proxy[rand.Intn(len(Proxy))])
									if err != nil {
										LogWendy.Fatal(err)
									}
									client.Transport.(*http.Transport).Proxy = http.ProxyURL(parse)
								}
								Url = FormatURL(Url)
								if !IsURLTail(Url) && !IsPathTail(Dir) {
									Url = FormatURLTail(Url)
								}
								result, err := client.Get(Url + Dir)
								if err != nil {
									LogWendy.Fatal(err)
								}
								if result.StatusCode == 200 || result.StatusCode == 301 || result.StatusCode == 302 {
									LogWendy.Info("Yes｜" + Url + "｜" + Dir + "｜" + Url + Dir)
								}
								limiter.Release()
								wg.Done()
							}(Proxy, j, Url)
						}
						limiter.Release()
						wg.Done()
					}(Dir[sliceLowHigh*(i-1):sliceLowHigh*i], Proxy, v)
				}
			}
			limiter.Release()
			wg.Done()
		}(Dir, Url[sliceLowHigh*(i-1):sliceLowHigh*i], Proxy)
	}
	wg.Wait()
	LogWendy.Info("耗时：" + time.Now().Sub(time.Unix(startTime, 0)).String())
}

/*
待改进：
设置URL、目录字典合适的分片
根据请求成功率设置同时并发线程数
根据某个URL的请求成功率设置请求超时时间
根据某个URL的请求成功率放弃某个URL
*/
