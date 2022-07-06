package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZSystemsMgr struct {
	*_BaseMgr
}

// ZSystemsMgr open func
func ZSystemsMgr(db *gorm.DB) *_ZSystemsMgr {
	if db == nil {
		panic(fmt.Errorf("ZSystemsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZSystemsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_systems"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZSystemsMgr) GetTableName() string {
	return "z_systems"
}

// Reset 重置gorm会话
func (obj *_ZSystemsMgr) Reset() *_ZSystemsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZSystemsMgr) Get() (result ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZSystemsMgr) Gets() (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZSystemsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZSystemsMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTheme theme获取 主题
func (obj *_ZSystemsMgr) WithTheme(theme int8) Option {
	return optionFunc(func(o *options) { o.query["theme"] = theme })
}

// WithTitle title获取 网站title
func (obj *_ZSystemsMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithKeywords keywords获取 网站关键字
func (obj *_ZSystemsMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

// WithDescription description获取 网站描述
func (obj *_ZSystemsMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithRecordNumber record_number获取 备案号
func (obj *_ZSystemsMgr) WithRecordNumber(recordNumber string) Option {
	return optionFunc(func(o *options) { o.query["record_number"] = recordNumber })
}

// WithCreatedAt created_at获取
func (obj *_ZSystemsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZSystemsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZSystemsMgr) GetByOption(opts ...Option) (result ZSystems, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZSystemsMgr) GetByOptions(opts ...Option) (results []*ZSystems, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZSystemsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZSystems, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where(options.query)
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
func (obj *_ZSystemsMgr) GetFromID(id uint) (result ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZSystemsMgr) GetBatchFromID(ids []uint) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTheme 通过theme获取内容 主题
func (obj *_ZSystemsMgr) GetFromTheme(theme int8) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`theme` = ?", theme).Find(&results).Error

	return
}

// GetBatchFromTheme 批量查找 主题
func (obj *_ZSystemsMgr) GetBatchFromTheme(themes []int8) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`theme` IN (?)", themes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 网站title
func (obj *_ZSystemsMgr) GetFromTitle(title string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 网站title
func (obj *_ZSystemsMgr) GetBatchFromTitle(titles []string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromKeywords 通过keywords获取内容 网站关键字
func (obj *_ZSystemsMgr) GetFromKeywords(keywords string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`keywords` = ?", keywords).Find(&results).Error

	return
}

// GetBatchFromKeywords 批量查找 网站关键字
func (obj *_ZSystemsMgr) GetBatchFromKeywords(keywordss []string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`keywords` IN (?)", keywordss).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 网站描述
func (obj *_ZSystemsMgr) GetFromDescription(description string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 网站描述
func (obj *_ZSystemsMgr) GetBatchFromDescription(descriptions []string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromRecordNumber 通过record_number获取内容 备案号
func (obj *_ZSystemsMgr) GetFromRecordNumber(recordNumber string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`record_number` = ?", recordNumber).Find(&results).Error

	return
}

// GetBatchFromRecordNumber 批量查找 备案号
func (obj *_ZSystemsMgr) GetBatchFromRecordNumber(recordNumbers []string) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`record_number` IN (?)", recordNumbers).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZSystemsMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZSystemsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZSystemsMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZSystemsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZSystemsMgr) FetchByPrimaryKey(id uint) (result ZSystems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZSystems{}).Where("`id` = ?", id).First(&result).Error

	return
}
