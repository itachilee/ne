package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StockAMgr struct {
	*_BaseMgr
}

// StockAMgr open func
func StockAMgr(db *gorm.DB) *_StockAMgr {
	if db == nil {
		panic(fmt.Errorf("StockAMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockAMgr{_BaseMgr: &_BaseMgr{DB: db.Table("stock_a"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StockAMgr) GetTableName() string {
	return "stock_a"
}

// Reset 重置gorm会话
func (obj *_StockAMgr) Reset() *_StockAMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_StockAMgr) Get() (result StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StockAMgr) Gets() (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_StockAMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(StockA{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCompanyCode company_code获取
func (obj *_StockAMgr) WithCompanyCode(companyCode int) Option {
	return optionFunc(func(o *options) { o.query["company_code"] = companyCode })
}

// WithCompantName compant_name获取
func (obj *_StockAMgr) WithCompantName(compantName string) Option {
	return optionFunc(func(o *options) { o.query["compant_name"] = compantName })
}

// WithCode code获取
func (obj *_StockAMgr) WithCode(code int) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_StockAMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDate date获取
func (obj *_StockAMgr) WithDate(date string) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_StockAMgr) GetByOption(opts ...Option) (result StockA, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StockAMgr) GetByOptions(opts ...Option) (results []*StockA, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_StockAMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]StockA, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(StockA{}).Where(options.query)
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

// GetFromCompanyCode 通过company_code获取内容
func (obj *_StockAMgr) GetFromCompanyCode(companyCode int) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`company_code` = ?", companyCode).Find(&results).Error

	return
}

// GetBatchFromCompanyCode 批量查找
func (obj *_StockAMgr) GetBatchFromCompanyCode(companyCodes []int) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`company_code` IN (?)", companyCodes).Find(&results).Error

	return
}

// GetFromCompantName 通过compant_name获取内容
func (obj *_StockAMgr) GetFromCompantName(compantName string) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`compant_name` = ?", compantName).Find(&results).Error

	return
}

// GetBatchFromCompantName 批量查找
func (obj *_StockAMgr) GetBatchFromCompantName(compantNames []string) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`compant_name` IN (?)", compantNames).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_StockAMgr) GetFromCode(code int) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_StockAMgr) GetBatchFromCode(codes []int) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_StockAMgr) GetFromName(name string) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_StockAMgr) GetBatchFromName(names []string) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_StockAMgr) GetFromDate(date string) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_StockAMgr) GetBatchFromDate(dates []string) (results []*StockA, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockA{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
