package errno

var (
	// ErrUserAlreadyExist 代表用户已经存在.
	ErrUserAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exist."}

	// ErrUserNotFound 表示未找到用户.
	ErrUserNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "User was not found."}

	// ErrPasswordIncorrect 表示密码不正确.
	ErrPasswordIncorrect = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password was incorrect."}

	ErrUserCreateFailed = &Errno{HTTP: 500, Code: "InternalError.UserCreateFailed", Message: "User create failed."}

	ErrUserUpdateFailed        = &Errno{HTTP: 500, Code: "InternalError.UserUpdateFailed", Message: "User update failed."}
	ErrLoadImportantItemFailed = &Errno{HTTP: 500, Code: "InternalError.LoadImportantItemFailed", Message: "Load important item failed."}
	ErrUpdatePwdFailed         = &Errno{HTTP: 500, Code: "InternalError.UpdatePwdFailed", Message: "Update pwd failed."}
	ErrPwdDuplicate            = &Errno{HTTP: 400, Code: "InvalidParameter.PwdDuplicate", Message: "Password was duplicate."}
)
