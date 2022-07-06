package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZCategoriesMgr struct {
	*_BaseMgr
}

// ZCategoriesMgr open func
func ZCategoriesMgr(db *gorm.DB) *_ZCategoriesMgr {
	if db == nil {
		panic(fmt.Errorf("ZCategoriesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZCategoriesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_categories"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZCategoriesMgr) GetTableName() string {
	return "z_categories"
}

// Reset 重置gorm会话
func (obj *_ZCategoriesMgr) Reset() *_ZCategoriesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZCategoriesMgr) Get() (result ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZCategoriesMgr) Gets() (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZCategoriesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZCategoriesMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 分类名
func (obj *_ZCategoriesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDisplayName display_name获取 分类别名
func (obj *_ZCategoriesMgr) WithDisplayName(displayName string) Option {
	return optionFunc(func(o *options) { o.query["display_name"] = displayName })
}

// WithSeoDesc seo_desc获取 seo描述
func (obj *_ZCategoriesMgr) WithSeoDesc(seoDesc string) Option {
	return optionFunc(func(o *options) { o.query["seo_desc"] = seoDesc })
}

// WithParentID parent_id获取 父类ID
func (obj *_ZCategoriesMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithCreatedAt created_at获取
func (obj *_ZCategoriesMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZCategoriesMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZCategoriesMgr) GetByOption(opts ...Option) (result ZCategories, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZCategoriesMgr) GetByOptions(opts ...Option) (results []*ZCategories, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZCategoriesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZCategories, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where(options.query)
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
func (obj *_ZCategoriesMgr) GetFromID(id uint) (result ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZCategoriesMgr) GetBatchFromID(ids []uint) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名
func (obj *_ZCategoriesMgr) GetFromName(name string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分类名
func (obj *_ZCategoriesMgr) GetBatchFromName(names []string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDisplayName 通过display_name获取内容 分类别名
func (obj *_ZCategoriesMgr) GetFromDisplayName(displayName string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`display_name` = ?", displayName).Find(&results).Error

	return
}

// GetBatchFromDisplayName 批量查找 分类别名
func (obj *_ZCategoriesMgr) GetBatchFromDisplayName(displayNames []string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`display_name` IN (?)", displayNames).Find(&results).Error

	return
}

// GetFromSeoDesc 通过seo_desc获取内容 seo描述
func (obj *_ZCategoriesMgr) GetFromSeoDesc(seoDesc string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`seo_desc` = ?", seoDesc).Find(&results).Error

	return
}

// GetBatchFromSeoDesc 批量查找 seo描述
func (obj *_ZCategoriesMgr) GetBatchFromSeoDesc(seoDescs []string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`seo_desc` IN (?)", seoDescs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父类ID
func (obj *_ZCategoriesMgr) GetFromParentID(parentID int) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父类ID
func (obj *_ZCategoriesMgr) GetBatchFromParentID(parentIDs []int) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZCategoriesMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZCategoriesMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZCategoriesMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZCategoriesMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZCategoriesMgr) FetchByPrimaryKey(id uint) (result ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByZCategoriesDisplayNameIndex  获取多个内容
func (obj *_ZCategoriesMgr) FetchIndexByZCategoriesDisplayNameIndex(displayName string) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`display_name` = ?", displayName).Find(&results).Error

	return
}

// FetchIndexByZCategoriesParentIDIndex  获取多个内容
func (obj *_ZCategoriesMgr) FetchIndexByZCategoriesParentIDIndex(parentID int) (results []*ZCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZCategories{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}
