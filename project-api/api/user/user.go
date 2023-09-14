package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
	"time"
)

var LoginFaildCode int32 = 0
var DBERORR = "数据连接错误"

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (u *HandlerUser) register(c *gin.Context) {
	//1.接收参数 参数模型
	//result := &common.Result{}
	username := c.Query("username")
	password := c.Query("password")
	//2。调用grpc进行登录
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msg := &login.DouyinUserRegisterRequest{
		Password: &password,
		Username: &username,
	}
	registerRsp, err := LoginServiceClient.Register(ctx, msg)
	if err != nil {
		c.JSON(http.StatusOK, &login.DouyinUserLoginResponse{
			StatusCode: &LoginFaildCode,
			StatusMsg:  &errs.UserERROR,
		})
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, registerRsp)
}

func (u *HandlerUser) login(c *gin.Context) {
	//1.接收参数 参数模型
	//result := &common.Result{}
	username := c.Query("username")
	password := c.Query("password")
	//2。调用grpc进行登录
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msg := &login.DouyinUserLoginRequest{
		Password: &password,
		Username: &username,
	}
	loginRsp, err := LoginServiceClient.Login(ctx, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, &login.DouyinUserLoginResponse{
			StatusCode: &LoginFaildCode,
			StatusMsg:  &errs.UserERROR,
		})
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, loginRsp)
}

func (u *HandlerUser) feed(c *gin.Context) {
	//1.接收参数 参数模型
	//result := &common.Result{}
	lastTimeInt64, _ := strconv.ParseInt(c.Query("last_time"), 10, 64)
	token := c.Query("token")
	//2。调用grpc进行登录
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feed := &login.DouyinFeedRequest{
		LatestTime: &lastTimeInt64,
		Token:      &token,
	}
	FeedRsp, err := LoginServiceClient.Feed(ctx, feed)
	if err != nil {
		c.JSON(http.StatusBadRequest, &login.DouyinFeedResponse{
			StatusCode: &LoginFaildCode,
			StatusMsg:  &errs.UserERROR,
		})
		return
	}
	c.JSON(http.StatusOK, FeedRsp)
}
