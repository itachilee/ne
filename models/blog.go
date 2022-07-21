package models

import (
	"time"
)

// BlogGushici [...]
type BlogGushici struct {
	ID         uint64    `gorm:"primaryKey;column:ID" json:"-"`
	CreatedAt  time.Time `gorm:"column:CreatedAt" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:UpdatedAt" json:"updatedAt"`
	DeletedAt  time.Time `gorm:"column:DeletedAt" json:"deletedAt"`
	Author     string    `gorm:"column:Author" json:"author"`         // '作者'
	Paragraphs string    `gorm:"column:Paragraphs" json:"paragraphs"` // '正文'
	Note       string    `gorm:"column:Note" json:"note"`             // '备注'
	Title      string    `gorm:"column:Title" json:"title"`           // '标题'
	Category   string    `gorm:"column:Category" json:"category"`     // '分类'
	FromFile   string    `gorm:"column:FromFile" json:"fromFile"`     // '来源文件'
	Content    string    `gorm:"column:Content" json:"content"`       // '正文'
	Origin     string    `gorm:"column:Origin" json:"origin"`         // '起源'
}

// TableName get sql table name.获取数据库表名
func (m *BlogGushici) TableName() string {
	return "blog_Gushici"
}

// BlogGushiciColumns get sql column name.获取数据库列名
var BlogGushiciColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	Author     string
	Paragraphs string
	Note       string
	Title      string
	Category   string
	FromFile   string
	Content    string
	Origin     string
}{
	ID:         "ID",
	CreatedAt:  "CreatedAt",
	UpdatedAt:  "UpdatedAt",
	DeletedAt:  "DeletedAt",
	Author:     "Author",
	Paragraphs: "Paragraphs",
	Note:       "Note",
	Title:      "Title",
	Category:   "Category",
	FromFile:   "FromFile",
	Content:    "Content",
	Origin:     "Origin",
}

// BlogHitokoto [...]
type BlogHitokoto struct {
	ID         uint   `gorm:"column:id" json:"id"`
	UUID       string `gorm:"column:uuid" json:"uuid"`
	Hitokoto   string `gorm:"column:hitokoto" json:"hitokoto"`
	Type       string `gorm:"column:type" json:"type"`
	From       string `gorm:"column:from" json:"from"`
	FromWho    string `gorm:"column:from_who" json:"fromWho"`
	Creator    string `gorm:"column:creator" json:"creator"`
	CreatorUID uint   `gorm:"column:creator_uid" json:"creatorUid"`
	Reviewer   uint   `gorm:"column:reviewer" json:"reviewer"`
	CommitFrom string `gorm:"column:commit_from" json:"commitFrom"`
	CreatedAt  string `gorm:"column:created_at" json:"createdAt"`
	Length     uint   `gorm:"column:length" json:"length"`
}

// TableName get sql table name.获取数据库表名
func (m *BlogHitokoto) TableName() string {
	return "blog_hitokoto"
}

// BlogHitokotoColumns get sql column name.获取数据库列名
var BlogHitokotoColumns = struct {
	ID         string
	UUID       string
	Hitokoto   string
	Type       string
	From       string
	FromWho    string
	Creator    string
	CreatorUID string
	Reviewer   string
	CommitFrom string
	CreatedAt  string
	Length     string
}{
	ID:         "id",
	UUID:       "uuid",
	Hitokoto:   "hitokoto",
	Type:       "type",
	From:       "from",
	FromWho:    "from_who",
	Creator:    "creator",
	CreatorUID: "creator_uid",
	Reviewer:   "reviewer",
	CommitFrom: "commit_from",
	CreatedAt:  "created_at",
	Length:     "length",
}

// StockA [...]
type StockA struct {
	CompanyCode int    `gorm:"column:company_code" json:"companyCode"`
	CompantName string `gorm:"column:compant_name" json:"compantName"`
	Code        int    `gorm:"column:code" json:"code"`
	Name        string `gorm:"column:name" json:"name"`
	Date        string `gorm:"column:date" json:"date"`
}

// TableName get sql table name.获取数据库表名
func (m *StockA) TableName() string {
	return "stock_a"
}

// StockAColumns get sql column name.获取数据库列名
var StockAColumns = struct {
	CompanyCode string
	CompantName string
	Code        string
	Name        string
	Date        string
}{
	CompanyCode: "company_code",
	CompantName: "compant_name",
	Code:        "code",
	Name:        "name",
	Date:        "date",
}

