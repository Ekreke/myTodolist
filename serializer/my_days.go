package serializer

import "github.com/ekreke/myTodolist/model"

type My_Days struct {
	Items       []model.Item `json:"item"`
	ProCurToken string       `json:"pro_cur_token"`
}
