package main

import (
	"context"
	"encoding/base64"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
	"io/ioutil"
	"net/http"
)

func start(ctx context.Context, event events.APIGatewayRequestContext) (*events.APIGatewayResponse, error) {
	resp, err := http.Get(event.Path[11:])
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayResponse{
		IsBase64Encoded: true,
		StatusCode:      200,
		Headers:         nil,
		Body:            base64.StdEncoding.EncodeToString(result),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(start)
}
