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
	client, err := wechat.NewClient(wechat.NewMiniProgramAuthOpts()...)
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
