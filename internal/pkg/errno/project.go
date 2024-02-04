package errno

var (
	ErrProjectJoin  = &Errno{HTTP: 400, Code: "FailedOperation.ProjectJoin", Message: "Project join failed."}
	ErrProjectInfo  = &Errno{HTTP: 404, Code: "ResourceNotFound.ProjectInfo", Message: "Project info not found."}
	ErrListProjects = &Errno{HTTP: 404, Code: "ResourceNotFound.ListProjects", Message: "List projects not found."}
)
