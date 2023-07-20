package dao

import (
	"backend/conf/mysql"
	"backend/internal/dao/query"
	"backend/internal/models"
	"backend/pkg/constants"
	"context"
	"strconv"
)

func SaveResult(ctx context.Context, loserId, winnerId string) error {

	intLoserId, _ := strconv.Atoi(loserId)
	intWinnerId, _ := strconv.Atoi(winnerId)

	err := mysql.Q.Transaction(func(tx *query.Query) error {
		userTx := tx.User
		// 更新输家
		_, err := userTx.WithContext(ctx).
			Where(userTx.ID.Eq(int32(intLoserId))).Update(userTx.Rating, userTx.Rating.Add(-3))
		if err != nil {
			return err
		}
		_, err = userTx.WithContext(ctx).
			Where(userTx.ID.Eq(int32(intWinnerId))).Update(userTx.Rating, userTx.Rating.Add(5))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetRankByLimit(ctx context.Context, page int) (users []*models.User, count int64, err error) {

	err = mysql.Q.Transaction(func(tx *query.Query) error {
		userTx := tx.User
		count, err = userTx.WithContext(ctx).Count()
		if err != nil {
			return err
		}

		users, err = userTx.WithContext(ctx).Order(userTx.Rating.Desc()).Limit(constants.RankPageSize).Offset((page - 1) * constants.RankPageSize).Find()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, -1, err
	}

	return users, count, nil
}
