package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StockBMgr struct {
	*_BaseMgr
}

// StockBMgr open func
func StockBMgr(db *gorm.DB) *_StockBMgr {
	if db == nil {
		panic(fmt.Errorf("StockBMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockBMgr{_BaseMgr: &_BaseMgr{DB: db.Table("stock_b"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StockBMgr) GetTableName() string {
	return "stock_b"
}

// Reset 重置gorm会话
func (obj *_StockBMgr) Reset() *_StockBMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_StockBMgr) Get() (result StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StockBMgr) Gets() (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_StockBMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(StockB{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCompanyCode company_code获取
func (obj *_StockBMgr) WithCompanyCode(companyCode int) Option {
	return optionFunc(func(o *options) { o.query["company_code"] = companyCode })
}

// WithCompantName compant_name获取
func (obj *_StockBMgr) WithCompantName(compantName string) Option {
	return optionFunc(func(o *options) { o.query["compant_name"] = compantName })
}

// WithCode code获取
func (obj *_StockBMgr) WithCode(code int) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_StockBMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDate date获取
func (obj *_StockBMgr) WithDate(date string) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_StockBMgr) GetByOption(opts ...Option) (result StockB, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StockBMgr) GetByOptions(opts ...Option) (results []*StockB, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_StockBMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]StockB, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(StockB{}).Where(options.query)
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
func (obj *_StockBMgr) GetFromCompanyCode(companyCode int) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`company_code` = ?", companyCode).Find(&results).Error

	return
}

// GetBatchFromCompanyCode 批量查找
func (obj *_StockBMgr) GetBatchFromCompanyCode(companyCodes []int) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`company_code` IN (?)", companyCodes).Find(&results).Error

	return
}

// GetFromCompantName 通过compant_name获取内容
func (obj *_StockBMgr) GetFromCompantName(compantName string) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`compant_name` = ?", compantName).Find(&results).Error

	return
}

// GetBatchFromCompantName 批量查找
func (obj *_StockBMgr) GetBatchFromCompantName(compantNames []string) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`compant_name` IN (?)", compantNames).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_StockBMgr) GetFromCode(code int) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_StockBMgr) GetBatchFromCode(codes []int) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_StockBMgr) GetFromName(name string) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_StockBMgr) GetBatchFromName(names []string) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_StockBMgr) GetFromDate(date string) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_StockBMgr) GetBatchFromDate(dates []string) (results []*StockB, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StockB{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
