package e

var MsgFlags = map[int]string{
	SUCCESS:                             "ok",
	ERROR:                               "fail",
	INVALID_PARAMS:                      "invalid params",
	ERROR_NOT_EXIST_USER:                "user not exist",
	ERROR_PASSWORD:                      "password error",
	ERROR_USER_EXIST:                    "user exist",
	ERROR_USER_NOT_LOGIN:                "user not login",
	ERROR_DB:                            "database error",
	ERROR_AUTH_CHECK_TOKEN_FAIL:         "token auth check: token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:      "token auth check: token timeout",
	ERROR_AUTH_TOKEN:                    "token auth: token error",
	ERROR_PRODUCTS_CURSOR_TOKEN_INVALID: "products cursor token invalid",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
