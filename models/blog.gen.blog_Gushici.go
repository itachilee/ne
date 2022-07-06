package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _BlogGushiciMgr struct {
	*_BaseMgr
}

// BlogGushiciMgr open func
func BlogGushiciMgr(db *gorm.DB) *_BlogGushiciMgr {
	if db == nil {
		panic(fmt.Errorf("BlogGushiciMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BlogGushiciMgr{_BaseMgr: &_BaseMgr{DB: db.Table("blog_Gushici"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BlogGushiciMgr) GetTableName() string {
	return "blog_Gushici"
}

// Reset 重置gorm会话
func (obj *_BlogGushiciMgr) Reset() *_BlogGushiciMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BlogGushiciMgr) Get() (result BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BlogGushiciMgr) Gets() (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_BlogGushiciMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_BlogGushiciMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithCreatedAt CreatedAt获取
func (obj *_BlogGushiciMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreatedAt"] = createdAt })
}

// WithUpdatedAt UpdatedAt获取
func (obj *_BlogGushiciMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["UpdatedAt"] = updatedAt })
}

// WithDeletedAt DeletedAt获取
func (obj *_BlogGushiciMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["DeletedAt"] = deletedAt })
}

// WithAuthor Author获取 '作者'
func (obj *_BlogGushiciMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["Author"] = author })
}

// WithParagraphs Paragraphs获取 '正文'
func (obj *_BlogGushiciMgr) WithParagraphs(paragraphs string) Option {
	return optionFunc(func(o *options) { o.query["Paragraphs"] = paragraphs })
}

// WithNote Note获取 '备注'
func (obj *_BlogGushiciMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["Note"] = note })
}

// WithTitle Title获取 '标题'
func (obj *_BlogGushiciMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["Title"] = title })
}

// WithCategory Category获取 '分类'
func (obj *_BlogGushiciMgr) WithCategory(category string) Option {
	return optionFunc(func(o *options) { o.query["Category"] = category })
}

// WithFromFile FromFile获取 '来源文件'
func (obj *_BlogGushiciMgr) WithFromFile(fromFile string) Option {
	return optionFunc(func(o *options) { o.query["FromFile"] = fromFile })
}

// WithContent Content获取 '正文'
func (obj *_BlogGushiciMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["Content"] = content })
}

// WithOrigin Origin获取 '起源'
func (obj *_BlogGushiciMgr) WithOrigin(origin string) Option {
	return optionFunc(func(o *options) { o.query["Origin"] = origin })
}

// GetByOption 功能选项模式获取
func (obj *_BlogGushiciMgr) GetByOption(opts ...Option) (result BlogGushici, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BlogGushiciMgr) GetByOptions(opts ...Option) (results []*BlogGushici, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_BlogGushiciMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]BlogGushici, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where(options.query)
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

// GetFromID 通过ID获取内容
func (obj *_BlogGushiciMgr) GetFromID(id uint64) (result BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`ID` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BlogGushiciMgr) GetBatchFromID(ids []uint64) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过CreatedAt获取内容
func (obj *_BlogGushiciMgr) GetFromCreatedAt(createdAt time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`CreatedAt` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_BlogGushiciMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`CreatedAt` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过UpdatedAt获取内容
func (obj *_BlogGushiciMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`UpdatedAt` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_BlogGushiciMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`UpdatedAt` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过DeletedAt获取内容
func (obj *_BlogGushiciMgr) GetFromDeletedAt(deletedAt time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`DeletedAt` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找
func (obj *_BlogGushiciMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`DeletedAt` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromAuthor 通过Author获取内容 '作者'
func (obj *_BlogGushiciMgr) GetFromAuthor(author string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Author` = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量查找 '作者'
func (obj *_BlogGushiciMgr) GetBatchFromAuthor(authors []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Author` IN (?)", authors).Find(&results).Error

	return
}

// GetFromParagraphs 通过Paragraphs获取内容 '正文'
func (obj *_BlogGushiciMgr) GetFromParagraphs(paragraphs string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Paragraphs` = ?", paragraphs).Find(&results).Error

	return
}

// GetBatchFromParagraphs 批量查找 '正文'
func (obj *_BlogGushiciMgr) GetBatchFromParagraphs(paragraphss []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Paragraphs` IN (?)", paragraphss).Find(&results).Error

	return
}

// GetFromNote 通过Note获取内容 '备注'
func (obj *_BlogGushiciMgr) GetFromNote(note string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Note` = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量查找 '备注'
func (obj *_BlogGushiciMgr) GetBatchFromNote(notes []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Note` IN (?)", notes).Find(&results).Error

	return
}

// GetFromTitle 通过Title获取内容 '标题'
func (obj *_BlogGushiciMgr) GetFromTitle(title string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 '标题'
func (obj *_BlogGushiciMgr) GetBatchFromTitle(titles []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCategory 通过Category获取内容 '分类'
func (obj *_BlogGushiciMgr) GetFromCategory(category string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Category` = ?", category).Find(&results).Error

	return
}

// GetBatchFromCategory 批量查找 '分类'
func (obj *_BlogGushiciMgr) GetBatchFromCategory(categorys []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Category` IN (?)", categorys).Find(&results).Error

	return
}

// GetFromFromFile 通过FromFile获取内容 '来源文件'
func (obj *_BlogGushiciMgr) GetFromFromFile(fromFile string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`FromFile` = ?", fromFile).Find(&results).Error

	return
}

// GetBatchFromFromFile 批量查找 '来源文件'
func (obj *_BlogGushiciMgr) GetBatchFromFromFile(fromFiles []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`FromFile` IN (?)", fromFiles).Find(&results).Error

	return
}

// GetFromContent 通过Content获取内容 '正文'
func (obj *_BlogGushiciMgr) GetFromContent(content string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 '正文'
func (obj *_BlogGushiciMgr) GetBatchFromContent(contents []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromOrigin 通过Origin获取内容 '起源'
func (obj *_BlogGushiciMgr) GetFromOrigin(origin string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Origin` = ?", origin).Find(&results).Error

	return
}

// GetBatchFromOrigin 批量查找 '起源'
func (obj *_BlogGushiciMgr) GetBatchFromOrigin(origins []string) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`Origin` IN (?)", origins).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BlogGushiciMgr) FetchByPrimaryKey(id uint64) (result BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`ID` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxBlogGushiciDeletedAt  获取多个内容
func (obj *_BlogGushiciMgr) FetchIndexByIDxBlogGushiciDeletedAt(deletedAt time.Time) (results []*BlogGushici, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BlogGushici{}).Where("`DeletedAt` = ?", deletedAt).Find(&results).Error

	return
}
