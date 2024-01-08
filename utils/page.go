package utils

import (
	"encoding/base64"
	"encoding/json"

	"github.com/ekreke/myTodolist/pkg/logging"
)

// record page info
type Page struct {
	// the next id is the end record of the previous page
	NextID        string `json:"next_id"`
	NextTimeAtUTC int64  `json:"next_time_at_utc"`
	PageSize      int64  `json:"page_size"`
}

// encode page tokne
func Encode(page *Page) string {
	b, err := json.Marshal(page)
	if err != nil {
		logging.Info("encode page token err:", err)
		return ""
	}
	return (base64.StdEncoding.EncodeToString(b))
}

// decode page token
func Decode(pagetoken string) Page {
	var result Page
	if len(pagetoken) == 0 {
		logging.Info("decode page token err: empty token")
		return result
	}
	bytes, err := base64.StdEncoding.DecodeString(pagetoken)
	if err != nil {
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}

	return result
}
