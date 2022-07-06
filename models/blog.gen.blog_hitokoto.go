package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BlogHitokotoMgr struct {
	*_BaseMgr
}

// BlogHitokotoMgr open func
func BlogHitokotoMgr(db *gorm.DB) *_BlogHitokotoMgr {
	if db == nil {
		panic(fmt.Errorf("BlogHitokotoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BlogHitokotoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("blog_hitokoto"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BlogHitokotoMgr) GetTableName() string {
	return "blog_hitokoto"
}

// Reset 重置gorm会话
func (obj *_BlogHitokotoMgr) Reset() *_BlogHitokotoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BlogHitokotoMgr) Get() (result BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BlogHitokotoMgr) Gets() (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_BlogHitokotoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BlogHitokotoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUUID uuid获取
func (obj *_BlogHitokotoMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithHitokoto hitokoto获取
func (obj *_BlogHitokotoMgr) WithHitokoto(hitokoto string) Option {
	return optionFunc(func(o *options) { o.query["hitokoto"] = hitokoto })
}

// WithType type获取
func (obj *_BlogHitokotoMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithFrom from获取
func (obj *_BlogHitokotoMgr) WithFrom(from string) Option {
	return optionFunc(func(o *options) { o.query["from"] = from })
}

// WithFromWho from_who获取
func (obj *_BlogHitokotoMgr) WithFromWho(fromWho string) Option {
	return optionFunc(func(o *options) { o.query["from_who"] = fromWho })
}

// WithCreator creator获取
func (obj *_BlogHitokotoMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithCreatorUID creator_uid获取
func (obj *_BlogHitokotoMgr) WithCreatorUID(creatorUID uint) Option {
	return optionFunc(func(o *options) { o.query["creator_uid"] = creatorUID })
}

// WithReviewer reviewer获取
func (obj *_BlogHitokotoMgr) WithReviewer(reviewer uint) Option {
	return optionFunc(func(o *options) { o.query["reviewer"] = reviewer })
}

// WithCommitFrom commit_from获取
func (obj *_BlogHitokotoMgr) WithCommitFrom(commitFrom string) Option {
	return optionFunc(func(o *options) { o.query["commit_from"] = commitFrom })
}

// WithCreatedAt created_at获取
func (obj *_BlogHitokotoMgr) WithCreatedAt(createdAt string) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithLength length获取
func (obj *_BlogHitokotoMgr) WithLength(length uint) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// GetByOption 功能选项模式获取
func (obj *_BlogHitokotoMgr) GetByOption(opts ...Option) (result BlogHitokoto, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BlogHitokotoMgr) GetByOptions(opts ...Option) (results []*BlogHitokoto, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_BlogHitokotoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]BlogHitokoto, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where(options.query)
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
func (obj *_BlogHitokotoMgr) GetFromID(id uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromID(ids []uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容
func (obj *_BlogHitokotoMgr) GetFromUUID(uuid string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromUUID(uuids []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`uuid` IN (?)", uuids).Find(&results).Error

	return
}

// GetFromHitokoto 通过hitokoto获取内容
func (obj *_BlogHitokotoMgr) GetFromHitokoto(hitokoto string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`hitokoto` = ?", hitokoto).Find(&results).Error

	return
}

// GetBatchFromHitokoto 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromHitokoto(hitokotos []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`hitokoto` IN (?)", hitokotos).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_BlogHitokotoMgr) GetFromType(_type string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromType(_types []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromFrom 通过from获取内容
func (obj *_BlogHitokotoMgr) GetFromFrom(from string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`from` = ?", from).Find(&results).Error

	return
}

// GetBatchFromFrom 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromFrom(froms []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`from` IN (?)", froms).Find(&results).Error

	return
}

// GetFromFromWho 通过from_who获取内容
func (obj *_BlogHitokotoMgr) GetFromFromWho(fromWho string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`from_who` = ?", fromWho).Find(&results).Error

	return
}

// GetBatchFromFromWho 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromFromWho(fromWhos []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`from_who` IN (?)", fromWhos).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容
func (obj *_BlogHitokotoMgr) GetFromCreator(creator string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromCreator(creators []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromCreatorUID 通过creator_uid获取内容
func (obj *_BlogHitokotoMgr) GetFromCreatorUID(creatorUID uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`creator_uid` = ?", creatorUID).Find(&results).Error

	return
}

// GetBatchFromCreatorUID 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromCreatorUID(creatorUIDs []uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`creator_uid` IN (?)", creatorUIDs).Find(&results).Error

	return
}

// GetFromReviewer 通过reviewer获取内容
func (obj *_BlogHitokotoMgr) GetFromReviewer(reviewer uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`reviewer` = ?", reviewer).Find(&results).Error

	return
}

// GetBatchFromReviewer 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromReviewer(reviewers []uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`reviewer` IN (?)", reviewers).Find(&results).Error

	return
}

// GetFromCommitFrom 通过commit_from获取内容
func (obj *_BlogHitokotoMgr) GetFromCommitFrom(commitFrom string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`commit_from` = ?", commitFrom).Find(&results).Error

	return
}

// GetBatchFromCommitFrom 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromCommitFrom(commitFroms []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`commit_from` IN (?)", commitFroms).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_BlogHitokotoMgr) GetFromCreatedAt(createdAt string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromCreatedAt(createdAts []string) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_BlogHitokotoMgr) GetFromLength(length uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_BlogHitokotoMgr) GetBatchFromLength(lengths []uint) (results []*BlogHitokoto, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogHitokoto{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
