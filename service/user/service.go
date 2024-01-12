package user

import (
	"context"
	"errors"
	"github.com/Flash-Pass/flash-pass-server/db"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/auth"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/wechatClient"
	"go.uber.org/zap"
)

type Service struct {
	userRepo     Repository
	wechatClient WechatClient
}

type IService interface {
	Login(ctx context.Context, mobile, password string) (token string, err error)
	LoginViaWeChat(ctx context.Context, code string) (token string, err error)
	Register(ctx context.Context, mobile, password string) (token string, err error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, openId, mobile string, userId int64) (*model.User, error)
}

type Repository interface {
	Create(ctx context.Context, mobile, password string) (*model.User, error)
	CheckPassword(ctx context.Context, mobile, password string) (userId int64, ok bool)
	GetUserById(ctx context.Context, userId int64) (*model.User, error)
	GetUserByOpenId(ctx context.Context, openId string) (*model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
}

type WechatClient interface {
	LoginByCode(code string) (wechatClient.IWeChatLoginByCodeResponse, error)
}

func NewService(userRepo Repository, wechatClient WechatClient) *Service {
	return &Service{
		userRepo:     userRepo,
		wechatClient: wechatClient,
	}
}

func (s Service) Login(ctx context.Context, mobile, password string) (token string, err error) {
	ctx, persist := db.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()
	logger := ctxlog.Extract(ctx)

	userId, ok := s.userRepo.CheckPassword(ctx, mobile, password)
	if !ok {
		logger.Warn("check password defeat", zap.String("mobile", mobile), zap.String("password", password))
		return "", errors.New("check password defeat")
	}

	claim := &auth.UserClaim{
		Id: userId,
	}

	token, err = auth.GenerateToken(ctx, claim)
	if err != nil {
		logger.Warn("generate token defeat", zap.Error(err), zap.Any("claim", claim))
		return "", err
	}

	return token, nil
}

func (s Service) LoginViaWeChat(ctx context.Context, code string) (token string, err error) {
	ctx, persist := db.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()
	logger := ctxlog.Extract(ctx)

	resp, err := s.wechatClient.LoginByCode(code)
	if err != nil {
		logger.Error("login by wechat defeat", zap.Error(err), zap.String("code", code))
		return "", err
	}

	user, err := s.userRepo.GetUserByOpenId(ctx, resp.GetOpenId())
	if err != nil {
		logger.Error("get user by open id defeat", zap.Error(err), zap.String("openId", resp.GetOpenId()))
		return "", err
	}

	claim := &auth.UserClaim{
		Id:         user.Id,
		OpenId:     resp.GetOpenId(),
		SessionKey: resp.GetSessionKey(),
		UnionId:    resp.GetUnionId(),
	}
	token, err = auth.GenerateToken(ctx, claim)
	if err != nil {
		logger.Error("generate token defeat", zap.Error(err), zap.Any("claim", claim))
		return "", err
	}

	return token, err
}

func (s Service) Register(ctx context.Context, mobile, password string) (token string, err error) {
	ctx, persist := db.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()
	logger := ctxlog.Extract(ctx)

	existedUser, err := s.userRepo.GetUserByMobile(ctx, mobile)
	if err == nil && existedUser != nil {
		logger.Error("user already existed", zap.Error(err), zap.Any("user", existedUser))
		return "", errors.New("user already existed")
	}

	user, err := s.userRepo.Create(ctx, mobile, password)
	if err != nil {
		logger.Error("create user defeat", zap.Error(err), zap.String("mobile", mobile), zap.String("password", password))
		return "", err
	}

	claim := &auth.UserClaim{
		Id: user.Id,
	}

	token, err = auth.GenerateToken(ctx, claim)
	if err != nil {
		logger.Error("generate token defeat", zap.Error(err), zap.Any("claim", claim))
		return "", err
	}

	return token, nil
}

func (s Service) Update(ctx context.Context, user *model.User) (u *model.User, err error) {
	ctx, persist := db.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()
	logger := ctxlog.Extract(ctx)

	user, err = s.userRepo.Update(ctx, user)
	if err != nil {
		logger.Error("update user defeat", zap.Error(err), zap.Any("user", user))
		return nil, err
	}

	return user, nil
}

func (s Service) GetUser(ctx context.Context, openId, mobile string, userId int64) (user *model.User, err error) {
	ctx, persist := db.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()
	logger := ctxlog.Extract(ctx)

	user = &model.User{}

	if openId != "" {
		user, err = s.userRepo.GetUserByOpenId(ctx, openId)
		if err != nil {
			logger.Error("get user by open id defeat", zap.Error(err), zap.String("openId", openId))
			return nil, err
		}
	}

	if user != nil && mobile != "" {
		user, err = s.userRepo.GetUserByMobile(ctx, mobile)
		if err != nil {
			logger.Error("get user by mobile defeat", zap.Error(err), zap.String("mobile", mobile))
			return nil, err
		}
	}

	if user != nil && userId != 0 {
		user, err = s.userRepo.GetUserById(ctx, userId)
		if err != nil {
			logger.Error("get user by user id defeat", zap.Error(err), zap.Int64("user id", userId))
			return nil, err
		}
	}

	user.Salt = ""
	user.Password = ""

	return user, nil
}

var _ IService = (*Service)(nil)
