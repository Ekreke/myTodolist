package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

func (b *userBiz) LoadNodes(ctx context.Context, req *v1.CommonRequestWizPagination, username string) (*v1.CommonResponseWizItemsAndPagination, error) {
	cursor := req.Pagination
	// if the cursor is nil , it means it's the first request ;
	if cursor == "" {
		items, npage, err := b.ds.Users().LoadNodes(0, 50, username)
		if err != nil {
			return nil, errno.ErrLoadNodesFailed
		}
		resp := &v1.CommonResponseWizItemsAndPagination{
			Items:      items,
			Pagination: string(npage.PageEncode()),
		}
		return resp, nil
	} else {
		c := token.Token(cursor)
		p := c.PageDecode()
		items, npage, err := b.ds.Users().LoadNodes(p.NextID, int(p.PageSize), username)
		if err != nil {
			return nil, errno.ErrLoadNodesFailed
		}
		resp := &v1.CommonResponseWizItemsAndPagination{
			Items:      items,
			Pagination: string(npage.PageEncode()),
		}
		return resp, nil
	}
}
