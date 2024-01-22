package errno

var (
	ErrUserAlreadyExist        = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exist."}
	ErrUserNotFound            = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "User was not found."}
	ErrPasswordIncorrect       = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password was incorrect."}
	ErrUserCreateFailed        = &Errno{HTTP: 500, Code: "InternalError.UserCreateFailed", Message: "User create failed."}
	ErrUserUpdateFailed        = &Errno{HTTP: 500, Code: "InternalError.UserUpdateFailed", Message: "User update failed."}
	ErrLoadImportantItemFailed = &Errno{HTTP: 500, Code: "InternalError.LoadImportantItemFailed", Message: "Load important item failed."}
	ErrUpdatePwdFailed         = &Errno{HTTP: 500, Code: "InternalError.UpdatePwdFailed", Message: "Update pwd failed."}
	ErrPwdDuplicate            = &Errno{HTTP: 400, Code: "InvalidParameter.PwdDuplicate", Message: "Password was duplicate."}
	ErrLoadMydayItemFailed     = &Errno{HTTP: 500, Code: "InternalError.LoadMydayItemFailed", Message: "Load myday item failed."}
	ErrLoadNodesFailed         = &Errno{HTTP: 500, Code: "InternalError.LoadNodesFailed", Message: "Load nodes failed."}
)
