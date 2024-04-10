package main

import (
	"context"
	"fmt"

	ernie "github.com/anhao/go-ernie"
)

func main() {
	requestStr := "介绍一下你自己"
	client := ernie.NewDefaultClient("/", "JowJnilOi5DKWiNgkxNlHZiz7kxiTsfA")
	completion, err := client.CreateErnieBotChatCompletion(context.Background(), ernie.ErnieBotRequest{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: requestStr,
			},
		},
	})
	if err != nil {
		fmt.Printf("ernie bot error: %v\n", err)
		return
	}

	fmt.Println(completion.BanRound)
	fmt.Println("-----------------------------")
	fmt.Println(completion.Result)
}
