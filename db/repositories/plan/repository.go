package plan

import (
	"context"
	"errors"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	plan query.IPlanDo
	task query.ITaskDo
}

type IRepository interface {
	Create(ctx context.Context, plan *model.Plan) (*model.Plan, error)
	IsPlanBelongToUser(ctx context.Context, planId, userId int64) (bool, error)
	Get(ctx context.Context, planId int64) (*model.Plan, error)
	Update(ctx context.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx context.Context, planId int64) error
	GetList(ctx context.Context, userId int64) ([]*model.Plan, error)
	GetPlanByTaskId(ctx context.Context, taskId int64) (*model.Plan, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		plan: query.Plan.WithContext(db.Statement.Context),
		task: query.Task.WithContext(db.Statement.Context),
	}
}

func (r *Repository) Create(ctx context.Context, plan *model.Plan) (*model.Plan, error) {
	logger := ctxlog.Extract(ctx)

	if err := r.plan.Create(plan); err != nil {
		logger.Error("create plan failed", zap.Error(err), zap.Any("plan", plan))
		return nil, err
	}

	return plan, nil
}

func (r *Repository) IsPlanBelongToUser(ctx context.Context, planId, userId int64) (bool, error) {
	logger := ctxlog.Extract(ctx)

	plan, err := r.Get(ctx, planId)
	if err != nil {
		return false, err
	}

	if plan.CreatedBy != userId {
		logger.Warn("plan not belong to user", zap.Int64("planId", planId), zap.Int64("userId", userId))
		return false, err
	}

	return true, nil
}

func (r *Repository) Get(ctx context.Context, planId int64) (*model.Plan, error) {
	logger := ctxlog.Extract(ctx)

	plan, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(planId)).First()
	if err != nil {
		logger.Error("get plan failed", zap.Error(err), zap.Int64("planId", planId))
		return nil, err
	}

	return plan, err
}

func (r *Repository) Update(ctx context.Context, plan *model.Plan) (*model.Plan, error) {
	logger := ctxlog.Extract(ctx)

	if _, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(plan.Id)).Updates(plan); err != nil {
		logger.Error("update defeat", zap.Any("plan", plan), zap.Error(err))
		return plan, err
	}

	return plan, nil
}

func (r *Repository) Delete(ctx context.Context, planId int64) error {
	logger := ctxlog.Extract(ctx)

	affected, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(planId),
		query.Plan.IsDeleted.Is(false)).Update(query.Plan.IsDeleted, true)
	if affected.RowsAffected == 0 {
		return errors.New("plan not found or plan has been deleted")
	}
	if err != nil {
		logger.Error("delete plan failed", zap.Error(err), zap.Int64("planId", planId))
		return err
	}

	return nil
}

func (r *Repository) GetList(ctx context.Context, userId int64) ([]*model.Plan, error) {
	logger := ctxlog.Extract(ctx)

	plans, err := r.plan.WithContext(ctx).Where(query.Plan.CreatedBy.Eq(userId)).Find()
	if err != nil {
		logger.Error("get plan list failed", zap.Error(err), zap.Int64("userId", userId))
		return nil, err
	}

	return plans, nil
}

func (r *Repository) GetPlanByTaskId(ctx context.Context, taskId int64) (*model.Plan, error) {
	logger := ctxlog.Extract(ctx)

	task, err := r.task.WithContext(ctx).Where(query.Task.Id.Eq(taskId)).First()
	if err != nil {
		logger.Error("get task failed", zap.Error(err), zap.Int64("taskId", taskId))
		return nil, err
	}

	plan, err := r.plan.WithContext(ctx).Where(query.Plan.Id.Eq(task.PlanId)).First()
	if err != nil {
		logger.Error("get plan failed", zap.Error(err), zap.Int64("planId", task.PlanId))
		return nil, err
	}

	return plan, nil
}

var _ IRepository = (*Repository)(nil)
