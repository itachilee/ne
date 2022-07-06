package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZCreditCardMgr struct {
	*_BaseMgr
}

// ZCreditCardMgr open func
func ZCreditCardMgr(db *gorm.DB) *_ZCreditCardMgr {
	if db == nil {
		panic(fmt.Errorf("ZCreditCardMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZCreditCardMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_credit_card"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZCreditCardMgr) GetTableName() string {
	return "z_credit_card"
}

// Reset 重置gorm会话
func (obj *_ZCreditCardMgr) Reset() *_ZCreditCardMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZCreditCardMgr) Get() (result ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZCreditCardMgr) Gets() (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZCreditCardMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZCreditCardMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNumber number获取
func (obj *_ZCreditCardMgr) WithNumber(number string) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithUserid userid获取
func (obj *_ZCreditCardMgr) WithUserid(userid int) Option {
	return optionFunc(func(o *options) { o.query["userid"] = userid })
}

// WithCreatedAt created_at获取
func (obj *_ZCreditCardMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZCreditCardMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZCreditCardMgr) GetByOption(opts ...Option) (result ZCreditCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZCreditCardMgr) GetByOptions(opts ...Option) (results []*ZCreditCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZCreditCardMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZCreditCard, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_ZCreditCardMgr) GetFromID(id uint) (result ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZCreditCardMgr) GetBatchFromID(ids []uint) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容
func (obj *_ZCreditCardMgr) GetFromNumber(number string) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`number` = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量查找
func (obj *_ZCreditCardMgr) GetBatchFromNumber(numbers []string) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`number` IN (?)", numbers).Find(&results).Error

	return
}

// GetFromUserid 通过userid获取内容
func (obj *_ZCreditCardMgr) GetFromUserid(userid int) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`userid` = ?", userid).Find(&results).Error

	return
}

// GetBatchFromUserid 批量查找
func (obj *_ZCreditCardMgr) GetBatchFromUserid(userids []int) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`userid` IN (?)", userids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZCreditCardMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZCreditCardMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZCreditCardMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZCreditCardMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZCreditCardMgr) FetchByPrimaryKey(id uint) (result ZCreditCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCreditCard{}).Where("`id` = ?", id).First(&result).Error

	return
}
