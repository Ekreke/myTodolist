package errno

var (
	ErrProjectJoin          = &Errno{HTTP: 400, Code: "FailedOperation.ProjectJoin", Message: "Project join failed."}
	ErrProjectInfo          = &Errno{HTTP: 404, Code: "ResourceNotFound.ProjectInfo", Message: "Project info not found."}
	ErrListProjects         = &Errno{HTTP: 404, Code: "ResourceNotFound.ListProjects", Message: "List projects not found."}
	ErrProjectCreate        = &Errno{HTTP: 400, Code: "FailedOperation.ProjectCreate", Message: "Project create failed."}
	ErrProjectDelete        = &Errno{HTTP: 400, Code: "FailedOperation.ProjectDelete", Message: "Project delete failed."}
	ErrProjectDeleteNode    = &Errno{HTTP: 400, Code: "FailedOperation.ProjectDeleteNode", Message: "Project delete node failed."}
	ErrProjectAddNode       = &Errno{HTTP: 400, Code: "FailedOperation.ProjectAddNode", Message: "Project add node failed."}
	ErrListProjectsICreated = &Errno{HTTP: 400, Code: "FailedOperation.ListProjectsICreated", Message: "List projects I created failed."}
	ErrProjectUpdate        = &Errno{HTTP: 400, Code: "FailedOperation.ProjectUpdate", Message: "Project update failed."}
	ErrProjectUpdateNode    = &Errno{HTTP: 400, Code: "FailedOperation.ProjectUpdateNode", Message: "Project update node failed."}
)
