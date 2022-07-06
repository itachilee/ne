package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ZPostTagMgr struct {
	*_BaseMgr
}

// ZPostTagMgr open func
func ZPostTagMgr(db *gorm.DB) *_ZPostTagMgr {
	if db == nil {
		panic(fmt.Errorf("ZPostTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZPostTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_post_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZPostTagMgr) GetTableName() string {
	return "z_post_tag"
}

// Reset 重置gorm会话
func (obj *_ZPostTagMgr) Reset() *_ZPostTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZPostTagMgr) Get() (result ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZPostTagMgr) Gets() (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZPostTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZPostTagMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPostID post_id获取 文章ID
func (obj *_ZPostTagMgr) WithPostID(postID int) Option {
	return optionFunc(func(o *options) { o.query["post_id"] = postID })
}

// WithTagID tag_id获取 标签ID
func (obj *_ZPostTagMgr) WithTagID(tagID int) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// GetByOption 功能选项模式获取
func (obj *_ZPostTagMgr) GetByOption(opts ...Option) (result ZPostTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZPostTagMgr) GetByOptions(opts ...Option) (results []*ZPostTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZPostTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZPostTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where(options.query)
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
func (obj *_ZPostTagMgr) GetFromID(id uint) (result ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZPostTagMgr) GetBatchFromID(ids []uint) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPostID 通过post_id获取内容 文章ID
func (obj *_ZPostTagMgr) GetFromPostID(postID int) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`post_id` = ?", postID).Find(&results).Error

	return
}

// GetBatchFromPostID 批量查找 文章ID
func (obj *_ZPostTagMgr) GetBatchFromPostID(postIDs []int) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`post_id` IN (?)", postIDs).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容 标签ID
func (obj *_ZPostTagMgr) GetFromTagID(tagID int) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`tag_id` = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量查找 标签ID
func (obj *_ZPostTagMgr) GetBatchFromTagID(tagIDs []int) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZPostTagMgr) FetchByPrimaryKey(id uint) (result ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByZPostTagPostIDIndex  获取多个内容
func (obj *_ZPostTagMgr) FetchIndexByZPostTagPostIDIndex(postID int) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`post_id` = ?", postID).Find(&results).Error

	return
}

// FetchIndexByZPostTagTagIDIndex  获取多个内容
func (obj *_ZPostTagMgr) FetchIndexByZPostTagTagIDIndex(tagID int) (results []*ZPostTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostTag{}).Where("`tag_id` = ?", tagID).Find(&results).Error

	return
}
