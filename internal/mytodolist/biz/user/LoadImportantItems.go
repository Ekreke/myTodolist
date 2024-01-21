package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

func (b *userBiz) LoadImportantItems(ctx context.Context, req *v1.ImportantRequest, username string) (*v1.ImportantResponse, token.Token, error) {
	cursor := req.Pagination
	// if the cursor is nil , it means it's the first request ;
	if cursor == "" {
		items, npage, err := b.ds.Users().GetImportantItems(0, 10, username)
		if err != nil {
			return nil, "", errno.ErrLoadImportantItemFailed
		}
		resp := &v1.ImportantResponse{
			ImportantItems: items,
		}
		return resp, npage.PageEncode(), nil
	} else {

		c := token.Token(cursor)
		p := c.PageDecode()
		items, npage, err := b.ds.Users().GetImportantItems(p.NextID, int(p.PageSize), username)
		if err != nil {
			return nil, "", errno.ErrLoadImportantItemFailed
		}
		resp := &v1.ImportantResponse{
			ImportantItems: items,
		}
		return resp, npage.PageEncode(), nil
	}
}
