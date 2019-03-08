# Wechat SDK

It is probably the best SDK in the world for developing WeChat App In Go.

## Usage

### Install

```
go get -u github.com/jony4/wechat
```

### Examples

```go
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
		wechat.SetGzip(true),
	}
	client, err := wechat.NewClient(opts...)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.MiniProgramAuth().
		SetAppID("xxx").  // 填写小程序的 appid
		SetSecret("xxx"). // 填写小程序的 secret
		SetJscode("xxx"). // 填写前端传入的 jscode
		Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
```
