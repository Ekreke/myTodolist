package mytodolist

import "github.com/ekreke/myTodolist/internal/pkg/model"

type CommonResquestWiz struct {
}

type CommonResponseWizMsg struct {
	Msg string `json:"msg"`
}

type CommonRequestWizPagination struct {
	Pagination string `form:"pagination"`
}

type CommonResponseWizItemsAndPagination struct {
	Items      []model.Items `json:"items"`
	Pagination string        `json:"pagination"`
}

type CollectionLoadItemsResp struct {
	Items []model.Items `json:"items"`
}
