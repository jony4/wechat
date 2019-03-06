package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jony4/wechat"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Duration(time.Second))
	defer cancel()
	opts := []wechat.ClientOptionFunc{
		wechat.SetBaseURI("www.baidu.com"),
	}
	client, err := wechat.NewClient(opts...)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.MiniProgramAuth().Set("test1", "test2").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
