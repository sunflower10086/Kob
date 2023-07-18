package dao

import (
	"backend/conf/mysql"
	"backend/internal/models"
	"backend/pkg/constants"
	"context"
)

func GetRecordByLimit(ctx context.Context, page int) ([]*models.Record, error) {
	recordDao := mysql.Q.Record
	records, err := recordDao.WithContext(ctx).Order(recordDao.ID).Limit(constants.RecordPageSize).Offset((page - 1) * constants.RecordPageSize).Find()
	if err != nil {
		return nil, err
	}

	return records, nil
}
