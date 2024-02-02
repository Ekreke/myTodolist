package mytodolist

type CollectionCreateRequest struct {
	Name string `form:"name"`
	Icon int64  `form:"icon"`
}

type CollectionUpdateRequest struct {
	CollectionId int64  `form:"collection_id"`
	Name         string `form:"name"`
	Icon         int64  `form:"icon"`
}

type CollectionAddItemRequest struct {
	ItemId       int64 `form:"item_id"`
	CollectionId int64 `form:"collection_id"`
}

type CollectionLoadItemsWizIdRequest struct {
	CollectionId int64 `form:"collection_id"`
}

type CollectionDeleteItemRequest struct {
	ItemId int64 `form:"item_id"`
}

type DeleteCollectionWizIDRequest struct {
	CollectionId int64 `form:"collection_id"`
}
