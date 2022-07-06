package app

import (
	"github.com/gin-gonic/gin"
	"github.com/itachilee/gin/pkg/convert"
)

const (
	MaxPageSize     = 100
	DefaultPageSize = 10
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return 10
	}
	if pageSize > MaxPageSize {
		return MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