// StockB [...]
type StockB struct {
	CompanyCode int    `gorm:"column:company_code" json:"companyCode"`
	CompantName string `gorm:"column:compant_name" json:"compantName"`
	Code        int    `gorm:"column:code" json:"code"`
	Name        string `gorm:"column:name" json:"name"`
	Date        string `gorm:"column:date" json:"date"`
}

// TableName get sql table name.获取数据库表名
func (m *StockB) TableName() string {
	return "stock_b"
}

// StockBColumns get sql column name.获取数据库列名
var StockBColumns = struct {
	CompanyCode string
	CompantName string
	Code        string
	Name        string
	Date        string
}{
	CompanyCode: "company_code",
	CompantName: "compant_name",
	Code:        "code",
	Name:        "name",
	Date:        "date",
}

// ZCategories [...]
type ZCategories struct {
	ID          uint      `gorm:"primaryKey;column:id" json:"-"`
	Name        string    `gorm:"column:name" json:"name"`                // 分类名
	DisplayName string    `gorm:"column:display_name" json:"displayName"` // 分类别名
	SeoDesc     string    `gorm:"column:seo_desc" json:"seoDesc"`         // seo描述
	ParentID    int       `gorm:"column:parent_id" json:"parentId"`       // 父类ID
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZCategories) TableName() string {
	return "z_categories"
}

// ZCategoriesColumns get sql column name.获取数据库列名
var ZCategoriesColumns = struct {
	ID          string
	Name        string
	DisplayName string
	SeoDesc     string
	ParentID    string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	DisplayName: "display_name",
	SeoDesc:     "seo_desc",
	ParentID:    "parent_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// ZCreditCard [...]
type ZCreditCard struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	Number    string    `gorm:"column:number" json:"number"`
	Userid    int       `gorm:"column:userid" json:"userid"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZCreditCard) TableName() string {
	return "z_credit_card"
}

// ZCreditCardColumns get sql column name.获取数据库列名
var ZCreditCardColumns = struct {
	ID        string
	Number    string
	Userid    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Number:    "number",
	Userid:    "userid",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// ZLinks [...]
type ZLinks struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`   // 友链名
	Link      string    `gorm:"column:link" json:"link"`   // 友链链接
	Order     int       `gorm:"column:order" json:"order"` // 排序
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZLinks) TableName() string {
	return "z_links"
}

