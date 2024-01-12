package user

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db"

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
	Create(ctx context.Context, mobile, password string) (*model.User, error)
	CheckPassword(ctx context.Context, mobile, password string) (userId int64, ok bool)
	GetUserById(ctx context.Context, userId int64) (*model.User, error)
	GetUserByOpenId(ctx context.Context, openId string) (*model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
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

func (r *Repository) Create(ctx context.Context, mobile, password string) (*model.User, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	salt, err := r.generator.GetSalt()
	if err != nil {
		logger.Error("get salt defeat", zap.Error(err))
		return nil, err
	}

	pwd, _ := r.encryptHandle.PasswordEncrypt(password, salt)

	user := &model.User{
		Base: model.Base{
			Id: r.snowflakeHandle.GetId().Int64(),
		},
		OpenId:   "",
		Username: mobile,
		Password: pwd,
		Salt:     salt,
		Nickname: "未来卷王",
		Mobile:   mobile,
		Avatar:   "",
	}

	if err := tx.User.Create(user); err != nil {
		logger.Error("create user defeat", zap.Error(err), zap.Any("user", user))
		return nil, err
	}

	return user, nil
}

func (r *Repository) CheckPassword(ctx context.Context, mobile, password string) (userId int64, ok bool) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	user, err := tx.User.WithContext(ctx).Where(query.User.Username.Eq(mobile)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("mobile", mobile))
		return 0, false
	}

	if pwd, _ := r.encryptHandle.PasswordEncrypt(password, user.Salt); pwd != user.Password {
		logger.Warn("check password defeat", zap.String("mobile", mobile), zap.String("password", password))
		return 0, false
	}

	return user.Id, true
}

func (r *Repository) GetUserByOpenId(ctx context.Context, openId string) (*model.User, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	user, err := tx.User.WithContext(ctx).Where(query.User.OpenId.Eq(openId)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("open id", openId))
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserById(ctx context.Context, userId int64) (*model.User, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	user, err := tx.User.WithContext(ctx).Where(query.User.Id.Eq(userId)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.Int64("user id", userId))
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByMobile(ctx context.Context, mobile string) (*model.User, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	user, err := tx.User.WithContext(ctx).Where(query.User.Mobile.Eq(mobile)).First()
	if err != nil {
		logger.Error("user not found", zap.Error(err), zap.String("mobile", mobile))
		return nil, err
	}

	return user, nil
}

func (r *Repository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	if _, err := tx.User.Where(query.User.Id.Eq(user.Id)).Update(
		query.User.Nickname, user.Nickname,
	); err != nil {
		logger.Error("update user defeat", zap.Error(err), zap.Any("user", user))
		return nil, err
	}

	user, _ = tx.User.Where(query.User.Id.Eq(user.Id)).First()
	return user, nil
}

var _ IRepository = (*Repository)(nil)
