package main

import (
	"context"
	"fmt"

	ernie "github.com/anhao/go-ernie"
)

func main() {
	requestStr := ""
	client := ernie.NewDefaultClient("LyQqEvytdvGNeeJFnj9hAYRJ", "JowJnilOi5DKWiNgkxNlHZiz7kxiTsfA")
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
	fmt.Println(completion.Result)
}
