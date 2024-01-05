package e

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "invalid params",
	ERROR_NOT_EXIST_USER: "user not exist",
	ERROR_PASSWORD:       "password error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
