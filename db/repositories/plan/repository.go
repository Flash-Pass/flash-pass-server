package plan

import (
	"errors"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	plan query.IPlanDo
}

type IRepository interface {
	Create(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	IsPlanBelongToUser(ctx *gin.Context, planId, userId uint64) (bool, error)
	Get(ctx *gin.Context, planId uint64) (*model.Plan, error)
	Update(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx *gin.Context, planId uint64) error
	GetList(ctx *gin.Context, userId uint64) ([]*model.Plan, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		plan: query.Plan.WithContext(db.Statement.Context),
	}
}

func (r *Repository) Create(ctx *gin.Context, plan *model.Plan) (*model.Plan, error) {
	logger := ctxlog.GetLogger(ctx)

	if err := r.plan.Create(plan); err != nil {
		logger.Error("create plan failed", zap.Error(err), zap.Any("plan", plan))
		return nil, err
	}

	return plan, nil
}

func (r *Repository) IsPlanBelongToUser(ctx *gin.Context, planId, userId uint64) (bool, error) {
	logger := ctxlog.GetLogger(ctx)

	plan, err := r.Get(ctx, planId)
	if err != nil {
		return false, err
	}

	if plan.CreatedBy != userId {
		logger.Warn("plan not belong to user", zap.Uint64("planId", planId), zap.Uint64("userId", userId))
		return false, err
	}

	return true, nil
}

func (r *Repository) Get(ctx *gin.Context, planId uint64) (*model.Plan, error) {
	logger := ctxlog.GetLogger(ctx)

	plan, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(planId)).First()
	if err != nil {
		logger.Error("get plan failed", zap.Error(err), zap.Uint64("planId", planId))
		return nil, err
	}

	return plan, err
}

func (r *Repository) Update(ctx *gin.Context, plan *model.Plan) (*model.Plan, error) {
	logger := ctxlog.GetLogger(ctx)

	if _, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(plan.Id)).Updates(plan); err != nil {
		logger.Error("update defeat", zap.Any("plan", plan), zap.Error(err))
		return plan, err
	}

	return plan, nil
}

func (r *Repository) Delete(ctx *gin.Context, planId uint64) error {
	logger := ctxlog.GetLogger(ctx)

	affected, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(planId),
		query.Plan.IsDeleted.Is(false)).Update(query.Plan.IsDeleted, true)
	if affected.RowsAffected == 0 {
		return errors.New("plan not found or plan has been deleted")
	}
	if err != nil {
		logger.Error("delete plan failed", zap.Error(err), zap.Uint64("planId", planId))
		return err
	}

	return nil
}

func (r *Repository) GetList(ctx *gin.Context, userId uint64) ([]*model.Plan, error) {
	logger := ctxlog.GetLogger(ctx)

	plans, err := r.plan.WithContext(ctx).Where(query.Plan.CreatedBy.Eq(userId)).Find()
	if err != nil {
		logger.Error("get plan list failed", zap.Error(err), zap.Uint64("userId", userId))
		return nil, err
	}

	return plans, nil
}

var _ IRepository = (*Repository)(nil)
