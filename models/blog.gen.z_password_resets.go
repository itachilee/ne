package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZPasswordResetsMgr struct {
	*_BaseMgr
}

// ZPasswordResetsMgr open func
func ZPasswordResetsMgr(db *gorm.DB) *_ZPasswordResetsMgr {
	if db == nil {
		panic(fmt.Errorf("ZPasswordResetsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZPasswordResetsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_password_resets"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZPasswordResetsMgr) GetTableName() string {
	return "z_password_resets"
}

// Reset 重置gorm会话
func (obj *_ZPasswordResetsMgr) Reset() *_ZPasswordResetsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZPasswordResetsMgr) Get() (result ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZPasswordResetsMgr) Gets() (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZPasswordResetsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithEmail email获取
func (obj *_ZPasswordResetsMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithToken token获取
func (obj *_ZPasswordResetsMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// WithCreatedAt created_at获取
func (obj *_ZPasswordResetsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZPasswordResetsMgr) GetByOption(opts ...Option) (result ZPasswordResets, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZPasswordResetsMgr) GetByOptions(opts ...Option) (results []*ZPasswordResets, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZPasswordResetsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZPasswordResets, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromEmail 通过email获取内容
func (obj *_ZPasswordResetsMgr) GetFromEmail(email string) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_ZPasswordResetsMgr) GetBatchFromEmail(emails []string) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromToken 通过token获取内容
func (obj *_ZPasswordResetsMgr) GetFromToken(token string) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`token` = ?", token).Find(&results).Error

	return
}

// GetBatchFromToken 批量查找
func (obj *_ZPasswordResetsMgr) GetBatchFromToken(tokens []string) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`token` IN (?)", tokens).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZPasswordResetsMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZPasswordResetsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByZPasswordResetsEmailIndex  获取多个内容
func (obj *_ZPasswordResetsMgr) FetchIndexByZPasswordResetsEmailIndex(email string) (results []*ZPasswordResets, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPasswordResets{}).Where("`email` = ?", email).Find(&results).Error

	return
}
