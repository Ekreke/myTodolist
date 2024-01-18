package token

import (
	"encoding/base64"
	"encoding/json"
)

type Token string
type Page struct {
	NextID   int   `json:"next_id"`
	PageSize int64 `json:"page_size"`
}

// Encode 返回分页token
func (p Page) PageEncode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

// Decode 解析分页信息
func (t Token) PageDecode() Page {
	var result Page
	if len(t) == 0 {
		return result
	}

	bytes, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}

	return result
}
