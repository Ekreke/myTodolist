package serializer

// common serilazer
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Msg:    "ok",
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
