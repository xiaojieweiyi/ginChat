package service

import (
	"ginchat/models"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "view/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	/*c.JSON(200, gin.H{
		"message": "welcome",
	}) */
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("view/user/register.html")
	if err != nil {
		panic(err)
	}
	err = ind.Execute(c.Writer, "register")
	if err != nil {
		panic(err)
	}
	/*c.JSON(200, gin.H{
		"message": "welcome",
	}) */
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("view/chat/index.html",
		"view/chat/head.html",
		"view/chat/foot.html",
		"view/chat/tabmenu.html",
		"view/chat/concat.html",
		"view/chat/group.html",
		"view/chat/profile.html",
		"view/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	ind.Execute(c.Writer, "index")
}

func Chat(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
