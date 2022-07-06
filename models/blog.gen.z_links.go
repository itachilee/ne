package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZLinksMgr struct {
	*_BaseMgr
}

// ZLinksMgr open func
func ZLinksMgr(db *gorm.DB) *_ZLinksMgr {
	if db == nil {
		panic(fmt.Errorf("ZLinksMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZLinksMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_links"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZLinksMgr) GetTableName() string {
	return "z_links"
}

// Reset 重置gorm会话
func (obj *_ZLinksMgr) Reset() *_ZLinksMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZLinksMgr) Get() (result ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZLinksMgr) Gets() (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZLinksMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZLinksMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 友链名
func (obj *_ZLinksMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithLink link获取 友链链接
func (obj *_ZLinksMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// WithOrder order获取 排序
func (obj *_ZLinksMgr) WithOrder(order int) Option {
	return optionFunc(func(o *options) { o.query["order"] = order })
}

// WithCreatedAt created_at获取
func (obj *_ZLinksMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZLinksMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZLinksMgr) GetByOption(opts ...Option) (result ZLinks, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZLinksMgr) GetByOptions(opts ...Option) (results []*ZLinks, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZLinksMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZLinks, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where(options.query)
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
func (obj *_ZLinksMgr) GetFromID(id uint) (result ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZLinksMgr) GetBatchFromID(ids []uint) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 友链名
func (obj *_ZLinksMgr) GetFromName(name string) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 友链名
func (obj *_ZLinksMgr) GetBatchFromName(names []string) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromLink 通过link获取内容 友链链接
func (obj *_ZLinksMgr) GetFromLink(link string) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`link` = ?", link).Find(&results).Error

	return
}

// GetBatchFromLink 批量查找 友链链接
func (obj *_ZLinksMgr) GetBatchFromLink(links []string) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`link` IN (?)", links).Find(&results).Error

	return
}

// GetFromOrder 通过order获取内容 排序
func (obj *_ZLinksMgr) GetFromOrder(order int) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`order` = ?", order).Find(&results).Error

	return
}

// GetBatchFromOrder 批量查找 排序
func (obj *_ZLinksMgr) GetBatchFromOrder(orders []int) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`order` IN (?)", orders).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZLinksMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZLinksMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZLinksMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZLinksMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZLinksMgr) FetchByPrimaryKey(id uint) (result ZLinks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZLinks{}).Where("`id` = ?", id).First(&result).Error

	return
}
