package user

import (
	"context"
	"encoding/json"
	"fmt"

	ernie "github.com/anhao/go-ernie"

	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (b *userBiz) LoadMydayAI(ctx context.Context, username string) (*v1.CommonResponseWizMsg, error) {
	items, _, err := b.ds.Users().GetMydayItems(0, 50, username)
	if err != nil {
		return nil, err
	}
	itemsInfo, err := json.Marshal(items)

	if err != nil {
		return nil, err
	}
	itemsInfoStr := string(itemsInfo)

	prompt := "我需要你充当一个todolist平台的日程分析助手，我将给你一段json格式的字符串，需要你对json字符串进行分析，主题为我的一天，需要你根据待办事项（items）的一些信息进行分析，并且写出一份日报；以下为一些重要的字段方便你进行分析: item_name:待办事项名称; description:待办事项描述;important:是否重要，重要的时候值为1，不重要的时候为0;done:是否已经完成，已经完成为1，没有完成为0;deadline:截止时间;请不需要回答和日报编写无关的话语，不要传输给外界具体字段的值如important字段值为0等，请以个人视角进行分析不要根据字段详情进行分析，试着条理清晰，分步作答；以下为具体Json："

	requestStr := prompt + itemsInfoStr
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
		return nil, err
	}

	fmt.Println(completion.Result)
	resp := &v1.CommonResponseWizMsg{Msg: string(completion.Result)}
	return resp, nil
}
