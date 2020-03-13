package handler

import (

	"context"
	"go.uber.org/zap"
	"hongbao/movies/src/share/errors"
	"hongbao/movies/src/share/pb"
	"hongbao/movies/src/share/utils/log"
	"hongbao/movies/src/user-srv/db"
)
type UserServiceExtHandler struct {
	logger *zap.Logger
}
//初始化
func NewUserServiceExtHandler() *UserServiceExtHandler {
	return &UserServiceExtHandler{
		logger: log.Instance(),
	}
}


//用户注册
func (u *UserServiceExtHandler) RegistAccount(ctx context.Context, req *pb.RegistAccountReq,rep *pb.RegistAccountRsp) error {
	userName := req.UserName
	password := req.Password
	email := req.Email
	user,err := db.SelectUserByEmail(email)

	if err != nil{
		u.logger.Error("error",zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user != nil {
		return errors.ErrorUserAlready
	}
	err = db.InsertUser(userName, password, email)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	return errors.ErrorUserSuccess
}
//用户登录
func (u *UserServiceExtHandler) LoginAccount(ctx context.Context, req *pb.LoginAccountReq, rsp *pb.LoginAccountRsp) error {
	email := req.Email
	password := req.Password
	user, err := db.SelectUserByPasswordName(email, password)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user == nil {
		return errors.ErrorUserLoginFailed
	}
	rsp.Email = user.Email.String
	rsp.Phone = user.Phone.String
	rsp.UserID = user.Id
	rsp.UserName = user.Name
	return nil
}
// 密码重置
func (u *UserServiceExtHandler) ResetAccount(ctx context.Context, req *pb.ResetAccountReq, rsp *pb.ResetAccountRsp) error {
	return nil
}

//评分
func (u *UserServiceExtHandler) WantScore(ctx context.Context, req *pb.WantScoreReq, rsp *pb.WantScoreRsp) error {

	orderNum, err := db.SelectOrderByUidMid(req.MovieId, req.UserId)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	if orderNum == 0 {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorScoreForbid
	}
	err = db.UpdateOrderScore(req.MovieId, req.Score)
	if err != nil {
		u.logger.Error("error", zap.Error(err))
		return errors.ErrorUserFailed
	}
	return nil
}
// 修改用户信息
func (u *UserServiceExtHandler) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileReq, rsp *pb.UpdateUserProfileRsp) error {

	userEmail := req.UserEmail
	userName := req.UserName
	userPhone := req.UserPhone
	userID := req.UserID
	if userEmail != "" {
		err := db.UpdateUserEmailProfile(userEmail, userID)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userName != "" {
		err := db.UpdateUserNameProfile(userName, userID)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userPhone != "" {
		err := db.UpdateUserPhoneProfile(userPhone, userID)
		if err != nil {
			u.logger.Error("error", zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	return nil
}

