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

func GetRankByLimit(ctx context.Context, page int) ([]*models.User, error) {
	userDao := mysql.Q.User
	users, err := userDao.WithContext(ctx).Order(userDao.Rating.Desc()).Limit(constants.RankPageSize).Offset((page - 1) * constants.RankPageSize).Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}
