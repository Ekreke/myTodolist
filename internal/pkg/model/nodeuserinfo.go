package model

type Nodeuserinfo struct {
	UserIds     []int64 `json:"user_ids"`
	UserDoneIds []int64 `json:"user_done_ids"`
}
