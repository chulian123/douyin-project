package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/router"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	//初始化grpc的客户端连接
	InitRpcUserClient()
	h := New()
	r.POST("/douyin/user/register/", h.register)
	r.POST("/douyin/user/login/", h.login)
	r.GET("/douyin/feed/", h.feed)

}
