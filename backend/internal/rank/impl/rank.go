package impl

import (
	"backend/internal"
	"backend/internal/dao"
	"backend/internal/rank"
	"backend/pkg/constants"
	"backend/pkg/ioc"
	"context"
	"log"
	"os"
)

var impl = &RankServiceImpl{}

func init() {
	ioc.RegistryImpl(impl)
}

type RankServiceImpl struct {
	L *log.Logger
}

func (r *RankServiceImpl) Config() {
	r.L = log.New(os.Stderr, "  [rank] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (r *RankServiceImpl) Name() string {
	return internal.RankServiceName
}

func (r *RankServiceImpl) GetRankList(ctx context.Context, page int) (*rank.GetRankListResp, error) {
	users, err := dao.GetRankByLimit(ctx, page)
	if err != nil {
		return nil, err
	}
	var resp rank.GetRankListResp
	var items []rank.RankUser

	for i, user := range users {

		item := rank.RankUser{
			Id:       user.ID,
			Photo:    user.Photo,
			UserName: user.Username,
			Rating:   user.Rating,
			Number:   (page-1)*constants.RankPageSize + i + 1,
		}

		items = append(items, item)
	}

	resp.Users = &items
	resp.UserCount = len(items)
	return &resp, nil
}
