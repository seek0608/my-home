package route

import (
	"github.com/gin-gonic/gin"
	"my-home/logic/rbac"
	"my-home/utils"
	"net/http"
	"strconv"
)

var R *gin.Engine

var WHITE = [2]string{"", ""}

func init() {
	R = gin.Default()
	userGroup := R.Group("/user")
	{
		userGroup.POST("/login", login)
		userGroup.POST("/check", checkLogin)

	}
}

func checkLogin(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	customClaims, err := utils.ParseToken(authorization)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "login error",
		})
		return
	}
	c.JSON(200, gin.H{"data": WithOptions(WithData(customClaims.UserInfo))})
}

type Login struct {
	Code      string `json:"code"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}

func login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user Login
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	openId, err := rbac.GetOpenId(user.Code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "获取openid失败",
		})
		return
	}
	info := rbac.UserInfo{
		WxId:      openId,
		NickName:  user.NickName,
		AvatarUrl: user.AvatarUrl,
	}
	println(info.WxId)
	// 校验用户名和密码是否正确
	if info.WxId == "o6Vsm5QM8F7JSnCtGMTe_XO_KDO4" || info.WxId == "123" {
		// 生成Token
		tokenString, _ := utils.GenToken(info)
		c.JSON(http.StatusOK, gin.H{
			"msg": gin.H{
				"token": tokenString,
			},
			"success": true,
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

func Init(port int) {
	R.Run(":" + strconv.Itoa(port))
}
