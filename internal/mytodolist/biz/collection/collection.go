package collection

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

type CollectionBiz interface {
	AddItem(r *v1.CollectionAddItemRequest, username string) (*v1.CommonResponseWizMsg, error)
	Update(r *v1.CollectionUpdateRequest, username string) (*v1.CommonResponseWizMsg, error)
	Create(r *v1.CollectionCreateRequest, username string) (*v1.CommonResponseWizMsg, error)
	Delete(r *v1.DeleteCollectionWizIDRequest, username string) (*v1.CommonResponseWizMsg, error)
	DeleteItem(r *v1.CollectionDeleteItemRequest, username string) (*v1.CommonResponseWizMsg, error)
	LoadItems(r *v1.CollectionLoadItemsWizIdRequest, username string) (*v1.CollectionLoadItemsResp, error)
}

type collectionBiz struct {
	ds store.Istore
}

func New(ds store.Istore) *collectionBiz {
	return &collectionBiz{ds: ds}
}
