package dao

import (
	"backend/conf/mysql"
	"backend/internal/dao/query"
	"backend/internal/models"
	"backend/pkg/constants"
	"context"
)

func GetRecordByLimit(ctx context.Context, page int) ([]*models.Record, int64, error) {
	var records []*models.Record
	var count int64
	var err error

	err = mysql.Q.Transaction(func(tx *query.Query) error {
		recordTx := tx.Record
		count, err = recordTx.WithContext(ctx).Count()
		if err != nil {
			return err
		}

		records, err = recordTx.WithContext(ctx).Order(recordTx.ID).Limit(constants.RecordPageSize).Offset((page - 1) * constants.RecordPageSize).Find()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, -1, err
	}

	return records, count, nil
}
