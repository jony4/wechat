package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jony4/wechat"
)

func main() {
	logger := wechat.NewDefaultLogger()
	// init default client
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Duration(time.Second))
	defer cancel()
	opts := []wechat.ClientOptionFunc{
		wechat.SetInfoLog(logger),
		wechat.SetErrorLog(logger),
		wechat.SetTraceLog(logger),
	}
	client, err := wechat.NewClient(opts...)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.MiniProgramAuth().
		SetAppID("xxxx").
		SetSecret("xxxx").
		SetJscode("xxxx").
		Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
