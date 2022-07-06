package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZTagsMgr struct {
	*_BaseMgr
}

// ZTagsMgr open func
func ZTagsMgr(db *gorm.DB) *_ZTagsMgr {
	if db == nil {
		panic(fmt.Errorf("ZTagsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZTagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_tags"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZTagsMgr) GetTableName() string {
	return "z_tags"
}

// Reset 重置gorm会话
func (obj *_ZTagsMgr) Reset() *_ZTagsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZTagsMgr) Get() (result ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZTagsMgr) Gets() (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZTagsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZTags{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZTagsMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 标签名
func (obj *_ZTagsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDisplayName display_name获取 标签别名
func (obj *_ZTagsMgr) WithDisplayName(displayName string) Option {
	return optionFunc(func(o *options) { o.query["display_name"] = displayName })
}

// WithSeoDesc seo_desc获取 seo描述
func (obj *_ZTagsMgr) WithSeoDesc(seoDesc string) Option {
	return optionFunc(func(o *options) { o.query["seo_desc"] = seoDesc })
}

// WithNum num获取 被使用次数
func (obj *_ZTagsMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithCreatedAt created_at获取
func (obj *_ZTagsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZTagsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZTagsMgr) GetByOption(opts ...Option) (result ZTags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZTagsMgr) GetByOptions(opts ...Option) (results []*ZTags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZTagsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZTags, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where(options.query)
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
func (obj *_ZTagsMgr) GetFromID(id uint) (result ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZTagsMgr) GetBatchFromID(ids []uint) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 标签名
func (obj *_ZTagsMgr) GetFromName(name string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 标签名
func (obj *_ZTagsMgr) GetBatchFromName(names []string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDisplayName 通过display_name获取内容 标签别名
func (obj *_ZTagsMgr) GetFromDisplayName(displayName string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`display_name` = ?", displayName).Find(&results).Error

	return
}

// GetBatchFromDisplayName 批量查找 标签别名
func (obj *_ZTagsMgr) GetBatchFromDisplayName(displayNames []string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`display_name` IN (?)", displayNames).Find(&results).Error

	return
}

// GetFromSeoDesc 通过seo_desc获取内容 seo描述
func (obj *_ZTagsMgr) GetFromSeoDesc(seoDesc string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`seo_desc` = ?", seoDesc).Find(&results).Error

	return
}

// GetBatchFromSeoDesc 批量查找 seo描述
func (obj *_ZTagsMgr) GetBatchFromSeoDesc(seoDescs []string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`seo_desc` IN (?)", seoDescs).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 被使用次数
func (obj *_ZTagsMgr) GetFromNum(num int) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 被使用次数
func (obj *_ZTagsMgr) GetBatchFromNum(nums []int) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZTagsMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZTagsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZTagsMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZTagsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZTagsMgr) FetchByPrimaryKey(id uint) (result ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByZTagsDisplayNameIndex  获取多个内容
func (obj *_ZTagsMgr) FetchIndexByZTagsDisplayNameIndex(displayName string) (results []*ZTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZTags{}).Where("`display_name` = ?", displayName).Find(&results).Error

	return
}
