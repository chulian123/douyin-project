package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	"log"
	"test.com/project-common/encrypts"
	"test.com/project-common/errs"
	"test.com/project-common/jwts"
	"test.com/project-grpc/user/login"
	"test.com/project-user/internal/dao"
	"test.com/project-user/internal/data/member"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/tran"
	"test.com/project-user/internal/repo"
	"test.com/project-user/pkg/model"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache       repo.Cache
	memberRepo  repo.MemberRepo
	transaction tran.Transaction
}

func New() *LoginService {
	return &LoginService{
		cache:       dao.Rc,
		memberRepo:  dao.NewMemberDao(),
		transaction: dao.NewTransaction(),
	}
}

func (ls *LoginService) Login(ctx context.Context, msg *login.DouyinUserLoginRequest) (*login.DouyinUserLoginResponse, error) {
	c := context.Background()
	//1.去数据库查询 账号密码是否正确
	//pwd := encrypts.Md5(msg.Password)
	mem, err := ls.memberRepo.FindMember(c, msg.Username, msg.Password)
	if err != nil {
		return &login.DouyinUserLoginResponse{
			StatusCode: &errs.LoginFaildCode,
			StatusMsg:  &errs.DBError,
		}, err
	}
	if mem == nil {
		return &login.DouyinUserLoginResponse{
			StatusCode: &errs.LoginFaildCode,
			StatusMsg:  &errs.UserERROR,
		}, err
	}
	log.Println("用户id", mem.Id)
	// 2. 用jwt生成token
	id := mem.Id            // 获取指针指向的值
	int64Value := int64(id) // 将值转换为 int64 类型
	token := jwts.CreateToken(int64Value)
	if token == "" {
		return &login.DouyinUserLoginResponse{
			StatusCode: &errs.LoginFaildCode,
			StatusMsg:  &errs.TokenError,
		}, err
	}
	log.Printf("查询结果: %+v\n", mem)
	return &login.DouyinUserLoginResponse{
		StatusCode: &errs.LoginSuccessCode,
		StatusMsg:  &errs.LoginSuccessMsg,
		UserId:     &mem.Id,
		Token:      &token,
	}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *login.DouyinUserRegisterRequest) (*login.DouyinUserRegisterResponse, error) {
	// 1.接受参数
	c := context.Background()
	// 检查用户有没有被注册
	exist, err := ls.memberRepo.GetMemberByEmail(c, msg.Username)
	if err != nil {
		return &login.DouyinUserRegisterResponse{
			StatusCode: &errs.LoginFaildCode,
			StatusMsg:  &errs.DBError,
			UserId:     &errs.BasicUserCode,
		}, err
	}
	if exist {
		return &login.DouyinUserRegisterResponse{
			StatusCode: &errs.LoginFaildCode,
			StatusMsg:  &errs.UserExist,
		}, nil
	}
	userID := encrypts.GetUserID()
	log.Println("userID:", userID)

	mem := &member.Member{
		Id:   userID,
		Name: *msg.Username,
		Pwd:  *msg.Password,
	}
	log.Println("SaveMember开始")
	err = ls.transaction.Action(func(conn database.DbConn) error {
		err = ls.memberRepo.SaveMember(conn, c, *mem)
		if err != nil {
			zap.L().Error("Register db SaveMember error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		return nil
	})
	log.Println("用户id", mem.Id)
	// 2. 用jwt生成token
	id := mem.Id     // 获取指针指向的值
	int64Value := id // 将值转换为 int64 类型
	token := jwts.CreateToken(int64Value)
	if token == "" {
		return &login.DouyinUserRegisterResponse{
			StatusCode: &errs.LoginFaildCode,
			StatusMsg:  &errs.TokenError,
		}, nil
	}
	return &login.DouyinUserRegisterResponse{
		StatusCode: &errs.LoginSuccessCode,
		StatusMsg:  &errs.LoginSuccessMsg,
		UserId:     &mem.Id,
		Token:      &token,
	}, nil
}

//func (ls *LoginService) Feed(ctx context.Context, msg *login.DouyinFeedRequest) (*login.DouyinFeedResponse, error) {
//	c := context.Background()
//	//查询数据
//
//}
