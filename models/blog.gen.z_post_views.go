package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ZPostViewsMgr struct {
	*_BaseMgr
}

// ZPostViewsMgr open func
func ZPostViewsMgr(db *gorm.DB) *_ZPostViewsMgr {
	if db == nil {
		panic(fmt.Errorf("ZPostViewsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZPostViewsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_post_views"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZPostViewsMgr) GetTableName() string {
	return "z_post_views"
}

// Reset 重置gorm会话
func (obj *_ZPostViewsMgr) Reset() *_ZPostViewsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZPostViewsMgr) Get() (result ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZPostViewsMgr) Gets() (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZPostViewsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZPostViewsMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPostID post_id获取 文章ID
func (obj *_ZPostViewsMgr) WithPostID(postID int) Option {
	return optionFunc(func(o *options) { o.query["post_id"] = postID })
}

// WithNum num获取 阅读次数
func (obj *_ZPostViewsMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// GetByOption 功能选项模式获取
func (obj *_ZPostViewsMgr) GetByOption(opts ...Option) (result ZPostViews, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZPostViewsMgr) GetByOptions(opts ...Option) (results []*ZPostViews, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZPostViewsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZPostViews, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where(options.query)
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
func (obj *_ZPostViewsMgr) GetFromID(id uint) (result ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZPostViewsMgr) GetBatchFromID(ids []uint) (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPostID 通过post_id获取内容 文章ID
func (obj *_ZPostViewsMgr) GetFromPostID(postID int) (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`post_id` = ?", postID).Find(&results).Error

	return
}

// GetBatchFromPostID 批量查找 文章ID
func (obj *_ZPostViewsMgr) GetBatchFromPostID(postIDs []int) (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`post_id` IN (?)", postIDs).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 阅读次数
func (obj *_ZPostViewsMgr) GetFromNum(num int) (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 阅读次数
func (obj *_ZPostViewsMgr) GetBatchFromNum(nums []int) (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZPostViewsMgr) FetchByPrimaryKey(id uint) (result ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByZPostViewsPostIDIndex  获取多个内容
func (obj *_ZPostViewsMgr) FetchIndexByZPostViewsPostIDIndex(postID int) (results []*ZPostViews, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostViews{}).Where("`post_id` = ?", postID).Find(&results).Error

	return
}
