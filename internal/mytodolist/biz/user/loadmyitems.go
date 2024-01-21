package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

func (b *userBiz) LoadMyItems(ctx context.Context, req *v1.CommonRequestWizPagination, username string) (*v1.CommonResponseWizItemsAndPagination, error) {
	cursor := req.Pagination
	// if the cursor is nil , it means it's the first request ;
	if cursor == "" {
		items, npage, err := b.ds.Users().LoadMyItems(0, 10, username)
		if err != nil {
			return nil, errno.ErrLoadMyItemsFailed
		}
		resp := &v1.CommonResponseWizItemsAndPagination{
			Items:      items,
			Pagination: string(npage.PageEncode()),
		}
		return resp, nil
	} else {
		// it means the cursor  is not null
		// assert the cursor to Token
		c := token.Token(cursor)
		// get page from
		p := c.PageDecode()
		items, npage, err := b.ds.Users().LoadMyItems(p.NextID, int(p.PageSize), username)
		if err != nil {
			return nil, errno.ErrLoadMyItemsFailed
		}
		resp := &v1.CommonResponseWizItemsAndPagination{
			Items:      items,
			Pagination: string(npage.PageEncode()),
		}
		return resp, nil
	}
}
