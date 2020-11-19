package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-manager/constant"
	"go-manager/model"
	"log"
	"net/http"
	"strconv"
)

func SaveUser(c *gin.Context) {
	var entity model.User
	// 请求头为application/x-www-form-urlencoded
	//name := c.PostForm("name")
	// 请求头为application/json
	err := c.BindJSON(&entity)
	if err != nil {
		log.Println(err.Error())
	}
	if entity.Id == 0 {
		constant.Db.Create(&entity)
		fmt.Println("用户", entity, "已创建")
		c.JSON(http.StatusOK, "已创建")
		return
	}
	constant.Db.Model(&entity).Updates(map[string]interface{}{"name": entity.Name, "age": entity.Age})
	fmt.Println("用户", entity, "已修改")
	c.JSON(http.StatusOK, "已修改")
}

func GetUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("id转换失败")
	}
	entity := new(model.User)
	constant.Db.First(&entity, id)
	c.JSON(http.StatusOK, entity)
}