// ZLinksColumns get sql column name.获取数据库列名
var ZLinksColumns = struct {
	ID        string
	Name      string
	Link      string
	Order     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Link:      "link",
	Order:     "order",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// ZPasswordResets [...]
type ZPasswordResets struct {
	Email     string    `gorm:"column:email" json:"email"`
	Token     string    `gorm:"column:token" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZPasswordResets) TableName() string {
	return "z_password_resets"
}

// ZPasswordResetsColumns get sql column name.获取数据库列名
var ZPasswordResetsColumns = struct {
	Email     string
	Token     string
	CreatedAt string
}{
	Email:     "email",
	Token:     "token",
	CreatedAt: "created_at",
}

// ZPostCate [...]
type ZPostCate struct {
	ID     uint `gorm:"primaryKey;column:id" json:"-"`
	PostID int  `gorm:"column:post_id" json:"postId"` // 文章ID
	CateID int  `gorm:"column:cate_id" json:"cateId"` // 分类ID
}

// TableName get sql table name.获取数据库表名
func (m *ZPostCate) TableName() string {
	return "z_post_cate"
}

// ZPostCateColumns get sql column name.获取数据库列名
var ZPostCateColumns = struct {
	ID     string
	PostID string
	CateID string
}{
	ID:     "id",
	PostID: "post_id",
	CateID: "cate_id",
}

// ZPostTag [...]
type ZPostTag struct {
	ID     uint `gorm:"primaryKey;column:id" json:"-"`
	PostID int  `gorm:"column:post_id" json:"postId"` // 文章ID
	TagID  int  `gorm:"column:tag_id" json:"tagId"`   // 标签ID
}

// TableName get sql table name.获取数据库表名
func (m *ZPostTag) TableName() string {
	return "z_post_tag"
}

// ZPostTagColumns get sql column name.获取数据库列名
var ZPostTagColumns = struct {
	ID     string
	PostID string
	TagID  string
}{
	ID:     "id",
	PostID: "post_id",
	TagID:  "tag_id",
}

// ZPostViews [...]
type ZPostViews struct {
	ID     uint `gorm:"primaryKey;column:id" json:"-"`
	PostID int  `gorm:"column:post_id" json:"postId"` // 文章ID
	Num    int  `gorm:"column:num" json:"num"`        // 阅读次数
}

// TableName get sql table name.获取数据库表名
func (m *ZPostViews) TableName() string {
	return "z_post_views"
}

// ZPostViewsColumns get sql column name.获取数据库列名
var ZPostViewsColumns = struct {
	ID     string
	PostID string
	Num    string
}{
	ID:     "id",
	PostID: "post_id",
	Num:    "num",
}

// ZPosts [...]
type ZPosts struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	UID       string    `gorm:"column:uid" json:"uid"`           // uid
	UserID    int       `gorm:"column:user_id" json:"userId"`    // 用户ID
	Title     string    `gorm:"column:title" json:"title"`       // 标题
	Summary   string    `gorm:"column:summary" json:"summary"`   // 摘要
	Original  string    `gorm:"column:original" json:"original"` // 原文章内容
	Content   string    `gorm:"column:content" json:"content"`   // 文章内容
	Password  string    `gorm:"column:password" json:"password"` // 文章密码
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZPosts) TableName() string {
	return "z_posts"
}

// ZPostsColumns get sql column name.获取数据库列名
var ZPostsColumns = struct {
	ID        string
	UID       string
	UserID    string
	Title     string
	Summary   string
	Original  string
	Content   string
	Password  string
	DeletedAt string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UID:       "uid",
	UserID:    "user_id",
	Title:     "title",
	Summary:   "summary",
	Original:  "original",
	Content:   "content",
	Password:  "password",
	DeletedAt: "deleted_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// ZSystems [...]
type ZSystems struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"-"`
	Theme        int8      `gorm:"column:theme" json:"theme"`                // 主题
	Title        string    `gorm:"column:title" json:"title"`                // 网站title
	Keywords     string    `gorm:"column:keywords" json:"keywords"`          // 网站关键字
	Description  string    `gorm:"column:description" json:"description"`    // 网站描述
	RecordNumber string    `gorm:"column:record_number" json:"recordNumber"` // 备案号
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZSystems) TableName() string {
	return "z_systems"
}

// ZSystemsColumns get sql column name.获取数据库列名
var ZSystemsColumns = struct {
	ID           string
	Theme        string
	Title        string
	Keywords     string
	Description  string
	RecordNumber string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	Theme:        "theme",
	Title:        "title",
	Keywords:     "keywords",
	Description:  "description",
	RecordNumber: "record_number",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// ZTags [...]
type ZTags struct {
	ID          uint      `gorm:"primaryKey;column:id" json:"-"`
	Name        string    `gorm:"column:name" json:"name"`                // 标签名
	DisplayName string    `gorm:"column:display_name" json:"displayName"` // 标签别名
	SeoDesc     string    `gorm:"column:seo_desc" json:"seoDesc"`         // seo描述
	Num         int       `gorm:"column:num" json:"num"`                  // 被使用次数
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ZTags) TableName() string {
	return "z_tags"
}

// ZTagsColumns get sql column name.获取数据库列名
var ZTagsColumns = struct {
	ID          string
	Name        string
	DisplayName string
	SeoDesc     string
	Num         string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	DisplayName: "display_name",
	SeoDesc:     "seo_desc",
	Num:         "num",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// ZUsers [...]
type ZUsers struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"-"`
	Name            string    `gorm:"column:name" json:"name"`     // 用户名
	Email           string    `gorm:"column:email" json:"email"`   // 邮箱
	Status          int8      `gorm:"column:status" json:"status"` // 用户状态 0创建,1正常
	EmailVerifiedAt int64     `gorm:"column:email_verified_at" json:"emailVerifiedAt"`
	Password        string    `gorm:"column:password" json:"password"` // 密码
	RememberToken   string    `gorm:"column:remember_token" json:"rememberToken"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updatedAt"`
	Role            string    `json:"role" gorm:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *ZUsers) TableName() string {
	return "z_users"
}

// ZUsersColumns get sql column name.获取数据库列名
var ZUsersColumns = struct {
	ID              string
	Name            string
	Email           string
	Status          string
	EmailVerifiedAt string
	Password        string
	RememberToken   string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	Name:            "name",
	Email:           "email",
	Status:          "status",
	EmailVerifiedAt: "email_verified_at",
	Password:        "password",
	RememberToken:   "remember_token",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}
