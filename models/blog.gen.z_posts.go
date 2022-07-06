package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZPostsMgr struct {
	*_BaseMgr
}

// ZPostsMgr open func
func ZPostsMgr(db *gorm.DB) *_ZPostsMgr {
	if db == nil {
		panic(fmt.Errorf("ZPostsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZPostsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_posts"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZPostsMgr) GetTableName() string {
	return "z_posts"
}

// Reset 重置gorm会话
func (obj *_ZPostsMgr) Reset() *_ZPostsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZPostsMgr) Get() (result ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZPostsMgr) Gets() (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZPostsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZPostsMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUID uid获取 uid
func (obj *_ZPostsMgr) WithUID(uid string) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithUserID user_id获取 用户ID
func (obj *_ZPostsMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithTitle title获取 标题
func (obj *_ZPostsMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithSummary summary获取 摘要
func (obj *_ZPostsMgr) WithSummary(summary string) Option {
	return optionFunc(func(o *options) { o.query["summary"] = summary })
}

// WithOriginal original获取 原文章内容
func (obj *_ZPostsMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithContent content获取 文章内容
func (obj *_ZPostsMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithPassword password获取 文章密码
func (obj *_ZPostsMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithDeletedAt deleted_at获取
func (obj *_ZPostsMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreatedAt created_at获取
func (obj *_ZPostsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZPostsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZPostsMgr) GetByOption(opts ...Option) (result ZPosts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZPostsMgr) GetByOptions(opts ...Option) (results []*ZPosts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZPostsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZPosts, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where(options.query)
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
func (obj *_ZPostsMgr) GetFromID(id uint) (result ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZPostsMgr) GetBatchFromID(ids []uint) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUID 通过uid获取内容 uid
func (obj *_ZPostsMgr) GetFromUID(uid string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUID 批量查找 uid
func (obj *_ZPostsMgr) GetBatchFromUID(uids []string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_ZPostsMgr) GetFromUserID(userID int) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_ZPostsMgr) GetBatchFromUserID(userIDs []int) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_ZPostsMgr) GetFromTitle(title string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_ZPostsMgr) GetBatchFromTitle(titles []string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromSummary 通过summary获取内容 摘要
func (obj *_ZPostsMgr) GetFromSummary(summary string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`summary` = ?", summary).Find(&results).Error

	return
}

// GetBatchFromSummary 批量查找 摘要
func (obj *_ZPostsMgr) GetBatchFromSummary(summarys []string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`summary` IN (?)", summarys).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 原文章内容
func (obj *_ZPostsMgr) GetFromOriginal(original string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 原文章内容
func (obj *_ZPostsMgr) GetBatchFromOriginal(originals []string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 文章内容
func (obj *_ZPostsMgr) GetFromContent(content string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 文章内容
func (obj *_ZPostsMgr) GetBatchFromContent(contents []string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 文章密码
func (obj *_ZPostsMgr) GetFromPassword(password string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 文章密码
func (obj *_ZPostsMgr) GetBatchFromPassword(passwords []string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容
func (obj *_ZPostsMgr) GetFromDeletedAt(deletedAt time.Time) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找
func (obj *_ZPostsMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZPostsMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZPostsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZPostsMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZPostsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZPostsMgr) FetchByPrimaryKey(id uint) (result ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByZPostsUIDIndex  获取多个内容
func (obj *_ZPostsMgr) FetchIndexByZPostsUIDIndex(uid string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// FetchIndexByZPostsUserIDIndex  获取多个内容
func (obj *_ZPostsMgr) FetchIndexByZPostsUserIDIndex(userID int) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByZPostsTitleIndex  获取多个内容
func (obj *_ZPostsMgr) FetchIndexByZPostsTitleIndex(title string) (results []*ZPosts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZPosts{}).Where("`title` = ?", title).Find(&results).Error

	return
}
