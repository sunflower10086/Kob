package impl

import (
	"backend/conf/mysql"
	"backend/internal"
	"backend/internal/dao"
	"backend/internal/record"
	"backend/pkg/constants"
	"backend/pkg/ioc"
	"context"
	"log"
	"os"
)

var (
	impl = &RecordServiceImpl{}
)

func init() {
	ioc.RegistryImpl(impl)
}

type RecordServiceImpl struct {
	L *log.Logger
}

func (r *RecordServiceImpl) Config() {
	r.L = log.New(os.Stderr, "  [record] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (r *RecordServiceImpl) Name() string {
	return internal.RecordServiceName
}

func (r *RecordServiceImpl) GetList(ctx context.Context, page int) (*record.GetListResp, error) {
	records, count, err := dao.GetRecordByLimit(ctx, page)
	if err != nil {
		return nil, err
	}
	var resp record.GetListResp
	items := make([]record.Record, 0, constants.RecordPageSize)
	for _, rec := range records {
		userDao := mysql.Q.User
		userA, _ := userDao.WithContext(ctx).Where(userDao.ID.Eq(rec.AID)).First()
		userB, _ := userDao.WithContext(ctx).Where(userDao.ID.Eq(rec.BID)).First()

		result := "平局"
		if rec.Loser == "A" {
			result = "A输"
		} else if rec.Loser == "B" {
			result = "B输"
		}

		item := record.Record{
			APhoto:    userA.Photo,
			AUserName: userA.Username,
			BPhoto:    userB.Photo,
			BUserName: userB.Username,
			Result:    result,
			Record:    rec,
		}

		items = append(items, item)
	}

	resp.Records = &items
	resp.RecordsCount = count

	return &resp, nil
}
