package e

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "invalid params",
	ERROR_NOT_EXIST_USER: "user not exist",
	ERROR_PASSWORD:       "password error",
	ERROR_USER_EXIST:     "user exist",
	ERROR_USER_NOT_LOGIN: "user not login",
	ERROR_DB:             "database error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
