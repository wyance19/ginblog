package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询所有文章列表

//查询单个文章信息

//查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArt(pageSize, pageNum)
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑用户
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBind(&data)
	code = model.EditArt(id, &data)
	c.JSON(200, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)
	c.JSON(200, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
