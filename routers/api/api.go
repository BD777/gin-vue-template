package api

import (
	"gin-vue-template/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinContextHelper struct {
	c *gin.Context
}

func NewGinContextHelper(c *gin.Context) *GinContextHelper {
	return &GinContextHelper{c: c}
}

func (h *GinContextHelper) GetQueryInt(key string, defaultValue int) int {
	v := h.c.DefaultQuery(key, strconv.FormatInt(int64(defaultValue), 10))
	val, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("[GinContextHelper#GetQueryInt] parse key[%v] with value[%v] error[%v]", key, v, err)
		return defaultValue
	}
	return val
}

func (h *GinContextHelper) GetPagination() *models.Pagination {
	pagination := &models.Pagination{
		Page:     h.GetQueryInt("page", 1),
		PageSize: h.GetQueryInt("pagesize", 20),
	}
	return pagination
}
