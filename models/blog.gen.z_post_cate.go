package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ZPostCateMgr struct {
	*_BaseMgr
}

// ZPostCateMgr open func
func ZPostCateMgr(db *gorm.DB) *_ZPostCateMgr {
	if db == nil {
		panic(fmt.Errorf("ZPostCateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZPostCateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_post_cate"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZPostCateMgr) GetTableName() string {
	return "z_post_cate"
}

// Reset 重置gorm会话
func (obj *_ZPostCateMgr) Reset() *_ZPostCateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZPostCateMgr) Get() (result ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZPostCateMgr) Gets() (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZPostCateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZPostCateMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPostID post_id获取 文章ID
func (obj *_ZPostCateMgr) WithPostID(postID int) Option {
	return optionFunc(func(o *options) { o.query["post_id"] = postID })
}

// WithCateID cate_id获取 分类ID
func (obj *_ZPostCateMgr) WithCateID(cateID int) Option {
	return optionFunc(func(o *options) { o.query["cate_id"] = cateID })
}

// GetByOption 功能选项模式获取
func (obj *_ZPostCateMgr) GetByOption(opts ...Option) (result ZPostCate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZPostCateMgr) GetByOptions(opts ...Option) (results []*ZPostCate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZPostCateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZPostCate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where(options.query)
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
func (obj *_ZPostCateMgr) GetFromID(id uint) (result ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZPostCateMgr) GetBatchFromID(ids []uint) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPostID 通过post_id获取内容 文章ID
func (obj *_ZPostCateMgr) GetFromPostID(postID int) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`post_id` = ?", postID).Find(&results).Error

	return
}

// GetBatchFromPostID 批量查找 文章ID
func (obj *_ZPostCateMgr) GetBatchFromPostID(postIDs []int) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`post_id` IN (?)", postIDs).Find(&results).Error

	return
}

// GetFromCateID 通过cate_id获取内容 分类ID
func (obj *_ZPostCateMgr) GetFromCateID(cateID int) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`cate_id` = ?", cateID).Find(&results).Error

	return
}

// GetBatchFromCateID 批量查找 分类ID
func (obj *_ZPostCateMgr) GetBatchFromCateID(cateIDs []int) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`cate_id` IN (?)", cateIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZPostCateMgr) FetchByPrimaryKey(id uint) (result ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByZPostCatePostIDIndex  获取多个内容
func (obj *_ZPostCateMgr) FetchIndexByZPostCatePostIDIndex(postID int) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`post_id` = ?", postID).Find(&results).Error

	return
}

// FetchIndexByZPostCateCateIDIndex  获取多个内容
func (obj *_ZPostCateMgr) FetchIndexByZPostCateCateIDIndex(cateID int) (results []*ZPostCate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPostCate{}).Where("`cate_id` = ?", cateID).Find(&results).Error

	return
}
