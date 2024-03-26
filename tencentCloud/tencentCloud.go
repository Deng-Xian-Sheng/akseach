package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

func start(ctx context.Context, event events.APIGatewayRequest) (*events.APIGatewayResponse, error) {
	// 构建目标URL
	targetURL, err := url.Parse(event.Path[11:]) // 假设您还是需要去除某些前缀
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	if len(event.QueryString) > 0 {
		queryParams := url.Values{}
		for key, value := range event.QueryString {
			// 这里假设每个key只有一个value
			queryParams.Add(key, value[0])
		}
		targetURL.RawQuery = queryParams.Encode()
	}

	// 发起到目标站点的请求
	resp, err := http.Get(targetURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 构建并返回APIGatewayResponse，设置Content-Type为application/json
	return &events.APIGatewayResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(result),
	}, nil
}

func main() {
	cloudfunction.Start(start)
}
