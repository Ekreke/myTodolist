package errno

var (
	ErrDeleteItemFailed  = &Errno{HTTP: 500, Code: "InternalError.DeleteItemFailed", Message: "Delete my items failed."}
	ErrLoadItemsFailed   = &Errno{HTTP: 500, Code: "InternalError.LoadItemsFailed", Message: "Load items failed."}
	ErrLoadMyItemsFailed = &Errno{HTTP: 500, Code: "InternalError.LoadMyItemsFailed", Message: "Load my items failed."}
	ErrCreatItemsFailed  = &Errno{HTTP: 500, Code: "InternalError.CreatItemsFailed", Message: "Creat my items failed."}
	GetItemInfofailed    = &Errno{HTTP: 500, Code: "InternalError.GetItemInfofailed", Message: "Get item info failed."}
	SetItemDoneFailed    = &Errno{HTTP: 500, Code: "InternalError.SetItemDoneFailed", Message: "Set item done failed."}
	ErrUpdateItemFailed  = &Errno{HTTP: 500, Code: "InternalError.UpdateItemFailed", Message: "Update my items failed."}
	SetItemUnDoneFailed  = &Errno{HTTP: 500, Code: "InternalError.SetItemUnDoneFailed", Message: "Set item undone failed."}
)
