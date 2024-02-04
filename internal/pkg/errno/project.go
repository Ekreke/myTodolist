package errno

var (
	ErrProjectJoin = &Errno{HTTP: 400, Code: "FailedOperation.ProjectJoin", Message: "Project join failed."}
)
