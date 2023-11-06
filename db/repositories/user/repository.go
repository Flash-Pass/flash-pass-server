package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
)

type Repository struct {
	user            query.IUserDo
	generator       generator
	encryptHandle   encryptHandle
	snowflakeHandle snowflake.IHandle
}

type IRepository interface {
	Create(ctx *gin.Context, mobile, password string) (*model.User, error)
	CheckPassword(ctx *gin.Context, mobile, password string) (token string, ok bool)
	GetUserById(ctx *gin.Context, userId string) (*model.User, error)
	GetUserByOpenId(ctx *gin.Context, openId string) (*model.User, error)
	GetUserByMobile(ctx *gin.Context, mobile string) (*model.User, error)
	Update(ctx *gin.Context, user *model.User) (*model.User, error)
}

type generator interface {
	GetSalt() (string, error)
}

type encryptHandle interface {
	PasswordEncrypt(password, salt string) (string, error)
}

func NewRepository(db *gorm.DB, g generator, e encryptHandle, s snowflake.IHandle) *Repository {
	return &Repository{
		user:            query.User.WithContext(db.Statement.Context),
		generator:       g,
		encryptHandle:   e,
		snowflakeHandle: s,
	}
}

func (r *Repository) Create(ctx *gin.Context, mobile, password string) (*model.User, error) {
	logger := ctxlog.GetLogger(ctx)

	salt, err := r.generator.GetSalt()
	if err != nil {
		logger.Error("get salt defeat", zap.Error(err))
		return nil, err
	}

	pwd, _ := r.encryptHandle.PasswordEncrypt(password, salt)

	user := &model.User{
		Base: model.Base{
			Id: r.snowflakeHandle.GetId().String(),
		},
		OpenId:   "",
		Username: mobile,
		Password: pwd,
		Salt:     salt,
		Nickname: "未来卷王",
		Mobile:   mobile,
		Avatar:   "",
	}

	if err := r.user.Create(user); err != nil {
		logger.Error("create user defeat", zap.Error(err), zap.Any("user", user))
		return nil, err
	}

	return user, nil
}

func (r *Repository) CheckPassword(ctx *gin.Context, mobile, password string) (token string, ok bool) {
	logger := ctxlog.GetLogger(ctx)

	user, err := r.user.WithContext(ctx).Where(query.User.Username.Eq(mobile)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("mobile", mobile))
		return "", false
	}

	if pwd, _ := r.encryptHandle.PasswordEncrypt(password, user.Salt); pwd != user.Password {
		logger.Warn("check password defeat", zap.String("mobile", mobile), zap.String("password", password))
		return "", false
	}

	return user.Id, true
}

func (r *Repository) GetUserByOpenId(ctx *gin.Context, openId string) (*model.User, error) {
	logger := ctxlog.GetLogger(ctx)

	user, err := r.user.WithContext(ctx).Where(query.User.OpenId.Eq(openId)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("open id", openId))
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserById(ctx *gin.Context, userId string) (*model.User, error) {
	logger := ctxlog.GetLogger(ctx)

	user, err := r.user.WithContext(ctx).Where(query.User.Id.Eq(userId)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("user id", userId))
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByMobile(ctx *gin.Context, mobile string) (*model.User, error) {
	logger := ctxlog.GetLogger(ctx)

	user, err := r.user.WithContext(ctx).Where(query.User.Mobile.Eq(mobile)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("mobile", mobile))
		return nil, err
	}

	return user, nil
}

func (r *Repository) Update(ctx *gin.Context, user *model.User) (*model.User, error) {
	logger := ctxlog.GetLogger(ctx)

	if _, err := r.user.Where(query.User.Id.Eq(user.Id)).Update(
		query.User.Nickname, user.Nickname,
	); err != nil {
		logger.Error("update user defeat", zap.Error(err), zap.Any("user", user))
		return nil, err
	}

	user, _ = r.user.Where(query.User.Id.Eq(user.Id)).First()
	return user, nil
}

var _ IRepository = (*Repository)(nil)
