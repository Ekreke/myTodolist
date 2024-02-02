package collection

import v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"

func (c *collectionBiz) Update(r *v1.CollectionUpdateRequest, username string) (*v1.CommonResponseWizMsg, error) {
	resp, err := c.ds.Collection().Update(r.CollectionId, r.Icon, r.Name, username)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
