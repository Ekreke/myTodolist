package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

func (b *userBiz) LoadMydayItems(ctx context.Context, req *v1.MydayRequest, username string) (*v1.MydayResponse, error) {
	cursor := req.Pagination
	resp := &v1.MydayResponse{}

	if cursor == "" {
		items, npage, err := b.ds.Users().GetMydayItems(0, 10, username)
		if err != nil {
			return nil, errno.ErrLoadMydayItemFailed
		}
		resp = &v1.MydayResponse{
			Items:       items,
			CursorToken: string(npage.PageEncode()),
		}
		return resp, nil
	} else {
		c := token.Token(cursor)
		p := c.PageDecode()
		items, npage, err := b.ds.Users().GetMydayItems(p.NextID, int(p.PageSize), username)
		if err != nil {
			return nil, errno.ErrLoadMydayItemFailed
		}
		resp = &v1.MydayResponse{
			Items:       items,
			CursorToken: string(npage.PageEncode()),
		}
	}
	return resp, nil
}
