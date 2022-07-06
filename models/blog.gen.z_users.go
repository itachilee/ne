package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZUsersMgr struct {
	*_BaseMgr
}

// ZUsersMgr open func
func ZUsersMgr(db *gorm.DB) *_ZUsersMgr {
	if db == nil {
		panic(fmt.Errorf("ZUsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZUsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("z_users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZUsersMgr) GetTableName() string {
	return "z_users"
}

// Reset 重置gorm会话
func (obj *_ZUsersMgr) Reset() *_ZUsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZUsersMgr) Get() (result ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZUsersMgr) Gets() (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZUsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZUsersMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 用户名
func (obj *_ZUsersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEmail email获取 邮箱
func (obj *_ZUsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithStatus status获取 用户状态 0创建,1正常
func (obj *_ZUsersMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithEmailVerifiedAt email_verified_at获取
func (obj *_ZUsersMgr) WithEmailVerifiedAt(emailVerifiedAt int64) Option {
	return optionFunc(func(o *options) { o.query["email_verified_at"] = emailVerifiedAt })
}

// WithPassword password获取 密码
func (obj *_ZUsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithRememberToken remember_token获取
func (obj *_ZUsersMgr) WithRememberToken(rememberToken string) Option {
	return optionFunc(func(o *options) { o.query["remember_token"] = rememberToken })
}

// WithCreatedAt created_at获取
func (obj *_ZUsersMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZUsersMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZUsersMgr) GetByOption(opts ...Option) (result ZUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZUsersMgr) GetByOptions(opts ...Option) (results []*ZUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZUsersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZUsers, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where(options.query)
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
func (obj *_ZUsersMgr) GetFromID(id uint) (result ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZUsersMgr) GetBatchFromID(ids []uint) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 用户名
func (obj *_ZUsersMgr) GetFromName(name string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 用户名
func (obj *_ZUsersMgr) GetBatchFromName(names []string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_ZUsersMgr) GetFromEmail(email string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *_ZUsersMgr) GetBatchFromEmail(emails []string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 用户状态 0创建,1正常
func (obj *_ZUsersMgr) GetFromStatus(status int8) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 用户状态 0创建,1正常
func (obj *_ZUsersMgr) GetBatchFromStatus(statuss []int8) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromEmailVerifiedAt 通过email_verified_at获取内容
func (obj *_ZUsersMgr) GetFromEmailVerifiedAt(emailVerifiedAt int64) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`email_verified_at` = ?", emailVerifiedAt).Find(&results).Error

	return
}

// GetBatchFromEmailVerifiedAt 批量查找
func (obj *_ZUsersMgr) GetBatchFromEmailVerifiedAt(emailVerifiedAts []int64) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`email_verified_at` IN (?)", emailVerifiedAts).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_ZUsersMgr) GetFromPassword(password string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_ZUsersMgr) GetBatchFromPassword(passwords []string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromRememberToken 通过remember_token获取内容
func (obj *_ZUsersMgr) GetFromRememberToken(rememberToken string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`remember_token` = ?", rememberToken).Find(&results).Error

	return
}

// GetBatchFromRememberToken 批量查找
func (obj *_ZUsersMgr) GetBatchFromRememberToken(rememberTokens []string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`remember_token` IN (?)", rememberTokens).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZUsersMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZUsersMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZUsersMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZUsersMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZUsersMgr) FetchByPrimaryKey(id uint) (result ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByZUsersEmailUnique primary or index 获取唯一内容
func (obj *_ZUsersMgr) FetchUniqueByZUsersEmailUnique(email string) (result ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`email` = ?", email).First(&result).Error

	return
}

// FetchIndexByZUsersEmailIndex  获取多个内容
func (obj *_ZUsersMgr) FetchIndexByZUsersEmailIndex(email string) (results []*ZUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZUsers{}).Where("`email` = ?", email).Find(&results).Error

	return
}
