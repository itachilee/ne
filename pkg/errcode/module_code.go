package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取会员列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建会员失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新会员失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除会员失败")
	ErrorCountTagFail   = NewError(20010005, "统计会员失败")
)
