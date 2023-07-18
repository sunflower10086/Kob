// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"backend/internal/models"
)

func newBot(db *gorm.DB, opts ...gen.DOOption) bot {
	_bot := bot{}

	_bot.botDo.UseDB(db, opts...)
	_bot.botDo.UseModel(&models.Bot{})

	tableName := _bot.botDo.TableName()
	_bot.ALL = field.NewAsterisk(tableName)
	_bot.ID = field.NewInt32(tableName, "id")
	_bot.UserID = field.NewInt32(tableName, "user_id")
	_bot.Title = field.NewString(tableName, "title")
	_bot.Description = field.NewString(tableName, "description")
	_bot.Code = field.NewString(tableName, "code")
	_bot.Createtime = field.NewTime(tableName, "createtime")
	_bot.Modifytime = field.NewTime(tableName, "modifytime")

	_bot.fillFieldMap()

	return _bot
}

type bot struct {
	botDo

	ALL         field.Asterisk
	ID          field.Int32
	UserID      field.Int32
	Title       field.String
	Description field.String
	Code        field.String
	Createtime  field.Time
	Modifytime  field.Time

	fieldMap map[string]field.Expr
}

func (b bot) Table(newTableName string) *bot {
	b.botDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bot) As(alias string) *bot {
	b.botDo.DO = *(b.botDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bot) updateTableName(table string) *bot {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt32(table, "id")
	b.UserID = field.NewInt32(table, "user_id")
	b.Title = field.NewString(table, "title")
	b.Description = field.NewString(table, "description")
	b.Code = field.NewString(table, "code")
	b.Createtime = field.NewTime(table, "createtime")
	b.Modifytime = field.NewTime(table, "modifytime")

	b.fillFieldMap()

	return b
}

func (b *bot) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bot) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 7)
	b.fieldMap["id"] = b.ID
	b.fieldMap["user_id"] = b.UserID
	b.fieldMap["title"] = b.Title
	b.fieldMap["description"] = b.Description
	b.fieldMap["code"] = b.Code
	b.fieldMap["createtime"] = b.Createtime
	b.fieldMap["modifytime"] = b.Modifytime
}

func (b bot) clone(db *gorm.DB) bot {
	b.botDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bot) replaceDB(db *gorm.DB) bot {
	b.botDo.ReplaceDB(db)
	return b
}

type botDo struct{ gen.DO }

type IBotDo interface {
	gen.SubQuery
	Debug() IBotDo
	WithContext(ctx context.Context) IBotDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBotDo
	WriteDB() IBotDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBotDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBotDo
	Not(conds ...gen.Condition) IBotDo
	Or(conds ...gen.Condition) IBotDo
	Select(conds ...field.Expr) IBotDo
	Where(conds ...gen.Condition) IBotDo
	Order(conds ...field.Expr) IBotDo
	Distinct(cols ...field.Expr) IBotDo
	Omit(cols ...field.Expr) IBotDo
	Join(table schema.Tabler, on ...field.Expr) IBotDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBotDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBotDo
	Group(cols ...field.Expr) IBotDo
	Having(conds ...gen.Condition) IBotDo
	Limit(limit int) IBotDo
	Offset(offset int) IBotDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBotDo
	Unscoped() IBotDo
	Create(values ...*models.Bot) error
	CreateInBatches(values []*models.Bot, batchSize int) error
	Save(values ...*models.Bot) error
	First() (*models.Bot, error)
	Take() (*models.Bot, error)
	Last() (*models.Bot, error)
	Find() ([]*models.Bot, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Bot, err error)
	FindInBatches(result *[]*models.Bot, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Bot) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBotDo
	Assign(attrs ...field.AssignExpr) IBotDo
	Joins(fields ...field.RelationField) IBotDo
	Preload(fields ...field.RelationField) IBotDo
	FirstOrInit() (*models.Bot, error)
	FirstOrCreate() (*models.Bot, error)
	FindByPage(offset int, limit int) (result []*models.Bot, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBotDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b botDo) Debug() IBotDo {
	return b.withDO(b.DO.Debug())
}

func (b botDo) WithContext(ctx context.Context) IBotDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b botDo) ReadDB() IBotDo {
	return b.Clauses(dbresolver.Read)
}

func (b botDo) WriteDB() IBotDo {
	return b.Clauses(dbresolver.Write)
}

func (b botDo) Session(config *gorm.Session) IBotDo {
	return b.withDO(b.DO.Session(config))
}

func (b botDo) Clauses(conds ...clause.Expression) IBotDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b botDo) Returning(value interface{}, columns ...string) IBotDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b botDo) Not(conds ...gen.Condition) IBotDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b botDo) Or(conds ...gen.Condition) IBotDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b botDo) Select(conds ...field.Expr) IBotDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b botDo) Where(conds ...gen.Condition) IBotDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b botDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBotDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b botDo) Order(conds ...field.Expr) IBotDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b botDo) Distinct(cols ...field.Expr) IBotDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b botDo) Omit(cols ...field.Expr) IBotDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b botDo) Join(table schema.Tabler, on ...field.Expr) IBotDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b botDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBotDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b botDo) RightJoin(table schema.Tabler, on ...field.Expr) IBotDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b botDo) Group(cols ...field.Expr) IBotDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b botDo) Having(conds ...gen.Condition) IBotDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b botDo) Limit(limit int) IBotDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b botDo) Offset(offset int) IBotDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b botDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBotDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b botDo) Unscoped() IBotDo {
	return b.withDO(b.DO.Unscoped())
}

func (b botDo) Create(values ...*models.Bot) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b botDo) CreateInBatches(values []*models.Bot, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b botDo) Save(values ...*models.Bot) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b botDo) First() (*models.Bot, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Bot), nil
	}
}

func (b botDo) Take() (*models.Bot, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Bot), nil
	}
}

func (b botDo) Last() (*models.Bot, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Bot), nil
	}
}

func (b botDo) Find() ([]*models.Bot, error) {
	result, err := b.DO.Find()
	return result.([]*models.Bot), err
}

func (b botDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Bot, err error) {
	buf := make([]*models.Bot, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b botDo) FindInBatches(result *[]*models.Bot, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b botDo) Attrs(attrs ...field.AssignExpr) IBotDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b botDo) Assign(attrs ...field.AssignExpr) IBotDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b botDo) Joins(fields ...field.RelationField) IBotDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b botDo) Preload(fields ...field.RelationField) IBotDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b botDo) FirstOrInit() (*models.Bot, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Bot), nil
	}
}

func (b botDo) FirstOrCreate() (*models.Bot, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Bot), nil
	}
}

func (b botDo) FindByPage(offset int, limit int) (result []*models.Bot, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b botDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b botDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b botDo) Delete(models ...*models.Bot) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *botDo) withDO(do gen.Dao) *botDo {
	b.DO = *do.(*gen.DO)
	return b
}
