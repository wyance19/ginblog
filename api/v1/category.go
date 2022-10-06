package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加用户
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑用户
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBind(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(200, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCate(id)
	c.JSON(200, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
